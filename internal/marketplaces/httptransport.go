package marketplaces

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// MakeProductsHandler returns a handler for products endpoints.
func MakeProductsHandler(bs Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndpoint(bs),
		decodeGetProductsRequest,
		encodeResponse,
		opts...,
	)

	invalidateProductsCacheHandler := kithttp.NewServer(
		makeInvalidateProductsCacheEndpoint(bs),
		decodeInvalidateProductsCacheRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/{web_app_id}", getProductsHandler.ServeHTTP)
	r.Put("/{web_app_id}/invalidate", invalidateProductsCacheHandler.ServeHTTP)

	return r
}

// MakeOrdersHandler returns a handler for orders endpoints.
func MakeOrdersHandler(s Service, authMW endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	ep := makeCreateOrderEndpoint(s)
	ep = authMW(ep)

	createOrderHandler := kithttp.NewServer(
		ep,
		decodeCreateOrderRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Post("/{web_app_id}", createOrderHandler.ServeHTTP)
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
	webAppID := chi.URLParam(r, "web_app_id")
	if webAppID == "" {
		return nil, ErrorInvalidWebAppID
	}

	webAppUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return nil, ErrorInvalidWebAppID
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrorBadRequest
	}

	req.WebAppID = webAppUUID
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
	for _, e := range telegramusers.AuthenticationErrors {
		if errors.Is(err, e) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"error": e.Error(),
			})
			return
		}
	}

	switch {
	case errors.Is(err, ErrorProductsNotFound):
		w.WriteHeader(http.StatusNotFound)
		err = ErrorProductsNotFound
	case errors.Is(err, ErrorInvalidWebAppID):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, ErrorBadRequest):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorBadRequest
	case errors.Is(err, ErrorInvalidProductQuantity):
		w.WriteHeader(http.StatusBadRequest)
		err = ErrorInvalidProductQuantity
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": ErrorInternal.Error(),
	})
}
