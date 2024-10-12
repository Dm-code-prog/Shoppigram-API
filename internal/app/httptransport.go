package app

import (
	"github.com/go-kit/kit/endpoint"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeShopHandler returns a handler for products endpoints.
func MakeShopHandler(bs Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeErrorForHTTP),
	}

	getShopH := kithttp.NewServer(
		makeGetShopEndpoint(bs),
		decodeGetShopRequest,
		encodeResponse,
		opts...,
	)

	invalidateShopCacheH := kithttp.NewServer(
		makeInvalidateShopCacheEndpoint(bs),
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
		kithttp.ServerErrorEncoder(encodeErrorForHTTP),
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
