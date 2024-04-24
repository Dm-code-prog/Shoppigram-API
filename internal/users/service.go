package telegram_users

import (
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/dgraph-io/ristretto"
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
		AllowsPm     bool      `json:"allows_pm,omitempty"`
	}

	// AuthUserRequest defines the request for the AuthUser endpoint
	AuthUserRequest struct {
		User User `json:"user"`
	}

	// AuthUserResponse defines the response for the AuthUser endpoint
	AuthUserResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// Repository provides access to the user storage
	Repository interface {
		AuthUser(ctx context.Context, request AuthUserRequest) error
	}

	// Service provides user operations
	Service struct {
		repo  Repository
		log   *zap.Logger
		cache *ristretto.Cache
	}
)

const (
	authUserCacheKeyBase = "users.AuthUser:"
)

var (
	ErrorBadRequest = errors.New("bad request")
	ErrorInternal   = errors.New("internal server error")
)

// New creates a new user service
func New(repo Repository, log *zap.Logger, cache *ristretto.Cache) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if cache == nil {
		log.Fatal("cache *ristretto.Cache is nil, fatal")
	}

	return &Service{
		repo:  repo,
		log:   log,
		cache: cache,
	}
}

// AuthUser checks if a user record is present and if not, or if its
// information is not out of date, then updates it
func (s *Service) AuthUser(ctx context.Context, request AuthUserRequest) (AuthUserResponse, error) {
	externalId := strconv.Itoa(request.User.ExternalId)

	// Check if the request is cached
	key := authUserCacheKeyBase + externalId
	res, ok := s.cache.Get(key)
	if !ok {
		s.log.With(
			zap.String("method", "s.cache.Get"),
			zap.String("external_id", externalId),
		).Info("cache miss")
	}

	usr, ok2 := res.(User)
	if !ok2 {
		s.log.With(
			zap.String("method", "s.cache.Get"),
			zap.String("external_id", externalId),
		).Error("User type cache data casting")
		return AuthUserResponse{}, ErrorInternal
	}

	// FIXME: Most of the time it won't be equal because of the ID field
	if reflect.DeepEqual(request.User, usr) {
		// Cache the response
		s.cache.SetWithTTL(key, request.User, 0, 1*time.Hour)
		return AuthUserResponse{usr.ID}, nil
	}

	err := s.repo.AuthUser(ctx, request)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.AuthUser"),
			zap.String("external_id", externalId),
		).Error(err.Error())
		return AuthUserResponse{}, errors.Wrap(err, "s.repo.AuthUser")
	}

	// Cache the response
	s.cache.SetWithTTL(key, request.User, 0, 1*time.Hour)

	return AuthUserResponse{usr.ID}, nil
}
