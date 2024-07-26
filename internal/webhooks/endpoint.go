package webhooks

import (
	"context"
	"net/http"

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

func makeCloudPaymentEndpoint(s *CloudPaymentsService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(http.Request)
		if !ok {
			return nil, errors.New("Can not handle CloudPayments request")
		}

		err := s.HandleCloudPaymentsWebHook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to handle CloudPayments webhook")
		}

		// handle data (check if InvoiceId equals order id, and if pyed amount is correct)
		// return proper response
		return nil, nil
	}
}
