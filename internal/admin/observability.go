package admin

import (
	"context"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
	"go.uber.org/zap"
	"strconv"
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

// GetShops calls the underlying service's GetShops method
func (s *ServiceWithObservability) GetShops(ctx context.Context, request GetShopsRequest) (GetShopsResponse, error) {
	res, err := s.service.GetShops(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.GetShops", logger.SilentError(err))
	}

	return res, err
}

// CreateShop calls the underlying service's CreateShop method
func (s *ServiceWithObservability) CreateShop(ctx context.Context, request CreateShopRequest) (CreateShopResponse, error) {
	res, err := s.service.CreateShop(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.CreateShop", logger.SilentError(err))
	}

	return res, err
}

// UpdateShop calls the underlying service's UpdateShop method
func (s *ServiceWithObservability) UpdateShop(ctx context.Context, request UpdateShopRequest) error {
	err := s.service.UpdateShop(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			With(zap.String("web_app_id", request.ID.String())).
			Error("s.service.UpdateShop", logger.SilentError(err))
	}

	return err
}

// DeleteShop calls the underlying service's DeleteShop method
func (s *ServiceWithObservability) DeleteShop(ctx context.Context, request DeleteShopRequest) error {
	err := s.service.DeleteShop(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			With(zap.String("web_app_id", request.WebAppId.String())).
			Error("s.service.SoftDeleteShop", logger.SilentError(err))
	}

	return err
}

// CreateProduct calls the underlying service's CreateProduct method
func (s *ServiceWithObservability) CreateProduct(ctx context.Context, request CreateProductRequest) (CreateProductResponse, error) {
	res, err := s.service.CreateProduct(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateProduct", logger.SilentError(err))
	}

	return res, err
}

// UpdateProduct calls the underlying service's UpdateProduct method
func (s *ServiceWithObservability) UpdateProduct(ctx context.Context, request UpdateProductRequest) error {
	err := s.service.UpdateProduct(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.UpdateProduct", logger.SilentError(err))
	}

	return err
}

// DeleteProduct calls the underlying service's DeleteProduct method
func (s *ServiceWithObservability) DeleteProduct(ctx context.Context, request DeleteProductRequest) error {
	err := s.service.DeleteProduct(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.DeleteProduct", logger.SilentError(err))
	}

	return err
}

// GetOrders calls the underlying service's GetOrders method
func (s *ServiceWithObservability) GetOrders(ctx context.Context, request GetOrdersRequest) (GetOrdersResponse, error) {
	res, err := s.service.GetOrders(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("owner_external_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.GetOrders", logger.SilentError(err))
	}

	return res, err
}

// GetBalance calls the underlying service's GetBalance method
func (s *ServiceWithObservability) GetBalance(ctx context.Context, request GetBalanceRequest) (GetBalanceResponse, error) {
	res, err := s.service.GetBalance(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("owner_external_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.GetBalance", logger.SilentError(err))
	}

	return res, err
}

// CreateProductImageUploadURL calls the underlying service's CreateProductImageUploadURL method
func (s *ServiceWithObservability) CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error) {
	res, err := s.service.CreateProductImageUploadURL(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateProductImageUploadURL", logger.SilentError(err))
	}

	return res, err
}

// CreateShopLogoUploadURL calls the underlying service's CreateMarketplaceLogoUploadURL method
func (s *ServiceWithObservability) CreateShopLogoUploadURL(ctx context.Context, request CreateShopLogoUploadURLRequest) (CreateShopLogoUploadURLResponse, error) {
	res, err := s.service.CreateShopLogoUploadURL(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateShopLogoUploadURL", logger.SilentError(err))
	}

	return res, err
}

// GetTelegramChannels calls the underlying service's GetTelegramChannels method
func (s *ServiceWithObservability) GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error) {
	res, err := s.service.GetTelegramChannels(ctx, ownerExternalID)
	if err != nil {
		s.log.
			With(zap.String("owner_external_id", strconv.FormatInt(ownerExternalID, 10))).
			Error("s.service.GetTelegramChannels", logger.SilentError(err))
	}

	return res, err
}

// PublishShopBannerToChannel calls the underlying service's PublishMarketplaceBannerToChannel method
func (s *ServiceWithObservability) PublishShopBannerToChannel(ctx context.Context, request PublishShopBannerToChannelRequest) error {
	err := s.service.PublishShopBannerToChannel(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			With(zap.String("channel_id", strconv.FormatInt(request.ExternalChannelID, 10))).
			Error("s.service.PublishShopBannerToChannel", logger.SilentError(err))
	}

	return err
}
