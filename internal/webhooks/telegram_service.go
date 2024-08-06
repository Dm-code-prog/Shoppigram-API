package webhooks

import (
	"context"
	"encoding/json"
	"fmt"

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

	// NotifyChannelIntegrationSuccessRequest contains the data required to notify a user about a successful
	// channel integration with Shoppigram
	NotifyChannelIntegrationSuccessRequest struct {
		UserExternalID    int64
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	// TelegramService is the service for handling Telegram webhooks
	TelegramService struct {
		channelStorage    ChannelStorage
		notifier          Notifier
		log               *zap.Logger
		shoppigramBotID   int64
		shoppigramBotName string
	}

	// NotifyGreetingsRequest contains the initial greeting message
	// of the bot
	NotifyGreetingsRequest struct {
		UserExternalID  int64
		GreetingMessage string
	}

	// Notifier is the service for notifications
	// The interface requires a method for notifying a user about a successful
	// channel integration with Shoppigram
	Notifier interface {
		NotifyChannelIntegrationSuccess(ctx context.Context, request NotifyChannelIntegrationSuccessRequest) error
		NotifyGreetings(ctx context.Context, request NotifyGreetingsRequest) error
	}
)

// NewTelegram returns a new instance of the TelegramService
func NewTelegram(channelStorage ChannelStorage, notifier Notifier, log *zap.Logger, shoppigramBotID int64, shoppigramBotName string) *TelegramService {
	return &TelegramService{
		channelStorage:    channelStorage,
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
	case s.isUpdateTypeShoppigramBotAddedToChannelAsAdmin(update):
		return s.handleUpdateTypeShoppigramBotAddedToChannelAsAdmin(ctx, update)
	case s.isUpdateTypeStartCommand(update):
		return s.handleUpdateTypeStartCommand(ctx, update)
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

func (s *TelegramService) handleUpdateTypeShoppigramBotAddedToChannelAsAdmin(ctx context.Context, update tgbotapi.Update) error {
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
		return errors.Wrap(err, "s.channelStorage.CreateOrUpdateTelegramChannel")
	}

	err = s.notifier.NotifyChannelIntegrationSuccess(ctx, NotifyChannelIntegrationSuccessRequest{
		UserExternalID:    event.From.ID,
		ChannelExternalID: event.Chat.ID,
		ChannelTitle:      event.Chat.Title,
		ChannelName:       event.Chat.UserName,
	})
	if err != nil {
		return errors.Wrap(err, "s.notifier.NotifyChannelIntegrationSuccess")
	}

	return nil
}

func (s *TelegramService) handleUpdateTypeStartCommand(ctx context.Context, update tgbotapi.Update) error {
	// Send a button with the link to the mini app

	var greetingMessage = tgbotapi.EscapeText(
		tgbotapi.ModeMarkdownV2, `
Добро пожаловать в Shoppigram!  
  
С нами вы можете создать свой интернет-магазин в Telegram всего за несколько кликов. Никаких сложностей и технических навыков – всё максимально просто и интуитивно.  
  
✨ Что вы можете сделать с Shoppigram:  
- Создать витрину товаров.  
- Управлять ассортиментом.  
- Обрабатывать заказы.  
- Взаимодействовать с клиентами напрямую через Telegram.  `) +
		fmt.Sprintf(`
📌 Как это может выглядеть:  
[Магазин кроссовок](https://t.me/%s/sneakerboss) 
  
[Кофейня](https://t.me/%s/mycoffe)  

🛠 [Связаться с поддержкой](https://t.me/ShoppigramSupport)  
🌟 [Открыть бота](https://t.me/%s/app)

`, s.shoppigramBotName, s.shoppigramBotName, s.shoppigramBotName) +
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, `
Спасибо, что выбрали Shoppigram! Давайте вместе сделаем ваш бизнес ещё успешнее.
`)

	// Send the message to the user
	err := s.notifier.NotifyGreetings(ctx, NotifyGreetingsRequest{
		UserExternalID:  update.Message.From.ID,
		GreetingMessage: greetingMessage,
	})
	if err != nil {
		return errors.Wrap(err, "s.notifier.NotifyGreetings")
	}

	return nil
}
