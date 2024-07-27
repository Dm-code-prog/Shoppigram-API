package webhooks

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"go.uber.org/zap"
)

type serverErrorLogger struct {
	logger *zap.Logger
}

func (s serverErrorLogger) Handle(ctx context.Context, err error) {
	s.logger.Error("server error", logging.SilentError(err))
}

// MakeHandler returns a handler for the Telegram webhooks service.
func MakeHandler(s *Service, log *zap.Logger, secretToken string) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			xTelegramBotApiSecretToken := request.Header.Get("X-Telegram-Bot-Api-Secret-Token")
			return context.WithValue(ctx, "X-Telegram-Bot-Api-Secret-Token", xTelegramBotApiSecretToken)
		}),
		kithttp.ServerErrorHandler(serverErrorLogger{logger: log}),
	}

	authMw := makeWebhookAuthMiddleware(secretToken)

	ep := authMw(makeTelegramWebhookEndpoint(s))
	handler := kithttp.NewServer(ep, decodeTelegramWebhookRequest, encodeResponse, opts...)

	r := chi.NewRouter()
	r.Post("/", handler.ServeHTTP)
	return r
}

// MakeCloudPaymentsHandler returns a handler for the CloudPayments webhooks
func MakeCloudPaymentsHandler(s *CloudPaymentsService, log *zap.Logger, login string, password string) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerErrorHandler(serverErrorLogger{logger: log}),
	}

	ep := makeCloudPaymentEndpoint(s)
	handler := kithttp.NewServer(ep, decodeCloudPaymentsRequest, encodeResponse, opts...)

	r := chi.NewRouter()
	r.Post("/", handler.ServeHTTP)
	return r
}

func makeWebhookAuthMiddleware(secretToken string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			xTelegramBotApiSecretToken := ctx.Value("X-Telegram-Bot-Api-Secret-Token").(string)
			if xTelegramBotApiSecretToken != secretToken {
				return nil, errors.New("invalid secret token")
			}
			return next(ctx, request)
		}
	}
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func decodeTelegramWebhookRequest(_ context.Context, r *http.Request) (any, error) {
	var request tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeCloudPaymentsRequest(_ context.Context, r *http.Request) (any, error) {
	rBody := r.Body
	if rBody == nil {
		return nil, errors.New("Can't get request body")
	}
	return rBody, nil
}

func enclodeCloudPaymentsResponce(_ context.Context, w http.ResponseWriter, responce any) error {
	w.WriteHeader(http.StatusOK)

	// w.Write()

	return nil
}
