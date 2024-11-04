package admin

import (
	"github.com/shoppigram-com/marketplace-api/packages/gokithelper"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
)

// MakeHandlerV2 returns an HTTP handler for the admin service.
func MakeHandlerV2(bs Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeErrorHTTP),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	getShopsH := kithttp.NewServer(
		authMw(makeGetShopsEndpoint(bs)),
		gokithelper.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	getShopH := kithttp.NewServer(
		authMw(makeGetShopEndpoint(bs)),
		decodeGetShopRequest,
		encodeResponse,
		opts...,
	)

	createShopH := kithttp.NewServer(
		authMw(makeCreateShopEndpoint(bs)),
		decodeCreateShopRequest,
		encodeResponse,
		opts...,
	)

	updateShopH := kithttp.NewServer(
		authMw(makeUpdateShopEndpoint(bs)),
		decodeUpdateShopRequest,
		encodeResponse,
		opts...,
	)

	deleteShopH := kithttp.NewServer(
		authMw(makeDeleteShopEndpoint(bs)),
		decodeDeleteShopRequest,
		encodeResponse,
		opts...,
	)

	enableShopSyncH := kithttp.NewServer(
		authMw(makeEnableShopSyncEndpoint(bs)),
		decodeEnableShopSyncRequest,
		encodeResponse,
		opts...,
	)

	createShopLogoUploadURLH := kithttp.NewServer(
		authMw(makeCreateShopLogoUploadURLEndpoint(bs)),
		decodeCreateMarketplaceUploadLogoURLRequest,
		encodeResponse,
		opts...,
	)

	createProductH := kithttp.NewServer(
		authMw(makeCreateProductEndpoint(bs)),
		decodeCreateProductRequest,
		encodeResponse,
		opts...,
	)

	updateProductH := kithttp.NewServer(
		authMw(makeUpdateProductEndpoint(bs)),
		decodeUpdateProductRequest,
		encodeResponse,
		opts...,
	)

	deleteProductH := kithttp.NewServer(
		authMw(makeDeleteProductEndpoint(bs)),
		decodeDeleteProductRequest,
		encodeResponse,
		opts...,
	)

	getOrdersH := kithttp.NewServer(
		authMw(makeGetOrdersEndpoint(bs)),
		decodeGetOrdersRequest,
		encodeResponse,
		opts...,
	)

	createProductImageUploadURLH := kithttp.NewServer(
		authMw(makeCreateProductImageUploadURLEndpoint(bs)),
		decodeCreateProductImageUploadURLRequest,
		encodeResponse,
		opts...,
	)

	publishShopPostToChannelH := kithttp.NewServer(
		authMw(makePublishShopBannerToChannelEndpoint(bs)),
		decodePublishMarketplaceBannerToChannelRequest,
		encodeResponse,
		opts...,
	)

	getTelegramChannelsH := kithttp.NewServer(
		authMw(makeGetTelegramChannelsEndpoint(bs)),
		gokithelper.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	getBalanceH := kithttp.NewServer(
		authMw(makeGetBalanceEndpoint(bs)),
		gokithelper.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	r := chi.NewRouter()
	r.Get("/shops", getShopsH.ServeHTTP)
	r.Get("/shops/{web_app_id}", getShopH.ServeHTTP)
	r.Post("/shops", createShopH.ServeHTTP)
	r.Put("/shops/{web_app_id}", updateShopH.ServeHTTP)
	r.Delete("/shops/{web_app_id}", deleteShopH.ServeHTTP)
	r.Put("/shops/sync/{web_app_id}", enableShopSyncH.ServeHTTP)

	r.Post("/shops/products/{web_app_id}", createProductH.ServeHTTP)
	r.Put("/shops/products/{web_app_id}", updateProductH.ServeHTTP)
	r.Delete("/shops/products/{web_app_id}", deleteProductH.ServeHTTP)

	r.Get("/orders", getOrdersH.ServeHTTP)
	r.Get("/balance", getBalanceH.ServeHTTP)

	r.Post("/shops/products/upload-image-url/{web_app_id}", createProductImageUploadURLH.ServeHTTP)
	r.Post("/shops/upload-logo-url/{web_app_id}", createShopLogoUploadURLH.ServeHTTP)

	r.Post("/shops/publish-to-channel/{web_app_id}", publishShopPostToChannelH.ServeHTTP)
	r.Get("/telegram-channels", getTelegramChannelsH.ServeHTTP)

	return r
}
