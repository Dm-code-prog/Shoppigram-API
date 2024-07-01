package telegram_users

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	// User defines the structure for a Marketplace end user
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
)

type (
	// Repository provides access to the user storage
	Repository interface {
		CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (uuid.UUID, error)
	}

	Service interface {
		CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (CreateOrUpdateTgUserResponse, error)
	}

	// DefaultService provides user operations
	DefaultService struct {
		repo Repository
	}
)

// New creates a new user service
func New(repo Repository) *DefaultService {
	return &DefaultService{
		repo: repo,
	}
}

// CreateOrUpdateTgUser creates or updates a user record
func (s *DefaultService) CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (CreateOrUpdateTgUserResponse, error) {
	id, err := s.repo.CreateOrUpdateTgUser(ctx, request)
	if err != nil {
		return CreateOrUpdateTgUserResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return CreateOrUpdateTgUserResponse{id}, nil
}
