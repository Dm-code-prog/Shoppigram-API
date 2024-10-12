package app

import (
	"context"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
	"go.uber.org/zap"
)

type (
	// ServiceWithObservability wraps the Service with observability
	ServiceWithObservability struct {
		service Service
		log     *zap.Logger
	}
)

// NewServiceWithObservability returns a new instance of the ServiceWithObservability
func NewServiceWithObservability(service Service, log *zap.Logger) *ServiceWithObservability {
	return &ServiceWithObservability{
		service: service,
		log:     log,
	}
}

// GetShop calls the underlying service's GetShop method
func (s *ServiceWithObservability) GetShop(ctx context.Context, request GetShopRequest) (GetShopResponse, error) {
	res, err := s.service.GetShop(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.GetShop", logger.SilentError(err))
	}

	return res, err
}

// CreateOrder calls the underlying service's CreateOrder method
func (s *ServiceWithObservability) CreateOrder(ctx context.Context, request CreateOrderRequest) (CreateOrderResponse, error) {
	res, err := s.service.CreateOrder(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateOrder", logger.SilentError(err))
	}

	return res, err
}

// InvalidateShopCache calls the underlying service's InvalidateProductsCache method
func (s *ServiceWithObservability) InvalidateShopCache(ctx context.Context, request InvalidateShopCacheRequest) {
	s.service.InvalidateShopCache(ctx, request)
}

// GetOrder gets the products in order
func (s *ServiceWithObservability) GetOrder(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error) {
	resp, err := s.service.GetOrder(ctx, req)
	if err != nil {
		s.log.
			With(zap.String("order_id", req.OrderId.String())).
			With(zap.Int64("external_user_id", req.ExternalUserId)).
			Error("s.service.GetOrder", logger.SilentError(err))
	}

	return resp, nil
}
