package notifications

import (
	"context"
	"embed"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/templates/en"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/templates/ru"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	//go:embed templates/*/*/*.md
	templates      embed.FS
	validLangCodes []string = []string{"ru", "en"}
)

const (
	fallbackLanguage  = "ru"
	supportContactUrl = "https://t.me/ShoppigramSupport"

	orderNotifier                   = "new_order_notifications"
	newMarketplaceNotifierName      = "new_marketplace_notifications"
	verifiedMarketplaceNotifierName = "verified_marketplace_notifications"
	marketplaceBaseURL              = "https://web-app.shoppigram.com/app/"

	stateConfirmed = "confirmed"
	stateDone      = "done"
)

type (
	// Cursor defines the structure for a notify list cursor
	Cursor struct {
		Name            string
		CursorDate      time.Time
		LastProcessedID uuid.UUID
	}

	// Product is a marketplace product
	Product struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Quantity int       `json:"quantity"`
		Price    float64   `json:"price"`
	}

	// AddUserToNewOrderNotificationsRequest creates a new order notification
	// list entry for some marketplace
	AddUserToNewOrderNotificationsRequest struct {
		WebAppID    uuid.UUID
		AdminChatID int64
	}

	// NotifyChannelIntegrationSuccessRequest contains the data required to notify a user about a successful
	// channel integration with Shoppigram
	NotifyChannelIntegrationSuccessRequest struct {
		UserExternalID    int64
		UserLanguage      string
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	// NotifyChannelIntegrationFailureRequest contains the data required to notify a user about a failure
	// during channel integration with Shoppigram
	NotifyChannelIntegrationFailureRequest struct {
		UserExternalID    int64
		UserLanguage      string
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	// NotifyBotRemovedFromChannelRequest contains the data required to notify a user about a bot removal
	NotifyBotRemovedFromChannelRequest struct {
		UserExternalID    int64
		UserLanguage      string
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	// SendMarketplaceBannerParams is a struct for request params to send a marketplace banner to a Telegram channel
	// with a TWA link button markup
	SendMarketplaceBannerParams struct {
		WebAppLink    string
		Message       string
		ChannelChatID int64
	}

	// PinNotificationParams is a struct for request params to pin a message in a Telegram channel
	PinNotificationParams struct {
		ChatID    int64
		MessageID int64
	}

	// NotifyGreetingsRequest contains the initial greeting message
	NotifyGreetingsRequest struct {
		UserExternalID int64
		UserLanguage   string
	}

	adminNotitfication struct {
		Id       int64
		Language string
	}

	telegramButtonData struct {
		text string
		link string
	}

	pageDataParam struct {
		key   string
		value any
	}

	// Repository provides access to the user storage
	Repository interface {
		GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]adminNotitfication, error)

		GetReviewersNotificationList(ctx context.Context) ([]int64, error)

		GetNotifierCursor(ctx context.Context, name string) (Cursor, error)

		UpdateNotifierCursor(ctx context.Context, cur Cursor) error

		// GetProductCustomMessage returns a custom notification message for a product
		// empty string is returned if no message is found
		GetProductCustomMessage(ctx context.Context, productID uuid.UUID, state string) (string, error)

		// GetProductCustomMediaForward gets a custom media forward information for a product
		GetProductCustomMediaForward(ctx context.Context, productID uuid.UUID, state string) (fromChatID int64, messageID int64, err error)

		GetNotificationsForNewMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]NewMarketplaceNotification, error)
		GetNotificationsForVerifiedMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]VerifiedMarketplaceNotification, error)
		AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsRequest) error

		// GetNotificationsForOrders returns a list of orders that were updated since the last run
		// along with extra information about the buyer, seller and the shop.
		GetNotificationsForOrders(ctx context.Context, cursor Cursor) ([]OrderNotification, error)
	}

	// Service provides user operations
	Service struct {
		repo                               Repository
		log                                *zap.Logger
		ctx                                context.Context
		cancel                             context.CancelFunc
		newOrderProcessingTimer            time.Duration
		newMarketplaceProcessingTimer      time.Duration
		verifiedMarketplaceProcessingTimer time.Duration
		bot                                *tgbotapi.BotAPI
		botName                            string
		bucketUrl                          string
	}
)

// New creates a new Service
func New(repo Repository, log *zap.Logger, newOrderProcessingTimer time.Duration, newMarketplaceProcessingTimer time.Duration, verifiedMarketplaceProcessingTimer time.Duration, botToken string, botName string, bucketUrl string) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.With(
			zap.String("method", "tgbotapi.NewBotAPI"),
		).Fatal(err.Error())
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Service{
		repo:                               repo,
		log:                                log,
		ctx:                                ctx,
		cancel:                             cancel,
		newOrderProcessingTimer:            newOrderProcessingTimer,
		newMarketplaceProcessingTimer:      newMarketplaceProcessingTimer,
		verifiedMarketplaceProcessingTimer: verifiedMarketplaceProcessingTimer,
		bot:                                bot,
		botName:                            botName,
		bucketUrl:                          bucketUrl,
	}
}

// Shutdown stops all the notifications
func (s *Service) Shutdown() error {
	s.cancel()
	<-s.ctx.Done()
	return nil
}

// sendMessageToChat sends message specified in MsgText to chat with id chatID
func (s *Service) sendMessageToChat(chatID int64, msgTxt string) error {
	msg := tgbotapi.NewMessage(chatID, msgTxt)
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	_, err := s.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "bot.Send")
	}
	return nil
}

