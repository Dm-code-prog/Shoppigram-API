package notifications

import (
	"context"
	"embed"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"go.uber.org/zap"
)

//go:embed templates/*/*/*.md templates/*/*.json
var templates embed.FS
var validLangCodes []string = []string{"ru", "en"}

const fallbackLanguage = "ru"
const supportContactUrl = "https://t.me/ShoppigramSupport"
const pathToButtonsText = "/buttons.json"

type (
	// Cursor defines the structure for a notify list cursor
	Cursor struct {
		Name            string
		CursorDate      time.Time
		LastProcessedID uuid.UUID
	}

	// Product is a marketplace product
	Product struct {
		Name     string
		Quantity int
		Price    float64
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
		GetNotificationsForNewOrdersAfterCursor(ctx context.Context, cur Cursor) ([]NewOrderNotification, error)
		GetNotificationsForNewMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]NewMarketplaceNotification, error)
		GetNotificationsForVerifiedMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]VerifiedMarketplaceNotification, error)
		AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsRequest) error
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

const (
	newOrderNotifierName            = "new_order_notifications"
	newMarketplaceNotifierName      = "new_marketplace_notifications"
	verifiedMarketplaceNotifierName = "verified_marketplace_notifications"
	marketplaceBaseURL              = "https://web-app.shoppigram.com/app/"
)

