package notifications

import (
	"context"
	"embed"
	"github.com/shoppigram-com/marketplace-api/packages/cloudwatchcollector"
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

	metricsStatusOK = "OK"
	metricsStatusKO = "KO"
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

		GetNotificationsForNewMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]NewShopNotification, error)
		GetNotificationsForVerifiedMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]VerifiedShopNotification, error)
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
	}
)

// New creates a new Service
func New(repo Repository, log *zap.Logger, newOrderJobTimeout, newShopJobTimeout, verifiedShopJobTimeout time.Duration, botToken, botName string) *Service {
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
		newOrderProcessingTimer:            newOrderJobTimeout,
		newMarketplaceProcessingTimer:      newShopJobTimeout,
		verifiedMarketplaceProcessingTimer: verifiedShopJobTimeout,
		bot:                                bot,
		botName:                            botName,
	}
}

// Shutdown stops all the notifications
func (s *Service) Shutdown() error {
	s.cancel()
	<-s.ctx.Done()
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

// NotifyGreetings sends a greeting message to a user
func (s *Service) NotifyGreetings(_ context.Context, request NotifyGreetingsRequest) error {
	messageText, err := BuildGreetigsMessage(
		checkAndGetLangCode(request.UserLanguage),
	)
	if err != nil {
		return errors.Wrap(err, "BuildGreetigsMessage")
	}
	msg := tgbotapi.NewMessage(request.UserExternalID, messageText)
	_, err = s.SendMessage(msg)
	return err
}

// SendMarketplaceBanner sends a marketplace banner to a Telegram channel
func (s *Service) SendMarketplaceBanner(_ context.Context, params SendMarketplaceBannerParams) (int64, error) {
	msg := tgbotapi.NewMessage(params.ChannelChatID, params.Message)
	addButtonsToMessage(
		&msg,
		telegramButtonData{
			getTranslation("ru", "go-to-the-store"),
			params.WebAppLink,
		},
	)

	m, err := s.SendMessage(msg)
	if err != nil {
		return 0, errors.Wrap(err, "SendMessage")
	}

	return int64(m.MessageID), nil
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

// SendMessage sends a message, handles errors and publishes metrics
func (s *Service) SendMessage(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	message, err := s.bot.Send(msg)
	defer func() {
		status := metricsStatusOK
		if err != nil {
			status = metricsStatusKO
		}
		cloudwatchcollector.Increment("telegram_bot_api_send_message", cloudwatchcollector.Dimensions{
			"status": status,
		})
	}()
	if err != nil {
		return tgbotapi.Message{}, errors.Wrap(err, "send message to Telegram")
	}

	return message, nil
}
