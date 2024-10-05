package webhooks

import (
	"context"
	"encoding/json"
	"github.com/shoppigram-com/marketplace-api/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	fallbackLanguage = "ru"
)

type (
	// CreateOrUpdateTelegramChannelRequest contains the data about a Telegram channel, Shoppigram bot is added to
	CreateOrUpdateTelegramChannelRequest struct {
		ExternalID      int64
		Title           string
		Name            string
		OwnerExternalID int64
		IsPublic        bool
	}

	// DeleteTelegramChannelRequest specifies the ID of the channel to delete
	// from the database
	DeleteTelegramChannelRequest struct {
		ExternalID int64
	}

	// GetTelegramChannelOwnerRequest contains chat id of a channel
	GetTelegramChannelOwnerRequest struct {
		ChannelChatId int64
	}

	// GetTelegramChannelOwnerResponse contains channel's owner external id
	GetTelegramChannelOwnerResponse struct {
		ChatId int64
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

	// NotifyBotRemovedFromChannelRequest contains the data required to notify a user about a removal
	// of a Shoppigram bot from channel
	NotifyBotRemovedFromChannelRequest struct {
		UserExternalID    int64
		UserLanguage      string
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	// TelegramService is the service for handling Telegram webhooks
	TelegramService struct {
		repo              Repository
		notifier          Notifier
		log               *zap.Logger
		shoppigramBotID   int64
		shoppigramBotName string
	}

	// NotifyGreetingsRequest contains the initial greeting message
	// of the bot
	NotifyGreetingsRequest struct {
		UserExternalID int64
		UserLanguage   string
	}
)

// NewTelegram returns a new instance of the TelegramService
func NewTelegram(repo Repository, notifier Notifier, log *zap.Logger, shoppigramBotID int64, shoppigramBotName string) *TelegramService {
	return &TelegramService{
		repo:              repo,
		notifier:          notifier,
		log:               log,
		shoppigramBotID:   shoppigramBotID,
		shoppigramBotName: shoppigramBotName,
	}
}

// HandleTelegramWebhook is the entry point for a webhook request from Telegram.
//
// It is supposed to determine the type of Update and call the correct handler for the update.
// If there is no handler for the update, it should return nil and log the update as JSON.
//
// We differentiate the updates based on the optional fields of the Update struct. At most one optional field is present
// at any given update. However, we can have more than one handler for a given Telegram update type.
// In this case, each handler provides a function that determines if it can handle the update.
func (s *TelegramService) HandleTelegramWebhook(ctx context.Context, update tgbotapi.Update) error {
	switch {
	case s.isUpdateTypeAddedToChannel(update):
		return s.handleAddedToChannel(ctx, update)
	case s.isUpdateTypeStartCommand(update):
		return s.handleStartCommand(ctx, update)
	case s.isUpdateTypeRemovedFromChannel(update):
		return s.handleRemovedFromChannel(ctx, update)
	default:
		b, _ := json.MarshalIndent(update, "", "  ")
		s.log.Info(
			"received Telegram webhook, but we were unable to handle it",
			zap.Any("webhook_data", json.RawMessage(b)),
		)
	}

	return nil
}

func (s *TelegramService) handleAddedToChannel(ctx context.Context, update tgbotapi.Update) error {
	event := update.MyChatMember

	handleFailure := func(err error) error {
		_ = s.notifier.NotifyChannelIntegrationFailure(ctx, NotifyChannelIntegrationFailureRequest{
			UserExternalID:    event.From.ID,
			UserLanguage:      event.From.LanguageCode,
			ChannelExternalID: event.Chat.ID,
			ChannelTitle:      event.Chat.Title,
			ChannelName:       event.Chat.UserName,
		})

		s.log.Error("telegram channel integration failed", logging.SilentError(err))
		return nil
	}

	if !update.MyChatMember.NewChatMember.CanPinMessages ||
		!update.MyChatMember.NewChatMember.CanEditMessages ||
		!update.MyChatMember.NewChatMember.CanPostMessages {
		return handleFailure(errors.New("bot doesn't have required permissions"))
	}

	err := s.repo.CreateOrUpdateTelegramChannel(ctx, CreateOrUpdateTelegramChannelRequest{
		ExternalID:      event.Chat.ID,
		Title:           event.Chat.Title,
		Name:            event.Chat.UserName,
		OwnerExternalID: event.From.ID,
		// Private channels don't have a name (username), only title
		// Public channels have a unique name (username) and a title
		IsPublic: event.Chat.UserName != "",
	})
	if err != nil {
		return handleFailure(errors.Wrap(err, "s.channelStorage.CreateOrUpdateTelegramChannel"))
	}

	err = s.notifier.NotifyChannelIntegrationSuccess(ctx, NotifyChannelIntegrationSuccessRequest{
		UserExternalID:    event.From.ID,
		UserLanguage:      event.From.LanguageCode,
		ChannelExternalID: event.Chat.ID,
		ChannelTitle:      event.Chat.Title,
		ChannelName:       event.Chat.UserName,
	})
	if err != nil {
		return handleFailure(errors.Wrap(err, "s.notifier.NotifyChannelIntegrationSuccess"))
	}

	return nil
}

func (s *TelegramService) handleRemovedFromChannel(ctx context.Context, update tgbotapi.Update) error {
	event := update.MyChatMember
	lang := event.From.LanguageCode
	if lang == "" {
		lang = fallbackLanguage
	}

	err := s.repo.DeleteTelegramChannel(ctx, DeleteTelegramChannelRequest{
		ExternalID: event.Chat.ID,
	})
	if err != nil {
		return errors.Wrap(err, "s.channelStorage.DeleteTelegramChannel")
	}

	return nil
}

func (s *TelegramService) handleStartCommand(ctx context.Context, update tgbotapi.Update) error {
	err := s.notifier.NotifyGreetings(ctx, NotifyGreetingsRequest{
		UserExternalID: update.Message.From.ID,
		UserLanguage:   update.Message.From.LanguageCode,
	})
	if err != nil {
		return errors.Wrap(err, "s.notifier.NotifyGreetings")
	}

	return nil
}