// New creates a new user service
func New(repo Repository, log *zap.Logger, newOrderProcessingTimer time.Duration, newMarketplaceProcessingTimer time.Duration, verifiedMarketplaceProcessingTimer time.Duration, botToken string, botName string, bucketUrl string) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if newOrderProcessingTimer == 0 {
		log.Fatal("new order processing timer is not specified")
		return nil
	}
	if newMarketplaceProcessingTimer == 0 {
		log.Fatal("new marketplace processing timer is not specified")
		return nil
	}
	if botToken == "" {
		log.Fatal("bot token is not specified")
		return nil
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

// RunNewOrderNotifier starts a job that batch loads new orders
// and sends notifications to the owners of marketplaces
func (s *Service) RunNewOrderNotifier() error {
	ticker := time.NewTicker(s.newOrderProcessingTimer)
	for {
		select {
		case <-ticker.C:
			err := s.runNewOrderNotifierOnce()
			if err != nil {
				s.log.Error("runNewOrderNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runNewOrderNotifierOnce executes one iteration of loading a batch of new
// orders and sending notifications to the owners of marketplaces
func (s *Service) runNewOrderNotifierOnce() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, newOrderNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	orderNotifications, err := s.repo.GetNotificationsForNewOrdersAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForNewOrdersAfterCursor")
	}

	if len(orderNotifications) == 0 {
		return nil
	}

	s.log.With(
		zap.String("count", strconv.Itoa(len(orderNotifications))),
	).Info("sending notifications for new orders")
	err = s.sendNewOrderNotifications(orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendNewOrderNotifications")
	}

	lastElem := orderNotifications[len(orderNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.CreatedAt,
		LastProcessedID: lastElem.ID,
		Name:            newOrderNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// RunNewMarketplaceNotifier starts a job that batch loads new marketplaces
// and sends notifications to the reviewers of marketplaces
func (s *Service) RunNewMarketplaceNotifier() error {
	ticker := time.NewTicker(s.newMarketplaceProcessingTimer)
	for {
		select {
		case <-ticker.C:
			err := s.runNewMarketplaceNotifierOnce()
			if err != nil {
				s.log.Error("runNewMarketplaceNotifierOnce failed", logging.SilentError(err))
				s.log.Error("runNewMarketplaceNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runNewMarketplaceNotifierOnce executes one iteration of loading a batch of new
// marketplaces and sending notifications to the reviewers of marketplaces
func (s *Service) runNewMarketplaceNotifierOnce() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, newMarketplaceNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	marketplaceNotifications, err := s.repo.GetNotificationsForNewMarketplacesAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForNewMarketplacesAfterCursor")
	}

	if len(marketplaceNotifications) == 0 {
		return nil
	}

	s.log.With(
		zap.String("count", strconv.Itoa(len(marketplaceNotifications))),
	).Info("sending notifications for new marketplaces")

	err = s.sendNewMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		if strings.Contains(err.Error(), "chat not found") {
			s.log.With(
				zap.String("method", "s.sendNewMarketplaceNotifications"),
				zap.String("user_id", strconv.FormatInt(marketplaceNotifications[0].OwnerExternalID, 10)),
			).Warn("chat not found, skipping notification sending")
		} else {
			return errors.Wrap(err, "s.sendNewMarketplaceNotifications")
		}
	}

	lastElem := marketplaceNotifications[len(marketplaceNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.CreatedAt,
		LastProcessedID: lastElem.ID,
		Name:            newMarketplaceNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// RunVerifiedMarketplaceNotifier starts a job that batch loads verified marketplaces
// and sends notifications to the owners of those marketplaces
func (s *Service) RunVerifiedMarketplaceNotifier() error {
	ticker := time.NewTicker(s.verifiedMarketplaceProcessingTimer)

	for {
		select {
		case <-ticker.C:
			err := s.runVerifiedMarketplaceNotifierOnce()
			if err != nil {
				s.log.Error("runVerifiedMarketplaceNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runVerifiedMarketplaceNotifierOnce executes one iteration of loading a batch of
// verified marketplaces and sending notifications to the owners of those marketplaces
func (s *Service) runVerifiedMarketplaceNotifierOnce() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, verifiedMarketplaceNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	marketplaceNotifications, err := s.repo.GetNotificationsForVerifiedMarketplacesAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForVerifiedMarketplacesAfterCursor")
	}

	if len(marketplaceNotifications) == 0 {
		return nil
	}

	s.log.With(
		zap.String("count", strconv.Itoa(len(marketplaceNotifications))),
	).Info("sending notifications for verified marketplaces")
	err = s.sendVerifiedMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendVerifiedMarketplaceNotifications")
	}

	lastElem := marketplaceNotifications[len(marketplaceNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.VerifiedAt,
		LastProcessedID: lastElem.ID,
		Name:            verifiedMarketplaceNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// Shutdown stops all the notifications
func (s *Service) Shutdown() error {
	s.cancel()
	<-s.ctx.Done()
	return nil
}

// sendNewOrderNotifications sends batch of notifications for new orders
func (s *Service) sendNewOrderNotifications(orderNotifications []NewOrderNotification) error {
	for _, notification := range orderNotifications {
		nl, err := s.repo.GetAdminsNotificationList(s.ctx, notification.WebAppID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetAdminsNotificationList")
		}

		for _, v := range nl {
			notificationLang := s.checkAndGetLangCode(v.Language)
			adminMsgTxt, err := notification.BuildMessageAdmin(notificationLang)
			if err != nil {
				return errors.Wrap(err, "a.BuildMessageAdmin")
			}

			msg := tgbotapi.NewMessage(v.Id, adminMsgTxt)
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			tgLinkPath := notification.WebAppID.String() + "/order/" + notification.ID.String()
			tgLink, err := s.getTelegramLink(tgLinkPath)
			if err != nil {
				return errors.Wrap(err, "getTelegramLink()")
			}

			buttonText, err := getButtonText(notificationLang, "order-management")
			if err != nil {
				return errors.Wrap(err, "getButtonText(\"order-management\")")
			}
			addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

			_, err = s.bot.Send(msg)
			if err != nil {
				if strings.Contains(err.Error(), "chat not found") {
					s.log.With(
						zap.String("method", "bot.Send"),
						zap.String("user_id", strconv.FormatInt(v.Id, 10)),
					).Warn("chat not found")
					continue
				}
				return errors.Wrap(err, "s.sendMessageToChat")
			}
		}

		userLang := s.checkAndGetLangCode(notification.UserLanguage)

		customerMsgTxt, err := notification.BuildMessageCustomer(userLang)

		if err != nil {
			return errors.Wrap(err, "a.BuildMessageCustomer")
		}

		msg := tgbotapi.NewMessage(notification.ExternalUserID, customerMsgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		tgLinkPath := notification.WebAppID.String() + "/order/" + notification.ID.String()
		tgLink, err := s.getTelegramLink(tgLinkPath)
		if err != nil {
			return errors.Wrap(err, "getTelegramLink()")
		}
		buttonText, err := getButtonText(userLang, "view-order")
		if err != nil {
			return errors.Wrap(err, "getButtonText(\"view-order\")")
		}

		addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

		_, err = s.bot.Send(msg)
		if err != nil {
			if strings.Contains(err.Error(), "chat not found") {
				s.log.With(
					zap.String("method", "bot.Send"),
					zap.String("user_id", strconv.FormatInt(notification.ExternalUserID, 10)),
				).Warn("chat not found")
				continue
			}
			return errors.Wrap(err, "s.sendMessageToChat")
		}

	}

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

// sendNewMarketplaceNotifications sends batch of notifications for new marketplaces
func (s *Service) sendNewMarketplaceNotifications(marketplaceNotifications []NewMarketplaceNotification) error {
	reviewers, err := s.repo.GetReviewersNotificationList(s.ctx)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetReviewersNotificationList")
	}
	for _, n := range marketplaceNotifications {
		n.ImageBaseUrl = s.bucketUrl

		ownerLang := s.checkAndGetLangCode(n.OwnerLanguage)
		onVerificationMsgTxt, err := n.BuildMessageAdmin(ownerLang)
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}

		msg := tgbotapi.NewMessage(n.OwnerExternalID, onVerificationMsgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		tgLinkPath := n.ID.String()
		tgLink, err := s.getTelegramLink(tgLinkPath)
		if err != nil {
			return errors.Wrap(err, "getTelegramLink()")
		}
		buttonTextContactSupport, err := getButtonText(ownerLang, "contact-support")
		if err != nil {
			return errors.Wrap(err, "getButtonText(\"contact-support\")")
		}
		buttonTextViewStore, err := getButtonText(ownerLang, "view-store")
		if err != nil {
			return errors.Wrap(err, "getButtonText(\"view-store\")")
		}

		addTelegramButtonsToMessage(&msg,
			telegramButtonData{buttonTextContactSupport, supportContactUrl},
			telegramButtonData{buttonTextViewStore, tgLink},
		)

		_, err = s.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "bot.Send to chat:"+strconv.FormatInt(n.OwnerExternalID, 10))
		}

		for _, r := range reviewers {
			msgTxt, err := n.BuildMessageShoppigram("en")
			if err != nil {
				return errors.Wrap(err, "a.BuildMessageShoppigram")
			}
			err = s.sendMessageToChat(r, msgTxt)
			if err != nil {
				return errors.Wrap(err, "sendMessageToChat")
			}
		}

	}

	return nil
}

// sendVerifiedMarketplaceNotifications sends batch of notifications for verified marketplaces
func (s *Service) sendVerifiedMarketplaceNotifications(marketplaceNotifications []VerifiedMarketplaceNotification) error {
	for _, notification := range marketplaceNotifications {
		ownerLang := s.checkAndGetLangCode(notification.OwnerLanguage)
		msgTxt, err := notification.BuildMessage(ownerLang)
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}

		msg := tgbotapi.NewMessage(notification.OwnerExternalUserID, msgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		tgLinkPath := notification.ID.String()
		tgLink, err := s.getTelegramLink(tgLinkPath)
		if err != nil {
			return errors.Wrap(err, "getTelegramLink()")
		}

		buttonText, err := getButtonText(ownerLang, "continue-setting-up")
		if err != nil {
			return errors.Wrap(err, "getButtonText(\"continue-setting-up\")")
		}
		addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

		_, err = s.bot.Send(msg)
		if err != nil {
			if strings.Contains(err.Error(), "chat not found") {
				s.log.With(
					zap.String("method", "bot.Send"),
					zap.String("user_id", strconv.FormatInt(notification.OwnerExternalUserID, 10)),
				).Warn(err.Error())
				continue
			}
			return errors.Wrap(err, "bot.Send")
		}

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

	tgLink := "https://t.me/" + s.botName + "/app"
	buttonText, err := getButtonText(userLnag, "try-new-features")
	if err != nil {
		return errors.Wrap(err, "getButtonText(\"try-new-features\")")
	}
	addTelegramButtonsToMessage(&msg, telegramButtonData{buttonText, tgLink})

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
	buttonText, err := getButtonText("ru", "go-to-the-store")
	if err != nil {
		return 0, errors.Wrap(err, "getButtonText(\"go-to-the-store\")")
	}
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

	s.log.Warn("Language code is incorrect. Using fallback language",
		zap.String("passed value", lang),
		zap.String("fallback value", fallbackLanguage))

	return fallbackLanguage
}

func isLanguageValid(lang string) bool {
	for _, v := range validLangCodes {
		if lang == v {
			return true
		}
	}
	return false
}

func getButtonText(lang string, key string) (string, error) {
	var bt map[string]string
	data, err := templates.ReadFile("templates/" + lang + pathToButtonsText)
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile("+lang+".json)")
	}
	err = json.Unmarshal(data, &bt)
	if err != nil {
		return "", errors.Wrap(err, "json.Unmarshal(data, bt)")
	}
	return bt[key], nil
}
