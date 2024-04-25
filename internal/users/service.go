package telegram_users

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"
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

	// TelegramAuthUserRequest defines the request for the TelegramAuthUser endpoint
	// According to the https://core.telegram.org/bots/webapps#webappinitdata
	TelegramAuthUserRequest struct {
		// ASK: Do we need Chat, ChatInstance, ChatType and CanSendAfter fields?
		QueryID      string `json:"query_id,omitempty"`
		User         User   `json:"user"`
		ChatType     string `json:"chat_type,omitempty"`
		ChatInstance string `json:"chat_instance,omitempty"`
		CanSendAfter int    `json:"can_send_after,omitempty"`
		AuthDate     int    `json:"auth_date"`
		Hash         string `json:"hash"`
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
	telegramAuthUserCacheKeyBase      = "users.TelegramAuthUser:"
	telegramAuthUserRequestExpireTime = 30 * time.Second
)

var (
	ErrorBadRequest      = errors.New("bad request")
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

// sign performs Telegram payload signing with the specified key
// Payload itself slice of key-value pairs joined with "\n"
func sign(payload string, key string) string {
	skHmac := hmac.New(sha256.New, []byte("WebAppData"))
	skHmac.Write([]byte(key))

	impHmac := hmac.New(sha256.New, skHmac.Sum(nil))
	impHmac.Write([]byte(payload))

	return hex.EncodeToString(impHmac.Sum(nil))
}

// TelegramRequestValidation validates that request came from Telegram
func (s *Service) TelegramRequestValidation(ctx context.Context, initData string) error {
	// TODO: Get token string here
	token := ""

	// Parse passed init data as query string.
	q, err := url.ParseQuery(initData)
	if err != nil {
		// ASK: Shall we log some kind of an identificator here?
		s.log.With(
			zap.String("method", "url.ParseQuery"),
		).Error(ErrorBadRequest.Error())
		return errors.Wrap(err, "url.ParseQuery")
	}

	var (
		// Init data creation time.
		authDate time.Time
		// Init data sign.
		hash string
		// All found key-value pairs.
		pairs = make([]string, 0, len(q))
	)

	// Iterate over all key-value pairs of parsed parameters.
	for k, v := range q {
		// Store found sign.
		if k == "hash" {
			hash = v[0]
			continue
		}
		if k == "auth_date" {
			if i, err := strconv.Atoi(v[0]); err == nil {
				authDate = time.Unix(int64(i), 0)
			}
		}
		// Append new pair.
		pairs = append(pairs, k+"="+v[0])
	}

	// Sign is always required.
	if hash == "" {
		s.log.Error(ErrorSignMissing.Error())
		return ErrorSignMissing
	}

	// In case, auth date is zero, it means, we can not check if parameters
	// are expired.
	if authDate.IsZero() {
		s.log.Error(ErrorAuthDateMissing.Error())
		return ErrorAuthDateMissing
	}

	// Check if init data is expired.
	if authDate.Add(telegramAuthUserRequestExpireTime).Before(time.Now()) {
		s.log.Error(ErrorExpired.Error())
		return ErrorExpired
	}

	// According to docs, we sort all the pairs in alphabetical order.
	sort.Strings(pairs)

	// In case, our sign is not equal to found one, we should throw an error.
	if sign(strings.Join(pairs, "\n"), token) != hash {
		s.log.Error(ErrorSignInvalid.Error())
		return ErrorSignInvalid
	}
	return nil
}

// TelegramAuthUser creates or updates a user record
func (s *Service) TelegramAuthUser(ctx context.Context, request TelegramAuthUserRequest) (TelegramAuthUserResponse, error) {
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
