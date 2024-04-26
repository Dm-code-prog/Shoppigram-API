package telegram_users

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// User defines the structure for a Marketplace client
	User struct {
		ID           uuid.UUID `json:"id,omitempty"`
		ExternalId   int64     `json:"external_id"`
		IsBot        bool      `json:"is_bot,omitempty"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name,omitempty"`
		Username     string    `json:"username,omitempty"`
		LanguageCode string    `json:"language_code,omitempty"`
		IsPremium    bool      `json:"is_premium,omitempty"`
		AllowsPm     bool      `json:"allows_write_to_pm,omitempty"`
	}

	// CreateOrUpdateTgUserRequest defines the request for the CreateOrUpdateTgUser endpoint
	CreateOrUpdateTgUserRequest User

	// CreateOrUpdateTgUserResponse defines the response for the CreateOrUpdateTgUser endpoint
	CreateOrUpdateTgUserResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// Repository provides access to the user storage
	Repository interface {
		GetEndUserBotToken(ctx context.Context, webAppID uuid.UUID) (string, error)
		CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (uuid.UUID, error)
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

// GetEndUserBotToken gets user bot token
func (s *Service) getEndUserBotToken(ctx context.Context, webAppID uuid.UUID) (string, error) {
	return s.repo.GetEndUserBotToken(ctx, webAppID)
}

// CreateOrUpdateTgUser creates or updates a user record
func (s *Service) CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (CreateOrUpdateTgUserResponse, error) {
	id, err := s.repo.CreateOrUpdateTgUser(ctx, request)
	if err != nil {
		return CreateOrUpdateTgUserResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return CreateOrUpdateTgUserResponse{id}, nil
}
