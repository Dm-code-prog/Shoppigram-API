package telegram_users

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// User defines the structure for a Marketplace client
	User struct {
		ID           uuid.UUID `json:"id,omitempty"`
		ExternalId   int       `json:"external_id"`
		IsBot        bool      `json:"is_bot,omitempty"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name,omitempty"`
		Username     string    `json:"username,omitempty"`
		LanguageCode string    `json:"language_code,omitempty"`
		IsPremium    bool      `json:"is_premium,omitempty"`
		AllowsPm     bool      `json:"allows_write_to_pm,omitempty"`
	}

	// TelegramAuthUserRequest defines the request for the TelegramAuthUser endpoint
	TelegramAuthUserRequest struct {
		User User
	}

	// TelegramAuthUserResponse defines the response for the TelegramAuthUser endpoint
	TelegramAuthUserResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// Repository provides access to the user storage
	Repository interface {
		TelegramAuthUser(ctx context.Context, request TelegramAuthUserRequest) (uuid.UUID, error)
	}

	// Service provides user operations
	Service struct {
		repo Repository
		log  *zap.Logger
	}
)

const (
	telegramAuthUserCacheKeyBase = "users.TelegramAuthUser:"
)

var (
	ErrorBadRequest   = errors.New("bad request")
	ErrorUnauthorized = errors.New("unauthorized")
	ErrorInternal     = errors.New("internal server error")
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

// TelegramRequestValidation validates that request came from Telegram
func (s *Service) TelegramRequestValidation(ctx context.Context, request TelegramAuthUserRequest) bool {
	// TODO: Add request validation functionality
	return true
}

// TelegramAuthUser creates or updates a user record
func (s *Service) TelegramAuthUser(ctx context.Context, request TelegramAuthUserRequest) (TelegramAuthUserResponse, error) {
	if !s.TelegramRequestValidation(ctx, request) {
		// ASK: Should it be an error or just a warning?
		s.log.With(
			zap.String("method", "s.TelegramRequestValidation"),
			zap.String("external_id", strconv.Itoa(request.User.ExternalId)),
		).Error(ErrorUnauthorized.Error())
		return TelegramAuthUserResponse{}, errors.Wrap(ErrorUnauthorized, "s.TelegramRequestValidation")
	}

	id, err := s.repo.TelegramAuthUser(ctx, request)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.TelegramAuthUser"),
			zap.String("external_id", strconv.Itoa(request.User.ExternalId)),
		).Error(err.Error())
		return TelegramAuthUserResponse{}, errors.Wrap(err, "s.repo.TelegramAuthUser")
	}

	return TelegramAuthUserResponse{id}, nil
}
