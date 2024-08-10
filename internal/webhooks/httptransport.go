package webhooks

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	kitauth "github.com/go-kit/kit/auth/basic"
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

// MakeTelegramHandler returns a handler for the Telegram webhooks service.
func MakeTelegramHandler(s *TelegramService, log *zap.Logger, secretToken string) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			xTelegramBotApiSecretToken := request.Header.Get("X-Telegram-Bot-Api-Secret-Token")
			return context.WithValue(ctx, "X-Telegram-Bot-Api-Secret-Token", xTelegramBotApiSecretToken)
		}),
		kithttp.ServerErrorHandler(serverErrorLogger{logger: log}),
	}

	authMw := makeTelegramWebhookAuthMiddleware(secretToken)

	ep := authMw(makeTelegramWebhookEndpoint(s))
	handler := kithttp.NewServer(ep, decodeTelegramWebhookRequest, encodeTelegramResponse, opts...)

	r := chi.NewRouter()
	r.Post("/", handler.ServeHTTP)
	return r
}

// MakeCloudPaymentsHandlers returns a handler for the CloudPayments webhooks
func MakeCloudPaymentsHandlers(s *CloudPaymentsService, log *zap.Logger, login string, password string) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerErrorHandler(serverErrorLogger{logger: log}),
		kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			return context.WithValue(ctx, kithttp.ContextKeyRequestAuthorization, request.Header.Get("Authorization"))
		}),
	}

	checkEndpoint := makeCloudPaymentCheckEndpoint(s)
	basicAuthMiddleware := kitauth.AuthMiddleware(login, password, "check")

	checkHandler := kithttp.NewServer(
		basicAuthMiddleware(checkEndpoint),
		decodeCloudPaymentsCheckRequest,
		encodeCloudPaymentsCheckResponse,
		opts...)

	payHandler := kithttp.NewServer(
		basicAuthMiddleware(makeCloudPaymentPayEndpoint(s)),
		decodeCloudPaymentsPayRequest,
		encodeCloudPaymentsPayResponse,
		opts...)

	router := chi.NewRouter()
	router.Post("/check", checkHandler.ServeHTTP)
	router.Post("/pay", payHandler.ServeHTTP)
	return router
}

func makeTelegramWebhookAuthMiddleware(secretToken string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
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
	_, _ = w.Write([]byte(ErrorInternalServerError.Error()))
}

