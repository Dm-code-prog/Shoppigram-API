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
			return nil, ErrorBadRequest
		}

		err := s.HandleTelegramWebhook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(ErrorCantHandle, "s.HandleTelegramWebhook(ctx, data)")
		}

		return nil, nil
	}
}

// Cloud Payments Check Responses
const (
	cloudPaymentsCheckResponceCodeSuccess              = 0
	cloudPaymentsCheckResponceCodeWrongInvoiceID       = 10
	cloudPaymentsCheckResponceCodeAccountIncorrect     = 11
	cloudPaymentsCheckResponceCodeWrongSum             = 12
	cloudPaymentsCheckResponceCodeCantHandleThePayment = 13
	cloudPaymentsCheckResponceCodeTransactionExpired   = 20
)

func makeCloudPaymentCheckEndpoint(s *CloudPaymentsService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(CloudPaymentsCheckRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		resp, err := s.HandleCloudPaymentsCheckWebHook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(ErrorCantHandle, "s.HandleCloudPaymentsCheckWebHook(ctx, data)")
		}
		return resp, nil
	}
}
