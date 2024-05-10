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
		ID      uuid.UUID `json:"id"`
		Name    string    `json:"name"`
		LogoURL string    `json:"logo_url"`
	}

	// GetMarketplacesResponse defines the response for the GetMarketplaces endpoint
	GetMarketplacesResponse struct {
		Marketplaces []Marketplace `json:"marketplaces"`
	}

	// Repository provides access to the user storage
	Repository interface {
		GetMarketplaces(ctx context.Context, userID int64) (GetMarketplacesResponse, error)
	}

	// Service provides user operations
	Service struct {
		repo Repository
		log  *zap.Logger
	}
)

var (
	ErrorInvalidUserID = errors.New("invalid user id")
	ErrorUserNotFound  = errors.New("user not found")
	ErrorInternal      = errors.New("internal server error")
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
func (s *Service) GetMarketplaces(ctx context.Context, userID int64) (GetMarketplacesResponse, error) {
	marketplaces, err := s.repo.GetMarketplaces(ctx, userID)
	if err != nil {
		if !errors.Is(err, ErrorUserNotFound) {
			s.log.With(
				zap.String("method", "s.repo.GetProducts"),
				zap.String("user_id", strconv.FormatInt(userID, 10)),
			).Error(err.Error())
		}
		return GetMarketplacesResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return marketplaces, nil
}
