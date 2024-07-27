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

// Cloud Paymet Check Responces
const (
	CloudPaymentsCheckResponceCode_success              = 0
	CloudPaymentsCheckResponceCode_wrongInvoiceID       = 10
	CloudPaymentsCheckResponceCode_accountIncorrect     = 11
	CloudPaymentsCheckResponceCode_wrongSum             = 12
	CloudPaymentsCheckResponceCode_cantHandleThePayment = 13
	CloudPaymentsCheckResponceCode_transactionExpired   = 20
)

func makeCloudPaymentCheckEndpoint(s *CloudPaymentsService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(CloudPaymentsCheckRequest)
		if !ok {
			return nil, errors.New("Can't handle CloudPayment check request. Wrong request format")
		}
		resp, err := s.HandleCloudPaymentsCheckWebHook(ctx, data)
		if err != nil {
			return nil, errors.Wrap(err, "s.HandleCloudPaymentsCheckWebHook(ctx, data)")
		}
		return resp, nil
	}
}
