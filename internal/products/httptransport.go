package products

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(s *Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndpoint(s),
		decodeGetProductsRequest,
		encodeResponse,
		opts...,
	)

	invalidateProductsCacheHandler := kithttp.NewServer(
		makeInvalidateProductsCacheEndpoint(s),
		decodeInvalidateProductsCacheRequest,
		encodeResponse,
		opts...,
	)

	createOrderHandler := kithttp.NewServer(
		authMw(makeCreateOrderEndpoint(s)),
		decodeCreateOrderRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/products/{web_app_id}", getProductsHandler.ServeHTTP)
	r.Put("/products/{web_app_id}/invalidate", invalidateProductsCacheHandler.ServeHTTP)
	r.Post("/orders/{web_app_id}", createOrderHandler.ServeHTTP)

	return r
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return GetProductsRequest{}, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return GetProductsRequest{}, errors.Wrap(ErrorInvalidWebAppID, "uuid.Parse")
	}

	return GetProductsRequest{
		WebAppID: webAppUUID,
	}, nil
}

func decodeInvalidateProductsCacheRequest(_ context.Context, r *http.Request) (interface{}, error) {
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return InvalidateProductsCacheRequest{}, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return InvalidateProductsCacheRequest{}, errors.Wrap(ErrorInvalidWebAppID, "uuid.Parse")
	}

	return InvalidateProductsCacheRequest{
		WebAppID: webAppUUID,
	}, nil
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

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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
		err = ErrorBadRequest
	case errors.Is(err, telegramusers.ErrorInitDataIsMissing):
		w.WriteHeader(http.StatusBadRequest)
		err = telegramusers.ErrorInitDataIsMissing
	case errors.Is(err, ErrorInvalidWebAppID):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInvalidWebAppID
	case errors.Is(err, ErrorInvalidProductQuantity):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInvalidProductQuantity
	case errors.Is(err, telegramusers.ErrorInitDataIsInvalid):
		w.WriteHeader(http.StatusUnauthorized)
		err = telegramusers.ErrorInitDataIsInvalid
	case errors.Is(err, ErrorProductsNotFound):
		w.WriteHeader(http.StatusNotFound)
		err = ErrorProductsNotFound
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err = ErrorInternal
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
