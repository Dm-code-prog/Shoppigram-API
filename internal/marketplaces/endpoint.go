package marketplaces

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
)

// makeGetProductsEndpoint constructs a GetProducts endpoint wrapping the service.
//
// Path: GET /api/v1/public/products/{web_app_id}
func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(GetProductsRequest)
		if !ok {
			return GetProductsResponse{}, ErrorInvalidWebAppID
		}
		v0, err := s.GetProducts(ctx, req)
		if err != nil {
			return GetProductsResponse{}, errors.Wrap(err, "s.GetProducts")
		}
		return v0, nil
	}
}

// makeInvalidateProductsCacheEndpoint constructs a InvalidateProductsCache endpoint wrapping the service.
//
// Path: PUT /api/v1/public/products/{web_app_id}/invalidate
func makeInvalidateProductsCacheEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(InvalidateProductsCacheRequest)
		if !ok {
			return nil, ErrorInvalidWebAppID
		}
		s.InvalidateProductsCache(ctx, req)
		return nil, nil
	}
}

// makeGetProductsEndpoint constructs a CreateOrder endpoint wrapping the service.
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
// Path: GET /api/v1/public/order/<order_id>
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

		responce, err := s.GetOrder(ctx, request)
		if err != nil {
			return nil, err
		}
		return responce, nil
	}
}
