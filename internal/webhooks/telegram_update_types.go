package webhooks

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (s *TelegramService) isUpdateTypeAddedToChannel(update tgbotapi.Update) bool {
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

	return true
}

func (s *TelegramService) isUpdateTypeRemovedFromChannel(update tgbotapi.Update) bool {
	if update.MyChatMember == nil {
		return false
	}
	event := update.MyChatMember

	if event.Chat.Type != "channel" {
		return false
	}

	if event.NewChatMember.User.ID != s.shoppigramBotID {
		return false
	}

	if event.NewChatMember.Status == "administrator" {
		return false
	}

	return true
}

func (s *TelegramService) isUpdateTypeStartCommand(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}

	if update.Message.Text == "/start" {
		return true
	}

	return false
}
