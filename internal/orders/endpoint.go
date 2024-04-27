package orders

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

// makeGetProductsEndpoint constructs a CreateOrder endpoint wrapping the service.
//
// Path: PUT /api/v1/public/orders/{web_app_id}
func makeCreateOrderEndpoint(svc *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateOrderRequest)
		if !ok {
			return CreateOrderResponse{}, ErrorBadRequest
		}

		res, err := svc.CreateOrder(ctx, req)
		if err != nil {
			return CreateOrderResponse{}, errors.Wrap(err, "svc.CreateOrder()")
		}

		return res, nil
	}
}
