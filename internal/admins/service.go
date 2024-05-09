package admins

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Marketplace defines the structure for a Marketplace
	Marketplace struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		ImageURL string    `json:"image_url"`
	}

	// GetMarketplacesByUserIDResponse defines the response for the GetMarketplacesByUserID endpoint
	GetMarketplacesByUserIDResponse struct {
		Marketplaces []Marketplace `json:"marketplaces"`
	}

	// Repository provides access to the user storage
	Repository interface {
		GetMarketplacesByUserID(ctx context.Context, userID int64) ([]Marketplace, error)
	}

	// Service provides user operations
	Service struct {
		repo Repository
		log  *zap.Logger
	}
)

var (
	ErrorBadRequest        = errors.New("bad request")
	ErrorUserNotFound      = errors.New("user not found")
	ErrorInitDataIsMissing = errors.New("init data is missing, it must be present in x-init-data header")
	ErrorInitDataNotFound  = errors.New("init data not found")
	ErrorInitDataIsInvalid = errors.New("init data is invalid")
	ErrorInitDataIsEmpty   = errors.New("init data is empty")
	ErrorWebAppNotFound    = errors.New("web app id not found")
	ErrorInternal          = errors.New("internal server error")
)

// New creates a new user service
func New(repo Repository, log *zap.Logger) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}

	return &Service{
		repo: repo,
		log:  log,
	}
}

// CreateOrUpdateTgUser creates or updates a user record
func (s *Service) GetMarketplacesByUserID(ctx context.Context, userID int32) (GetMarketplacesByUserIDResponse, error) {
	marketplaces, err := s.repo.GetMarketplacesByUserID(ctx, userID)
	if err != nil {
		return GetMarketplacesByUserIDResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return GetMarketplacesByUserIDResponse{Marketplaces: marketplaces}, nil
}
