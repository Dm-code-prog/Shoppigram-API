package webhooks

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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

	// ChannelStorage is an interface for storing and retrieving data about Telegram channels
	// from our own storage
	ChannelStorage interface {
		CreateOrUpdateTelegramChannel(ctx context.Context, req CreateOrUpdateTelegramChannelRequest) error
	}

	Service struct {
		channelStorage  ChannelStorage
		log             *zap.Logger
		shoppigramBotID int64
	}
)

// New returns a new instance of the Service
func New(channelStorage ChannelStorage, log *zap.Logger, shoppigramBotID int64) *Service {
	return &Service{
		channelStorage:  channelStorage,
		log:             log,
		shoppigramBotID: shoppigramBotID,
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
func (s *Service) HandleTelegramWebhook(ctx context.Context, update tgbotapi.Update) error {
	switch {
	case s.isUpdateTypeShoppigramBotAddedToChannelAsAdmin(update):
		return s.handleUpdateTypeShoppigramBotAddedToChannelAsAdmin(ctx, update)
	default:
		b, err := json.MarshalIndent(update, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshal data")
		}
		s.log.Info(
			"received Telegram webhook, but we don't have a handler for this type of update",
			zap.Any("webhook_data", json.RawMessage(b)),
		)
	}

	return nil
}

func (s *Service) handleUpdateTypeShoppigramBotAddedToChannelAsAdmin(ctx context.Context, update tgbotapi.Update) error {
	event := update.MyChatMember

	err := s.channelStorage.CreateOrUpdateTelegramChannel(ctx, CreateOrUpdateTelegramChannelRequest{
		ExternalID:      event.Chat.ID,
		Title:           event.Chat.Title,
		Name:            event.Chat.UserName,
		OwnerExternalID: event.From.ID,
		// Private channels don't have a name (username), only title
		// Public channels have a unique name (username) and a title
		IsPublic: event.Chat.UserName != "",
	})
	if err != nil {
		return errors.Wrap(err, "failed to create or update a Telegram channel")
	}

	return nil
}

func (s *Service) isUpdateTypeShoppigramBotAddedToChannelAsAdmin(update tgbotapi.Update) bool {
	if update.MyChatMember == nil {
		return false
	}
	event := update.MyChatMember

	if event.Chat.Type != "channel" {
		return false
	}

	if event.NewChatMember.Status != "administrator" {
		return false
	}

	if event.NewChatMember.User.ID != s.shoppigramBotID {
		return false
	}

	if !event.NewChatMember.CanPostMessages {
		return false
	}

	return true
}
