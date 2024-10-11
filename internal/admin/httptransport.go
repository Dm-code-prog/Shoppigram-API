package admin

import (
	"github.com/shoppigram-com/marketplace-api/packages/gokithelper"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
)

// MakeHandler returns a handler for the admin service.
func MakeHandler(bs Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeErrorHTTP),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	getMarketplacesHandler := kithttp.NewServer(
		authMw(makeGetMarketplacesEndpoint(bs)),
		gokithelper.DecodeEmptyRequest,
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
		gokithelper.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	getBalanceHandler := kithttp.NewServer(
		authMw(makeGetBalanceEndpoint(bs)),
		gokithelper.DecodeEmptyRequest,
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
	r.Get("/balance", getBalanceHandler.ServeHTTP)

	r.Post("/products/upload-image-url/{web_app_id}", createProductImageUploadURL.ServeHTTP)
	r.Post("/upload-logo-url/{web_app_id}", createMarketplaceUploadLogoURLHandler.ServeHTTP)

	r.Post("/publish-to-channel/{web_app_id}", publishMarketplaceBannerToChannelHandler.ServeHTTP)
	r.Get("/telegram-channels", getTelegramChannels.ServeHTTP)

	return r
}

// MakeHandlerV2 returns an HTTP handler for the admin service.
func MakeHandlerV2(bs Service, authMw endpoint.Middleware) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeErrorHTTP),
	}
	opts = append(opts, telegramusers.AuthServerBefore...)

	getShopsH := kithttp.NewServer(
		authMw(makeGetMarketplacesEndpoint(bs)),
		gokithelper.DecodeEmptyRequest,
		encodeResponse,
		opts...,
	)

	createShopH := kithttp.NewServer(
		authMw(makeCreateMarketplaceEndpoint(bs)),
		decodeCreateMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	updateShopH := kithttp.NewServer(
		authMw(makeUpdateMarketplaceEndpoint(bs)),
		decodeUpdateMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	deleteShopH := kithttp.NewServer(
		authMw(makeDeleteMarketplaceEndpoint(bs)),
		decodeDeleteMarketplaceRequest,
		encodeResponse,
		opts...,
	)

	createShopLogoUploadURLH := kithttp.NewServer(
		authMw(makeCreateMarketplaceLogoUploadURLEndpoint(bs)),
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
		authMw(makePublishMarketplaceBannerToChannelEndpoint(bs)),
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
	r.Post("/shops", createShopH.ServeHTTP)
	r.Put("/shops/{web_app_id}", updateShopH.ServeHTTP)
	r.Delete("/shops/{web_app_id}", deleteShopH.ServeHTTP)

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
