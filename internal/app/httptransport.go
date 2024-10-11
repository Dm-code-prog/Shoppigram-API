package app

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// MakeShopHandler returns a handler for products endpoints.
func MakeShopHandler(bs Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	getShopH := kithttp.NewServer(
		makeGetProductsEndpoint(bs),
		decodeGetShopRequest,
		encodeResponse,
		opts...,
	)

	invalidateShopCacheH := kithttp.NewServer(
		makeInvalidateProductsCacheEndpoint(bs),
		decodeInvalidateShopCacheRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/{id}", getShopH.ServeHTTP)
	r.Put("/{web_app_id}/invalidate", invalidateShopCacheH.ServeHTTP)

	return r
}

// MakeOrdersHandler returns a handler for orders endpoints.
func MakeOrdersHandler(s Service, authMW endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	createOrderHandler := kithttp.NewServer(
		authMW(makeCreateOrderEndpoint(s)),
		decodeCreateOrderRequest,
		encodeResponse,
		opts...,
	)

	getOrdersHandler := kithttp.NewServer(
		authMW(makeGetOrderEndpoint(s)),
		decodeGetOrderRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Post("/{web_app_id}", createOrderHandler.ServeHTTP)
	r.Get("/{order_id}", getOrdersHandler.ServeHTTP)
	return r
}

// decodeGetShopRequest decodes the request for the GetShop endpoint.
// The ID can be either a UUID or a short name. The request is malformed if the ID is missing.
func decodeGetShopRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		webAppID        uuid.UUID
		webAppShortName string
	)

	id := chi.URLParam(r, "id")
	if id == "" {
		return GetProductsRequest{}, ErrorInvalidWebAppID
	}

	webAppID, err := uuid.Parse(id)
	if err != nil {
		webAppShortName = id
	}

	return GetProductsRequest{
		WebAppID:        webAppID,
		WebAppShortName: webAppShortName,
	}, nil
}

func decodeInvalidateShopCacheRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func decodeCreateOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

func decodeGetOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetOrderRequest
	orderId := chi.URLParam(r, "order_id")
	if orderId == "" {
		return nil, ErrorBadRequest
	}

	asUUID, err := uuid.Parse(orderId)
	if err != nil {
		return nil, ErrorBadRequest
	}
	request.OrderId = asUUID

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
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
	case errors.Is(err, ErrorGetOrderNotPremited):
		w.WriteHeader(http.StatusForbidden)
		err = ErrorGetOrderNotPremited
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err = ErrorInternal
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
