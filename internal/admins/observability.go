package admins

import (
	"context"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
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

// GetMarketplaces calls the underlying service's GetMarketplaces method
func (s *ServiceWithObservability) GetMarketplaces(ctx context.Context, request GetMarketplacesRequest) (GetMarketplacesResponse, error) {
	res, err := s.service.GetMarketplaces(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.GetMarketplaces", logging.SilentError(err))
	}

	return res, err
}

// CreateMarketplace calls the underlying service's CreateMarketplace method
func (s *ServiceWithObservability) CreateMarketplace(ctx context.Context, request CreateMarketplaceRequest) (CreateMarketplaceResponse, error) {
	res, err := s.service.CreateMarketplace(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.CreateMarketplace", logging.SilentError(err))
	}

	return res, err
}

// UpdateMarketplace calls the underlying service's UpdateMarketplace method
func (s *ServiceWithObservability) UpdateMarketplace(ctx context.Context, request UpdateMarketplaceRequest) error {
	err := s.service.UpdateMarketplace(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			With(zap.String("web_app_id", request.ID.String())).
			Error("s.service.UpdateMarketplace", logging.SilentError(err))
	}

	return err
}

// DeleteMarketplace calls the underlying service's DeleteMarketplace method
func (s *ServiceWithObservability) DeleteMarketplace(ctx context.Context, request DeleteMarketplaceRequest) error {
	err := s.service.DeleteMarketplace(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("external_user_id", strconv.FormatInt(request.ExternalUserID, 10))).
			With(zap.String("web_app_id", request.WebAppId.String())).
			Error("s.service.DeleteMarketplace", logging.SilentError(err))
	}

	return err
}

// CreateProduct calls the underlying service's CreateProduct method
func (s *ServiceWithObservability) CreateProduct(ctx context.Context, request CreateProductRequest) (CreateProductResponse, error) {
	res, err := s.service.CreateProduct(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateProduct", logging.SilentError(err))
	}

	return res, err
}

// UpdateProduct calls the underlying service's UpdateProduct method
func (s *ServiceWithObservability) UpdateProduct(ctx context.Context, request UpdateProductRequest) error {
	err := s.service.UpdateProduct(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.UpdateProduct", logging.SilentError(err))
	}

	return err
}

// DeleteProduct calls the underlying service's DeleteProduct method
func (s *ServiceWithObservability) DeleteProduct(ctx context.Context, request DeleteProductRequest) error {
	err := s.service.DeleteProduct(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.DeleteProduct", logging.SilentError(err))
	}

	return err
}

// GetOrders calls the underlying service's GetOrders method
func (s *ServiceWithObservability) GetOrders(ctx context.Context, request GetOrdersRequest) (GetOrdersResponse, error) {
	res, err := s.service.GetOrders(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("owner_external_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.GetOrders", logging.SilentError(err))
	}

	return res, err
}

// GetBalance calls the underlying service's GetBalance method
func (s *ServiceWithObservability) GetBalance(ctx context.Context, request GetBalanceRequest) (GetBalanceResponse, error) {
	res, err := s.service.GetBalance(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("owner_external_id", strconv.FormatInt(request.ExternalUserID, 10))).
			Error("s.service.GetBalance", logging.SilentError(err))
	}

	return res, err
}

// CreateProductImageUploadURL calls the underlying service's CreateProductImageUploadURL method
func (s *ServiceWithObservability) CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error) {
	res, err := s.service.CreateProductImageUploadURL(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateProductImageUploadURL", logging.SilentError(err))
	}

	return res, err
}

// CreateMarketplaceLogoUploadURL calls the underlying service's CreateMarketplaceLogoUploadURL method
func (s *ServiceWithObservability) CreateMarketplaceLogoUploadURL(ctx context.Context, request CreateMarketplaceLogoUploadURLRequest) (CreateMarketplaceLogoUploadURLResponse, error) {
	res, err := s.service.CreateMarketplaceLogoUploadURL(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			Error("s.service.CreateMarketplaceLogoUploadURL", logging.SilentError(err))
	}

	return res, err
}

// GetTelegramChannels calls the underlying service's GetTelegramChannels method
func (s *ServiceWithObservability) GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error) {
	res, err := s.service.GetTelegramChannels(ctx, ownerExternalID)
	if err != nil {
		s.log.
			With(zap.String("owner_external_id", strconv.FormatInt(ownerExternalID, 10))).
			Error("s.service.GetTelegramChannels", logging.SilentError(err))
	}

	return res, err
}

// PublishMarketplaceBannerToChannel calls the underlying service's PublishMarketplaceBannerToChannel method
func (s *ServiceWithObservability) PublishMarketplaceBannerToChannel(ctx context.Context, request PublishMarketplaceBannerToChannelRequest) error {
	err := s.service.PublishMarketplaceBannerToChannel(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("web_app_id", request.WebAppID.String())).
			With(zap.String("channel_id", strconv.FormatInt(request.ExternalChannelID, 10))).
			Error("s.service.PublishMarketplaceBannerToChannel", logging.SilentError(err))
	}

	return err
}
