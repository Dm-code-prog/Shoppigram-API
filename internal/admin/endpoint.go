package admin

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
)

// makeGetShopEndpoint constructs a GetShops endpoint wrapping the service.
func makeGetShopEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		v0, err := s.GetShops(ctx, GetShopsRequest{
			ExternalUserID: usr.ExternalId,
		})
		if err != nil {
			return nil, errors.Wrap(err, "s.GetShops")
		}
		return v0, nil
	}
}

// makeCreateShopEndpoint creates a new endpoint for access to
// CreateShop service method
func makeCreateShopEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(CreateShopRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.CreateShop(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateShop")
		}

		return response, nil
	}
}

// makeDeleteShopEndpoint creates a new endpoint for access to
// SoftDeleteShop service method
func makeDeleteShopEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(DeleteShopRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.DeleteShop(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.DeleteProduct")
		}

		return nil, nil
	}
}

// makeUpdateShopEndpoint creates a new endpoint for access to
// UpdateShop service method
func makeUpdateShopEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(UpdateShopRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.UpdateShop(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.UpdateShop")
		}

		return nil, nil
	}
}

// makeCreateProductEndpoint creates a new endpoint for access to
// CreateProduct service method
func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(CreateProductRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.CreateProduct(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateProduct")
		}

		return response, nil
	}
}

// makeEnableShopSyncEndpoint creates a new endpoint for access to
// ConfigureShopSync service method
func makeEnableShopSyncEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(ConfigureShopSyncRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.ConfigureShopSync(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.ConfigureShopSync")
		}

		return nil, nil
	}
}

// makeUpdateProductEndpoint creates a new endpoint for access to
// UpdateProduct service method
//
// Path: PUT /api/v1/private/marketplaces/products/<web_app_id>
func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(UpdateProductRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.UpdateProduct(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.UpdateProduct")
		}

		return nil, nil
	}
}

// makeDeleteProductEndpoint creates a new endpoint for access to
// DeleteProduct service method
//
// Path: DELETE /api/v1/private/marketplaces/products/<web_app_id>
func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(DeleteProductRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.DeleteProduct(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.DeleteProduct")
		}

		return nil, nil
	}
}

// makeGetOrdersEndpoint creates a new endpoint for access to
// GetOrders service method
func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(GetOrdersRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.GetOrders(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.GetOrders")
		}

		return response, nil
	}
}

// makeGetBalanceEndpoint creates a new endpoint for access to
// GetBalance service method
func makeGetBalanceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		var request GetBalanceRequest
		request.ExternalUserID = usr.ExternalId
		response, err := s.GetBalance(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.GetBalance")
		}

		return response, nil
	}
}

// makeCreateProductImageUploadURLEndpoint creates a new endpoint for access to
// CreateProductImageUploadURL service method
func makeCreateProductImageUploadURLEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(CreateProductImageUploadURLRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.CreateProductImageUploadURL(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateUploadURL")
		}

		return response, nil
	}
}

// makeCreateShopLogoUploadURLEndpoint creates a new endpoint for access to
// CreateShopLogoUploadURL service method
func makeCreateShopLogoUploadURLEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(CreateShopLogoUploadURLRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.CreateShopLogoUploadURL(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateUploadURL")
		}

		return response, nil
	}
}

// makePublishShopBannerToChannelEndpoint creates a new endpoint for access to
// PublishShopBannerToChannel service method
func makePublishShopBannerToChannelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(PublishShopBannerToChannelRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.PublishShopBannerToChannel(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.PublishShopBannerToChannel")
		}

		return nil, nil
	}
}

// makeGetTelegramChannelsEndpoint creates a new endpoint for access to
// GetTelegramChannels service method
func makeGetTelegramChannelsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		v0, err := s.GetTelegramChannels(ctx, usr.ExternalId)
		if err != nil {
			return nil, errors.Wrap(err, "s.GetTelegramChannels")
		}
		return v0, nil
	}
}
