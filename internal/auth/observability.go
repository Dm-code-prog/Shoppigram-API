package auth

import (
	"context"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
	"go.uber.org/zap"
	"strconv"
)

type (
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

// CreateOrUpdateTgUser calls the underlying service's CreateOrUpdateTgUser method
func (s *ServiceWithObservability) CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (CreateOrUpdateTgUserResponse, error) {
	res, err := s.service.CreateOrUpdateTgUser(ctx, request)
	if err != nil {
		s.log.
			With(zap.String("telegram_user_id", strconv.FormatInt(request.ExternalId, 10))).
			Error("failed to create or update telegram user", logger.SilentError(err))
	}

	return res, err
}
