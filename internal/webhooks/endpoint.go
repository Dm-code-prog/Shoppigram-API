package webhooks

import (
	"context"
	"io"

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

func makeCloudPaymentEndpoint(s *CloudPaymentsService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		data, ok := request.(io.ReadCloser)
		if !ok {
			return nil, errors.New("Flailed to cast request to io.ReadCloser")
		}
		resp, err := s.HandleCloudPaymentsWebHook(ctx, data)
		if err != nil && resp != nil {
			return resp, errors.Wrap(err, "Failed to handle CloudPayments webhook")
		}

		return nil, errors.New("Can't handle CloudPayment request Unknown request")
	}
}
