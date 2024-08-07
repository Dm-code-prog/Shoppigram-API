package admins

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/shoppigram-com/marketplace-api/internal/httputils"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
)

// MakeHandler returns a handler for the admin service.
func MakeHandler(bs Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	getMarketplacesHandler := kithttp.NewServer(
		authMw(makeGetMarketplacesEndpoint(bs)),
		httputils.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	createMarketplaceHandler := kithttp.NewServer(
		authMw(makeCreateMarketplaceEndpoint(bs)),
		decodeCreateMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	updateMarketplaceHandler := kithttp.NewServer(
		authMw(makeUpdateMarketplaceEndpoint(bs)),
		decodeUpdateMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	deleteMarketplaceHandler := kithttp.NewServer(
		authMw(makeDeleteMarketplaceEndpoint(bs)),
		decodeDeleteMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	createMarketplaceUploadLogoURLHandler := kithttp.NewServer(
		authMw(makeCreateMarketplaceLogoUploadURLEndpoint(bs)),
		decodeCreateMarketplaceUploadLogoURLRequest,
		encodeResponse,
		opts...,
	)

	createProductHandler := kithttp.NewServer(
		authMw(makeCreateProductEndpoint(bs)),
		decodeCreateProductRequest,
		encodeResponse,
		opts...,
	)

	updateProductHandler := kithttp.NewServer(
		authMw(makeUpdateProductEndpoint(bs)),
		decodeUpdateProductRequest,
		encodeResponse,
		opts...,
	)

	deleteProductHandler := kithttp.NewServer(
		authMw(makeDeleteProductEndpoint(bs)),
		decodeDeleteProductRequest,
		encodeResponse,
		opts...,
	)

	getOrdersHandler := kithttp.NewServer(
		authMw(makeGetOrdersEndpoint(bs)),
		decodeGetOrdersRequest,
		encodeResponse,
		opts...,
	)

	createProductImageUploadURL := kithttp.NewServer(
		authMw(makeCreateProductImageUploadURLEndpoint(bs)),
		decodeCreateProductImageUploadURLRequest,
		encodeResponse,
		opts...,
	)

	publishMarketplaceBannerToChannelHandler := kithttp.NewServer(
		authMw(makePublishMarketplaceBannerToChannelEndpoint(bs)),
		decodePublishMarketplaceBannerToChannelRequest,
		encodeResponse,
		opts...,
	)

	getTelegramChannels := kithttp.NewServer(
		authMw(makeGetTelegramChannelsEndpoint(bs)),
		httputils.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/", getMarketplacesHandler.ServeHTTP)
	r.Post("/", createMarketplaceHandler.ServeHTTP)
	r.Put("/{web_app_id}", updateMarketplaceHandler.ServeHTTP)
	r.Delete("/{web_app_id}", deleteMarketplaceHandler.ServeHTTP)

	r.Post("/products/{web_app_id}", createProductHandler.ServeHTTP)
	r.Put("/products/{web_app_id}", updateProductHandler.ServeHTTP)
	r.Delete("/products/{web_app_id}", deleteProductHandler.ServeHTTP)

	r.Get("/orders", getOrdersHandler.ServeHTTP)

	r.Post("/products/upload-image-url/{web_app_id}", createProductImageUploadURL.ServeHTTP)
	r.Post("/upload-logo-url/{web_app_id}", createMarketplaceUploadLogoURLHandler.ServeHTTP)

	r.Post("/publish-to-channel/{web_app_id}", publishMarketplaceBannerToChannelHandler.ServeHTTP)
	r.Get("/telegram-channels", getTelegramChannels.ServeHTTP)

	return r
}

var badRequestErrors = []error{
	ErrorBadRequest,
	ErrorInvalidShortName,
	ErrorInvalidName,
	ErrorNotUniqueShortName,
	ErrorAdminNotFound,
	ErrorMaxMarketplacesExceeded,
	ErrorInvalidImageExtension,
	ErrorMaxProductsExceeded,
	ErrorInvalidProductCurrency,
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	for _, d := range badRequestErrors {
		if errors.Is(err, d) {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"error": d.Error(),
			})
			return
		}
	}

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

	if errors.Is(err, ErrorOpNotAllowed) {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error": ErrorOpNotAllowed.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": ErrorInternal.Error(),
	})
}
