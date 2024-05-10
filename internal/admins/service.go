package admins

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Marketplace defines the structure for a Marketplace
	Marketplace struct {
		ID         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		LogoURL    string    `json:"logo_url"`
		IsVerified bool      `json:"is_verified"`
	}

	// GetMarketplacesRequest defines the request for the GetMarketplaces endpoint
	GetMarketplacesRequest struct {
		ExternalUserID int64
	}
	// GetMarketplacesResponse defines the response for the GetMarketplaces endpoint
	GetMarketplacesResponse struct {
		Marketplaces []Marketplace `json:"marketplaces"`
	}

	// Repository provides access to the admin storage
	Repository interface {
		GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error)
	}

	// Service provides admin operations
	Service struct {
		repo Repository
		log  *zap.Logger
	}
)

var (
	ErrorInvalidAdminID = errors.New("invalid admin id")
	ErrorAdminNotFound  = errors.New("admin not found")
	ErrorInternal       = errors.New("internal server error")
)

// New creates a new admin service
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

// GetMarketplaces gets all marketplaces created by user
func (s *Service) GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error) {
	marketplaces, err := s.repo.GetMarketplaces(ctx, req)
	if err != nil {
		if !errors.Is(err, ErrorAdminNotFound) {
			s.log.With(
				zap.String("method", "s.repo.GetProducts"),
				zap.String("user_id", strconv.FormatInt(req.ExternalUserID, 10)),
			).Error(err.Error())
		}
		return GetMarketplacesResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return marketplaces, nil
}
