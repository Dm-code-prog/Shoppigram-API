package app

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
)

// makeGetShopEndpoint constructs a GetShop endpoint wrapping the service.
//
// Path: GET /api/v1/public/products/{web_app_id}
func makeGetShopEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(GetShopRequest)
		if !ok {
			return GetShopResponse{}, ErrorInvalidWebAppID
		}
		v0, err := s.GetShop(ctx, req)
		if err != nil {
			return GetShopResponse{}, errors.Wrap(err, "s.GetShop")
		}
		return v0, nil
	}
}

// makeInvalidateShopCacheEndpoint constructs a InvalidateShopCache endpoint wrapping the service.
//
// Path: PUT /api/v1/public/products/{web_app_id}/invalidate
func makeInvalidateShopCacheEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(InvalidateShopCacheRequest)
		if !ok {
			return nil, ErrorInvalidWebAppID
		}
		s.InvalidateShopCache(ctx, req)
		return nil, nil
	}
}

// makeGetShopEndpoint constructs a CreateOrder endpoint wrapping the service.
//
// Path: PUT /api/v1/public/orders/{web_app_id}
func makeCreateOrderEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(CreateOrderRequest)
		if !ok {
			return CreateOrderResponse{}, ErrorBadRequest
		}

		res, err := svc.CreateOrder(ctx, req)
		if err != nil {
			return CreateOrderResponse{}, errors.Wrap(err, "svc.CreateOrder")
		}

		return res, nil
	}
}

// makeGetOrderEndpoint creates a new endpoint for access to
// GetOrder service method
//
// Path: GET /api/v1/public/orders/<order_id>
func makeGetOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		request, ok := req.(GetOrderRequest)
		if !ok {
			return nil, ErrorBadRequest
		}

		request.ExternalUserId = usr.ExternalId

		response, err := s.GetOrder(ctx, request)
		if err != nil {
			return nil, err
		}

		return response, nil
	}
}
