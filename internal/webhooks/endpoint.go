package webhooks

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func makeTelegramWebhookEndpoint(log *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(tgbotapi.Update)
		if !ok {
			return nil, errors.New("invalid request")
		}

		// serialize data as pretty JSON and log it
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal data")
		}

		// Log the formatted JSON using zap.Any to preserve structure
		log.Info("received webhook", zap.Any("webhook_data", json.RawMessage(b)))
		return nil, nil
	}
}
