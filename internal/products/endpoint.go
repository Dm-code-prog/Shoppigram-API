package products

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

// makeGetProductsEndpoint constructs a GetProducts endpoint wrapping the service.
//
// Path: GET /api/v1/public/products/{web_app_id}
func makeGetProductsEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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
func makeInvalidateProductsCacheEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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
func makeCreateOrderEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateOrderRequest)
		if !ok {
			return CreateOrderResponse{}, ErrorBadRequest
		}

		res, err := s.CreateOrder(ctx, req)
		if err != nil {
			return CreateOrderResponse{}, errors.Wrap(err, "svc.CreateOrder")
		}

		return res, nil
	}
}