// AddUserToNewOrderNotifications creates a new order notification
// list entry for some marketplace
func (s *Service) AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsRequest) error {
	err := s.repo.AddUserToNewOrderNotifications(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.AddUserToNewOrderNotifications")
	}

	return nil
}

// NotifyChannelIntegrationSuccess notifies a user about a successful
// channel integration with Shoppigram
func (s *Service) NotifyChannelIntegrationSuccess(_ context.Context, request NotifyChannelIntegrationSuccessRequest) error {
	message := ChannelIntegrationSuccessNotification(request)
	userLnag := s.checkAndGetLangCode(message.UserLanguage)
	msgTxt, err := message.BuildMessage(userLnag)
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	buttonText := getTranslation(userLnag, "try-new-features")
	addTelegramButtonsToMessage(
		&msg,
		telegramButtonData{
			buttonText,
			"https://t.me/" + s.botName + "/app",
		},
	)

	_, err = s.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "bot.Send")
	}

	return nil
}

// NotifyGreetings sends a greeting message to a user
func (s *Service) NotifyGreetings(_ context.Context, request NotifyGreetingsRequest) error {
	userLang := s.checkAndGetLangCode(request.UserLanguage)
	messageText, err := BuildGreetigsMessage(userLang)
	if err != nil {
		return errors.Wrap(err, "BuildGreetigsMessage()")
	}
	msg := tgbotapi.NewMessage(request.UserExternalID, messageText)
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	_, err = s.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "bot.Send")
	}

	return nil
}

// SendMarketplaceBanner sends a marketplace banner to a Telegram channel
func (s *Service) SendMarketplaceBanner(_ context.Context, params SendMarketplaceBannerParams) (message int64, err error) {
	msg := tgbotapi.NewMessage(params.ChannelChatID, params.Message)
	buttonText := getTranslation("ru", "go-to-the-store")
	addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, params.WebAppLink})

	Message, err := s.bot.Send(msg)
	if err != nil {
		return 0, errors.Wrap(err, "bot.Send")
	}

	return int64(Message.MessageID), nil
}

// PinNotification pins a message in a Telegram channel
func (s *Service) PinNotification(_ context.Context, req PinNotificationParams) error {
	_, err := s.bot.Request(tgbotapi.PinChatMessageConfig{
		ChatID:    req.ChatID,
		MessageID: int(req.MessageID),
	})
	if err != nil {
		return errors.Wrap(err, "bot.PinChatMessage")
	}

	return nil
}

// NotifyChannelIntegrationFailure notifies a user about a failure
// happened during channel integration with Shoppigram
func (s *Service) NotifyChannelIntegrationFailure(_ context.Context, request NotifyChannelIntegrationFailureRequest) error {
	message := ChannelIntegrationFailureNotification(request)
	userLnag := s.checkAndGetLangCode(message.UserLanguage)
	msgTxt, err := message.BuildMessage(userLnag)
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	tgLink := createAddBotAsAdminLink(s.botName)
	buttonText := getTranslation(userLnag, "try-again")
	if err != nil {
		return errors.Wrap(err, "getButtonText(\"try-again\")")
	}
	addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

	_, err = s.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "bot.Send")
	}

	return nil
}

// NotifyChannelIntegrationFailure notifies a user about a failure
// happened during channel integration with Shoppigram
func (s *Service) NotifyBotRemovedFromChannel(_ context.Context, request NotifyBotRemovedFromChannelRequest) error {

	message := BotRemovedFromChannelNotification(request)
	userLnag := s.checkAndGetLangCode(message.UserLanguage)
	msgTxt, err := message.BuildMessage(userLnag)
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	tgLink := createAddBotAsAdminLink(s.botName)
	buttonText := getTranslation(userLnag, "add-bot-as-admin")
	if err != nil {
		return errors.Wrap(err, "getButtonText(\"add-bot-as-admin\")")
	}

	addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

	_, err = s.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "bot.Send")
	}

	return nil
}

func addTelegramButtonsToMessage(msg *tgbotapi.MessageConfig, messageData ...telegramButtonData) {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, v := range messageData {
		button := tgbotapi.NewInlineKeyboardButtonURL(v.text, v.link)
		row := tgbotapi.NewInlineKeyboardRow(button)
		rows = append(rows, row)
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		rows...,
	)
}

func (s *Service) getTelegramLink(path string, pageData ...pageDataParam) (string, error) {
	pageDataParams := make(map[string]any)
	for _, v := range pageData {
		pageDataParams[v.key] = v.value
	}

	tmaLink, err := TMALinkingScheme{
		PageName: "/app/" + path,
		PageData: pageDataParams,
	}.ToBase64String()
	if err != nil {
		return "", errors.Wrap(err, "TMALinkingScheme.ToBase64String()")
	}
	fullLink := "https://t.me/" + s.botName + "/app?startapp=" + tmaLink
	return fullLink, nil
}

func (s *Service) checkAndGetLangCode(lang string) string {
	if isLanguageValid(lang) {
		return lang
	}
	return fallbackLanguage
}

func (s *Service) handleTelegramSendError(err error, chatID int64) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "chat not found") {
		s.log.With(
			zap.String("method", "bot.Send"),
			zap.String("user_id", strconv.FormatInt(chatID, 10)),
		).Warn("chat not found")
		return nil
	}
	return errors.Wrap(err, "bot.Send")
}

func getTranslation(lang, key string) string {
	switch lang {
	case langRu:
		return ru.Translations[key]
	case langEn:
		return en.Translations[key]
	default:
		return ru.Translations[key]
	}
}

func createAddBotAsAdminLink(botName string) string {
	return "https://t.me/" + botName + "?startchannel&admin=post_messages+edit_messages+pin_messages"
}
