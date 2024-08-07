package notifications

import (
	"context"
	"embed"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"go.uber.org/zap"
)

//go:embed templates/*/*.md
var templates embed.FS

const supportContactUrl = "https://t.me/ShoppigramSupport"

type (
	// Cursor defines the structure for a notify list cursor
	Cursor struct {
		Name            string
		CursorDate      time.Time
		LastProcessedID uuid.UUID
	}

	// Product is a marketplace product
	Product struct {
		Name          string
		Quantity      int
		Price         float64
		PriceCurrency string
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
	}

	// Repository provides access to the user storage
	Repository interface {
		GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error)
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
	marketplaceURL                  = "https://web-app.shoppigram.com/app/"
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

type telegramButtonData struct {
	text string
	link string
}

func addTelegramButtonsToMessage(msg *tgbotapi.MessageConfig, messageData ...telegramButtonData) {
	buttons := make([]tgbotapi.InlineKeyboardButton, len(messageData))
	for _, v := range messageData {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonURL(v.text, v.link))
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)

}

type pageDataParam struct {
	key   string
	value any
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

// sendNewOrderNotifications sends batch of notifications for new orders
func (s *Service) sendNewOrderNotifications(orderNotifications []NewOrderNotification) error {
	for _, notification := range orderNotifications {
		nl, err := s.repo.GetAdminsNotificationList(s.ctx, notification.WebAppID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetAdminsNotificationList")
		}
		adminMsgTxt, err := notification.BuildMessageAdmin()
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageAdmin")
		}

		for _, v := range nl {
			msg := tgbotapi.NewMessage(v, adminMsgTxt)
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			tgLinkPath := notification.WebAppID.String() + "/order/" + "notification.ID.String()"
			tgLink, err := s.getTelegramLink(tgLinkPath)
			if err != nil {
				return errors.Wrap(err, "getTelegramLink()")
			}
			addTelegramButtonsToMessage(&msg, telegramButtonData{"Управление заказом", tgLink})

			_, err = s.bot.Send(msg)
			if err != nil {
				if strings.Contains(err.Error(), "chat not found") {
					s.log.With(
						zap.String("method", "bot.Send"),
						zap.String("user_id", strconv.FormatInt(v, 10)),
					).Warn("chat not found")
					continue
				}
				return errors.Wrap(err, "s.sendMessageToChat")
			}
		}

		customerMsgTxt, err := notification.BuildMessageCustomer()

		if err != nil {
			return errors.Wrap(err, "a.BuildMessageCustomer")
		}

		msg := tgbotapi.NewMessage(notification.ExternalUserID, customerMsgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		tgLinkPath := "/app/" + notification.WebAppID.String() + "/order/" + notification.ID.String()
		tgLink, err := s.getTelegramLink(tgLinkPath)
		if err != nil {
			return errors.Wrap(err, "getTelegramLink()")
		}
		addTelegramButtonsToMessage(&msg, telegramButtonData{"Посмотреть заказ", tgLink})

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
		onVerificationMsgTxt, err := n.BuildMessageAdmin()
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

		addTelegramButtonsToMessage(&msg,
			telegramButtonData{"Связаться с поддержкой", supportContactUrl},
			telegramButtonData{"Посмотреть магазин", tgLink},
		)

		_, err = s.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "bot.Send to chat:"+strconv.FormatInt(n.OwnerExternalID, 10))
		}

		for _, r := range reviewers {
			msgTxt, err := n.BuildMessageShoppigram()
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
		msgTxt, err := notification.BuildMessage()
		if err != nil {
			return errors.Wrap(err, "a.BuildMessageShoppigram")
		}

		msg := tgbotapi.NewMessage(notification.OwnerExternalUserID, msgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2
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
	msgTxt, err := message.BuildMessage()
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	_, err = s.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "bot.Send")
	}

	return nil
}

// NotifyGreetings sends a greeting message to a user
func (s *Service) NotifyGreetings(_ context.Context, request NotifyGreetingsRequest) error {
	messageText, err := BuildGreetigsMessage()
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
	button := tgbotapi.NewInlineKeyboardButtonURL("Перейти в магазин", params.WebAppLink)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			button,
		))

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
