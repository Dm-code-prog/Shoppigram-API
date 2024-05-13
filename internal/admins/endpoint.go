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
