package notifications

import (
	"encoding/base64"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type (
	// TMALinkingScheme defines the JSON representation of the linking scheme for the Shoppigram TMA
	// Reference:
	// https://www.notion.so/WIP-RFC-TWA-page-linking-and-data-transfer-format-d8ba392b9b19475b80be8aeed415ea30?pvs=4
	TMALinkingScheme struct {
		PageName string `json:"page_name"`
		// We have to use any in here, because of Go constraints
		// The actual type is any "simple" JSON value, like string, number, boolean, or null
		PageData map[string]any `json:"page_data"`
	}
)

// ToBase64String converts the TMALinkingScheme to a base64 string
// That can be decoded and parsed as JSON
func (t TMALinkingScheme) ToBase64String() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func addButtonsToMessage(msg *tgbotapi.MessageConfig, messageData ...telegramButtonData) {
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

func (s *Service) makeMiniAppLink(path string, pageData ...pageDataParam) (string, error) {
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

func makeAddBotAsAdminToChannelLink() string {
	return "https://t.me/" + botName + "?startchannel&admin=post_messages+edit_messages+pin_messages"
}