func encodeTelegramResponse(_ context.Context, w http.ResponseWriter, response any) error {
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

func decodeCloudPaymentsCheckRequest(_ context.Context, r *http.Request) (any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Parse the query string
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	// Convert and map to the struct
	var req CloudPaymentsCheckRequest

	if v := values.Get("TransactionId"); v != "" {
		req.TransactionId, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := values.Get("Amount"); v != "" {
		req.Amount, _ = strconv.ParseFloat(v, 64)
	}
	req.Currency = values.Get("Currency")
	req.PaymentAmount = values.Get("PaymentAmount")
	req.PaymentCurrency = values.Get("PaymentCurrency")
	req.DateTime = values.Get("DateTime")

	if v := values.Get("CardId"); v != "" {
		req.CardId = &v
	}
	req.CardFirstSix = values.Get("CardFirstSix")
	req.CardLastFour = values.Get("CardLastFour")
	req.CardType = values.Get("CardType")
	req.CardExpDate = values.Get("CardExpDate")

	if v := values.Get("TestMode"); v != "" {
		req.TestMode = v == "1" || v == "true"
	}
	req.Status = values.Get("Status")
	req.OperationType = values.Get("OperationType")

	if v := values.Get("InvoiceId"); v != "" {
		req.InvoiceId = &v
	}
	if v := values.Get("AccountId"); v != "" {
		req.AccountId = &v
	}
	if v := values.Get("SubscriptionId"); v != "" {
		req.SubscriptionId = &v
	}
	if v := values.Get("TokenRecipient"); v != "" {
		req.TokenRecipient = &v
	}
	if v := values.Get("Name"); v != "" {
		req.Name = &v
	}
	if v := values.Get("Email"); v != "" {
		req.Email = &v
	}
	if v := values.Get("IpAddress"); v != "" {
		req.IpAddress = &v
	}
	if v := values.Get("IpCountry"); v != "" {
		req.IpCountry = &v
	}
	if v := values.Get("IpCity"); v != "" {
		req.IpCity = &v
	}
	if v := values.Get("IpRegion"); v != "" {
		req.IpRegion = &v
	}
	if v := values.Get("IpDistrict"); v != "" {
		req.IpDistrict = &v
	}
	if v := values.Get("IpLatitude"); v != "" {
		req.IpLatitude = &v
	}
	if v := values.Get("IpLongitude"); v != "" {
		req.IpLongitude = &v
	}
	if v := values.Get("Issuer"); v != "" {
		req.Issuer = &v
	}
	if v := values.Get("IssuerBankCountry"); v != "" {
		req.IssuerBankCountry = &v
	}
	if v := values.Get("Description"); v != "" {
		req.Description = &v
	}
	if v := values.Get("CardProduct"); v != "" {
		req.CardProduct = &v
	}
	if v := values.Get("PaymentMethod"); v != "" {
		req.PaymentMethod = &v
	}
	if v := values.Get("Data"); v != "" {
		req.Data = &v
	}

	return req, nil
}

func encodeCloudPaymentsCheckResponse(_ context.Context, w http.ResponseWriter, response any) error {
	castedResponse, ok := response.(CloudPaymentsResponse)
	if !ok {
		return ErrorInternalServerError
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(castedResponse); err != nil {
		return ErrorInternalServerError
	}

	return nil
}

func decodeCloudPaymentsPayRequest(_ context.Context, r *http.Request) (any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Parse the query string
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	// Convert and map to the struct
	var req CloudPaymentsPayRequest

	if v := values.Get("TransactionId"); v != "" {
		req.TransactionId, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := values.Get("Amount"); v != "" {
		req.Amount, _ = strconv.ParseFloat(v, 64)
	}
	req.Currency = values.Get("Currency")
	req.PaymentAmount = values.Get("PaymentAmount")
	req.PaymentCurrency = values.Get("PaymentCurrency")
	req.DateTime = values.Get("DateTime")

	if v := values.Get("CardId"); v != "" {
		req.CardId = &v
	}
	req.CardFirstSix = values.Get("CardFirstSix")
	req.CardLastFour = values.Get("CardLastFour")
	req.CardType = values.Get("CardType")
	req.CardExpDate = values.Get("CardExpDate")

	if v := values.Get("TestMode"); v != "" {
		req.TestMode = v == "1" || v == "true"
	}
	req.Status = values.Get("Status")
	req.OperationType = values.Get("OperationType")
	req.GatewayName = values.Get("GatewayName")

	if v := values.Get("InvoiceId"); v != "" {
		req.InvoiceId = &v
	}
	if v := values.Get("AccountId"); v != "" {
		req.AccountId = &v
	}
	if v := values.Get("SubscriptionId"); v != "" {
		req.SubscriptionId = &v
	}
	if v := values.Get("Name"); v != "" {
		req.Name = &v
	}
	if v := values.Get("Email"); v != "" {
		req.Email = &v
	}
	if v := values.Get("IpAddress"); v != "" {
		req.IpAddress = &v
	}
	if v := values.Get("IpCountry"); v != "" {
		req.IpCountry = &v
	}
	if v := values.Get("IpCity"); v != "" {
		req.IpCity = &v
	}
	if v := values.Get("IpRegion"); v != "" {
		req.IpRegion = &v
	}
	if v := values.Get("IpDistrict"); v != "" {
		req.IpDistrict = &v
	}
	if v := values.Get("IpLatitude"); v != "" {
		req.IpLatitude = &v
	}
	if v := values.Get("IpLongitude"); v != "" {
		req.IpLongitude = &v
	}
	if v := values.Get("Issuer"); v != "" {
		req.Issuer = &v
	}
	if v := values.Get("IssuerBankCountry"); v != "" {
		req.IssuerBankCountry = &v
	}
	if v := values.Get("Description"); v != "" {
		req.Description = &v
	}
	if v := values.Get("AuthCode"); v != "" {
		req.AuthCode = &v
	}
	if v := values.Get("Data"); v != "" {
		req.Data = &v
	}
	if v := values.Get("Token"); v != "" {
		req.Token = &v
	}
	if v := values.Get("TotalFee"); v != "" {
		fee, _ := strconv.ParseFloat(v, 64)
		req.TotalFee = &fee
	}
	if v := values.Get("CardProduct"); v != "" {
		req.CardProduct = &v
	}
	if v := values.Get("PaymentMethod"); v != "" {
		req.PaymentMethod = &v
	}
	if v := values.Get("FallBackScenarioDeclinedTransactionId"); v != "" {
		id, _ := strconv.ParseInt(v, 10, 64)
		req.FallBackScenarioDeclinedTransactionId = &id
	}
	if v := values.Get("Rrn"); v != "" {
		req.Rrn = &v
	}

	return req, nil
}

func encodeCloudPaymentsPayResponse(_ context.Context, w http.ResponseWriter, response any) error {
	castedResponse, ok := response.(CloudPaymentsResponse)
	if !ok {
		return ErrorInternalServerError
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(castedResponse); err != nil {
		return ErrorInternalServerError
	}

	return nil
}
