package products

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// makeGetProductsEndpoint constructs a GetProducts endpoint wrapping the service.
//
// Path: GET /api/v1/<web_app_id>/products
func makeGetProductsEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductsRequest)
		v0, err := s.GetProducts(ctx, req)
		if err != nil {
			return GetProductsResponse{}, err
		}
		return v0, nil
	}
}
