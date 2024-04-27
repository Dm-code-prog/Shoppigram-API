package orders

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"go.uber.org/zap"
	"net/http"
)

// MakeHandler returns a handler for the users service.
func MakeHandler(s *Service, us *telegramusers.Service, log *zap.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	ep := makeCreateOrderEndpoint(s)
	ep = telegramusers.MakeAuthMiddleware(us, log)(ep)

	createOrUpdateTgUserHandler := kithttp.NewServer(
		ep,
		decodeCreateOrderRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Post("/{web_app_id}", createOrUpdateTgUserHandler.ServeHTTP)

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if response != nil {
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrorBadRequest):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, telegramusers.ErrorInitDataIsMissing):
		w.WriteHeader(http.StatusBadRequest)
		err = telegramusers.ErrorInitDataIsMissing
	case errors.Is(err, ErrorInvalidProductQuantity):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInvalidProductQuantity
	case errors.Is(err, telegramusers.ErrorInitDataIsInvalid):
		w.WriteHeader(http.StatusUnauthorized)
		err = telegramusers.ErrorInitDataIsInvalid
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err = ErrorInternal
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeCreateOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	webAppId, err := telegramusers.GetWebAppIDFromContext(ctx)
	if err != nil {
		return nil, ErrorBadRequest
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrorBadRequest
	}

	req.WebAppID = webAppId
	return req, nil
}
