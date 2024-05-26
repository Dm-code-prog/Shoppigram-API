package webhooks

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func makeTelegramWebhookEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(tgbotapi.Update)
		if !ok {
			return nil, errors.New("invalid request")
		}

		err := s.HandleTelegramWebhook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to handle Telegram webhook")
		}

		return nil, nil
	}
}
