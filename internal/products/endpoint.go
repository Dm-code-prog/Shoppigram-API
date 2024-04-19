package products

import (
	"context"
	"github.com/go-kit/kit/endpoint"
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
			return GetProductsResponse{}, err
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
