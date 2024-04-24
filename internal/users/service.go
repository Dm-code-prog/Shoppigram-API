package telegram_users

import (
	"context"
	"net/http"
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
		ID           uuid.UUID `json:"id"`
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

	// Repository provides access to the user storage
	Repository interface {
		CreateUser(ctx context.Context, request AuthUserRequest) error
		GetUser(ctx context.Context, request AuthUserRequest) (User, error)
		UpdateUser(ctx context.Context, request AuthUserRequest) error
		DeleteUser(ctx context.Context, request AuthUserRequest) error
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
	ErrorNotFound   = errors.New("user not found")
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
func (s *Service) AuthUser(ctx context.Context, request AuthUserRequest) (int, error) {
	externalId := strconv.Itoa(request.User.ExternalId)

	// Check if the request is cached
	key := authUserCacheKeyBase + externalId
	if res, ok := s.cache.Get(key); ok {
		usr, ok2 := res.(User)
		if !ok2 {
			return http.StatusInternalServerError, nil
		}
		if !reflect.DeepEqual(request.User, usr) {
			err := s.repo.UpdateUser(ctx, request)
			if err != nil {
				s.log.With(
					zap.String("method", "s.repo.UpdateUser"),
					zap.String("external_id", externalId),
				).Error(err.Error())
				return http.StatusInternalServerError, errors.Wrap(err, "s.repo.UpdateUser")
			}
		}

		return http.StatusAccepted, nil
	} else {
		s.log.With(
			zap.String("external_id", externalId),
		).Info("cache miss")
	}

	res, err := s.repo.GetUser(ctx, request)
	if err != nil {
		if !errors.Is(err, ErrorNotFound) {
			s.log.With(
				zap.String("method", "s.repo.GetUser"),
				zap.String("external_id", externalId),
			).Error(err.Error())
			return http.StatusInternalServerError, errors.Wrap(err, "s.repo.GetUser")
		} else {
			err := s.repo.UpdateUser(ctx, request)
			if err != nil {
				s.log.With(
					zap.String("method", "s.repo.CreateUser"),
					zap.String("external_id", externalId),
				).Error(err.Error())
				return http.StatusInternalServerError, errors.Wrap(err, "s.repo.CreateUser")
			}
			return http.StatusAccepted, nil
		}
	}

	if !reflect.DeepEqual(request.User, res) {
		err := s.repo.UpdateUser(ctx, request)
		if err != nil {
			s.log.With(
				zap.String("method", "s.repo.AuthUser"),
				zap.String("external_id", externalId),
			).Error(err.Error())
		}

	}

	// Cache the response
	s.cache.SetWithTTL(key, res, 0, 1*time.Hour)

	return http.StatusAccepted, nil
}
