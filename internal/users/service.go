package telegram_users

import (
	"context"
	"strconv"
	"time"

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

	// CreateOrUpdateTelegramUserRequest defines the request for the CreateOrUpdateTelegramUser endpoint
	// According to the https://core.telegram.org/bots/webapps#webappinitdata
	CreateOrUpdateTelegramUserRequest struct {
		// ASK: Do we need Chat, ChatInstance, ChatType and CanSendAfter fields?
		QueryID      string `json:"query_id,omitempty"`
		User         User   `json:"user"`
		ChatType     string `json:"chat_type,omitempty"`
		ChatInstance string `json:"chat_instance,omitempty"`
		CanSendAfter int    `json:"can_send_after,omitempty"`
		AuthDate     int    `json:"auth_date"`
		Hash         string `json:"hash"`
	}

	// CreateOrUpdateTelegramUserResponse defines the response for the CreateOrUpdateTelegramUser endpoint
	CreateOrUpdateTelegramUserResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// Repository provides access to the user storage
	Repository interface {
		CreateOrUpdateTelegramUser(ctx context.Context, request CreateOrUpdateTelegramUserRequest) (uuid.UUID, error)
	}

	// Service provides user operations
	Service struct {
		repo Repository
		log  *zap.Logger
	}
)

const (
	createOrUpdateTelegramUserRequestExpireTime = 30 * time.Second
)

var (
	ErrorBadRequest      = errors.New("bad request")
	ErrorInvalidJSON     = errors.New("invalid JSON")
	ErrorUnauthorized    = errors.New("unauthorized")
	ErrorSignMissing     = errors.New("request sign is missing")
	ErrorAuthDateMissing = errors.New("request auth date is missing")
	ErrorExpired         = errors.New("request is expired")
	ErrorSignInvalid     = errors.New("request sign is invalid")
	ErrorInternal        = errors.New("internal server error")
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

// CreateOrUpdateTelegramUser creates or updates a user record
func (s *Service) CreateOrUpdateTelegramUser(ctx context.Context, request CreateOrUpdateTelegramUserRequest) (CreateOrUpdateTelegramUserResponse, error) {
	id, err := s.repo.CreateOrUpdateTelegramUser(ctx, request)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.CreateOrUpdateTelegramUser"),
			zap.String("external_id", strconv.Itoa(request.User.ExternalId)),
		).Error(err.Error())
		return CreateOrUpdateTelegramUserResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTelegramUser")
	}

	return CreateOrUpdateTelegramUserResponse{id}, nil
}
