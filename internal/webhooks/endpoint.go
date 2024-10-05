package webhooks

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

func makeTelegramWebhookEndpoint(s *TelegramService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(tgbotapi.Update)
		if !ok {
			return nil, ErrorBadRequest
		}

		err := s.HandleTelegramWebhook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(err, "s.HandleTelegramWebhook")
		}

		return nil, nil
	}
}

func makeCloudPaymentCheckEndpoint(s *CloudPaymentsService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(CloudPaymentsCheckRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		resp, err := s.HandleCloudPaymentsCheckWebHook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(err, "s.HandleCloudPaymentsCheckWebHook")
		}
		return resp, nil
	}
}

func makeCloudPaymentPayEndpoint(s *CloudPaymentsService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(CloudPaymentsPayRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		resp, err := s.HandleCloudPaymentsPayWebHook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(err, "s.HandleCloudPaymentsPayWebHook")
		}
		return resp, nil
	}
}
