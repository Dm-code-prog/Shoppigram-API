/**
This file implements the logic for sending notifications
about channel events, to be precise:
	- When our bot was successfully integrated into a channel
	- When our bot failed to be integrated into a channel
	- When our bot was removed from a channel
*/

package notifications

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

// NotifyChannelIntegrationSuccess notifies a user about a successful
// channel integration with Shoppigram
func (s *Service) NotifyChannelIntegrationSuccess(_ context.Context, request NotifyChannelIntegrationSuccessRequest) error {
	message := ChannelIntegrationSuccessNotification(request)

	langCode := checkAndGetLangCode(message.UserLanguage)
	msgTxt, err := message.BuildMessage(langCode)
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	addButtonsToMessage(
		&msg,
		telegramButtonData{
			getTranslation(langCode, "try-new-features"),
			"https://t.me/" + s.botName + "/app",
		},
	)

	_, err = s.SendMessage(msg)
	return err
}

// NotifyChannelIntegrationFailure notifies a user about a failure
// happened during channel integration with Shoppigram
func (s *Service) NotifyChannelIntegrationFailure(_ context.Context, request NotifyChannelIntegrationFailureRequest) error {
	message := ChannelIntegrationFailureNotification(request)

	langCode := checkAndGetLangCode(message.UserLanguage)
	msgTxt, err := message.BuildMessage(langCode)
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	addButtonsToMessage(
		&msg,
		telegramButtonData{
			getTranslation(langCode, "try-again"),
			makeAddBotAsAdminToChannelLink(),
		},
	)

	_, err = s.SendMessage(msg)
	return err
}

// NotifyBotRemovedFromChannel notifies a user about the fact that
// our bot was removed from a channel
func (s *Service) NotifyBotRemovedFromChannel(_ context.Context, request NotifyBotRemovedFromChannelRequest) error {
	message := BotRemovedFromChannelNotification(request)

	langCode := checkAndGetLangCode(message.UserLanguage)
	msgTxt, err := message.BuildMessage(langCode)
	if err != nil {
		return errors.Wrap(err, "message.BuildMessageShoppigram")
	}

	msg := tgbotapi.NewMessage(request.UserExternalID, msgTxt)
	addButtonsToMessage(
		&msg,
		telegramButtonData{
			makeAddBotAsAdminToChannelLink(),
			getTranslation(langCode, "add-bot-as-admin"),
		},
	)

	_, err = s.SendMessage(msg)
	return err
}
