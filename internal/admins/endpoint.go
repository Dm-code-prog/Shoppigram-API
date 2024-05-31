package admins

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"go.uber.org/zap"
)

// makeGetMarketplacesEndpoint constructs a GetMarketplaces endpoint wrapping the service.
//
// Path: GET /api/v1/private/marketplaces
func makeGetMarketplacesEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, _ any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
			return nil, err
		}

		v0, err := s.GetMarketplaces(ctx, GetMarketplacesRequest{
			ExternalUserID: usr.ExternalId,
		})
		if err != nil {
			return nil, errors.Wrap(err, "s.GetMarketplaces")
		}
		return v0, nil
	}
}

// makeCreateMarketplaceEndpoint creates a new endpoint for access to
// CreateMarketplace service method
//
// Path: POST /api/v1/private/marketplaces
func makeCreateMarketplaceEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
			return nil, err
		}

		request, ok := req.(CreateMarketplaceRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.CreateMarketplace(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateMarketplace")
		}

		return response, nil
	}
}

// makeUpdateMarketplaceEndpoint creates a new endpoint for access to
// UpdateMarketplace service method
//
// Path: PUT /api/v1/private/marketplaces/<web_app_id>
func makeUpdateMarketplaceEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
			return nil, err
		}

		request, ok := req.(UpdateMarketplaceRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.UpdateMarketplace(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.UpdateMarketplace")
		}

		return nil, nil
	}
}

// makeCreateProductEndpoint creates a new endpoint for access to
// CreateProduct service method
//
// Path: POST /api/v1/private/marketplaces/products/<web_app_id>
func makeCreateProductEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
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

// makeUpdateProductEndpoint creates a new endpoint for access to
// UpdateProduct service method
//
// Path: PUT /api/v1/private/marketplaces/products/<web_app_id>
func makeUpdateProductEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
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
func makeDeleteProductEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
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

// makeCreateProductImageUploadURLEndpoint creates a new endpoint for access to
// CreateProductImageUploadURL service method
//
// Path: POST /api/v1/private/marketplaces/products/upload-image-url/<web_app_id>
func makeCreateProductImageUploadURLEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
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

// makeCreateMarketplaceLogoUploadURLEndpoint creates a new endpoint for access to
// CreateMarketplaceLogoUploadURL service method
//
// Path: POST /api/v1/private/marketplaces/upload-logo-url/<web_app_id>
func makeCreateMarketplaceLogoUploadURLEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
			return nil, err
		}

		request, ok := req.(CreateMarketplaceLogoUploadURLRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		response, err := s.CreateMarketplaceLogoUploadURL(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateUploadURL")
		}

		return response, nil
	}
}

// makePublishMarketplaceBannerToChannelEndpoint creates a new endpoint for access to
// PublishMarketplaceBannerToChannel service method
//
// Path: POST /api/v1/private/marketplaces/publish-to-channel/<web_app_id>
func makePublishMarketplaceBannerToChannelEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
			return nil, err
		}

		request, ok := req.(PublishMarketplaceBannerToChannelRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserID = usr.ExternalId
		err = s.PublishMarketplaceBannerToChannel(ctx, request)
		if err != nil {
			// TODO:
			// Move to observability middleware
			s.log.With(
				zap.String("method", "s.PublishMarketplaceBannerToChannel"),
			).Error(err.Error())
			return nil, errors.Wrap(err, "s.PublishMarketplaceBannerToChannel")
		}

		return nil, nil
	}
}

// makeGetTelegramChannelsEndpoint creates a new endpoint for access to
// GetTelegramChannels service method
//
// Path: GET /api/v1/private/telegram-channels
func makeGetTelegramChannelsEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, _ any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
			).Error(err.Error())
			return nil, err
		}

		v0, err := s.GetTelegramChannels(ctx, usr.ExternalId)
		if err != nil {
			return nil, errors.Wrap(err, "s.GetTelegramChannels")
		}
		return v0, nil
	}
}
