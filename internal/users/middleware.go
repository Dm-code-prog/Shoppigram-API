package telegram_users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/zap"
	"time"
)

type (
	contextKeyT string
)

const (
	userKey     contextKeyT = "user_key"
	initDataKey contextKeyT = "init_data_key"
	webAppIDKey contextKeyT = "web_app_id"

	initDataTTL = time.Minute * 30
)

// PutUserToContext puts User data into context
func PutUserToContext(ctx context.Context, usr User) context.Context {
	ctx = context.WithValue(ctx, userKey, usr)
	return ctx
}

// GetUserFromContext retrieves the telegram user from context
//
// It will be available if PutUserToContext was previously called
func GetUserFromContext(ctx context.Context) (User, error) {
	usr, ok := ctx.Value(userKey).(User)
	if !ok {
		return usr, errors.New("user not found")
	}

	return usr, nil
}

// PutInitDataToContext puts Telegram initData into context
//
// By convention, it is passed in every request in X-Init-Data HTTP header
func PutInitDataToContext(ctx context.Context, initData string) context.Context {
	ctx = context.WithValue(ctx, initDataKey, initData)
	return ctx
}

// GetInitDataFromContext retrieves Telegram initData from context
//
// It will be available if PutInitDataToContext has been previously called
func GetInitDataFromContext(ctx context.Context) (string, error) {
	initData, ok := ctx.Value(initDataKey).(string)
	if !ok {
		return initData, errors.New("init data not found")
	}

	if initData == "" {
		return initData, errors.New("init data is empty")
	}

	return initData, nil
}

func PutWebAppIDToContext(ctx context.Context, webAppID string) context.Context {
	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		return ctx
	}

	ctx = context.WithValue(ctx, webAppIDKey, asUUID)
	return ctx
}

func GetWebAppIDFromContext(ctx context.Context) (uuid.UUID, error) {
	webAppID, ok := ctx.Value(webAppIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("web app id not found")
	}

	return webAppID, nil
}

// MakeAuthMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func MakeAuthMiddleware(s *Service, log *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			xInitData, err := GetInitDataFromContext(ctx)
			if err != nil {
				return nil, ErrorInitDataIsMissing
			}

			webAppID, err := GetWebAppIDFromContext(ctx)
			if err != nil {
				log.Error("web app id missing in the request context", zap.Error(err))
				return nil, err
			}

			log = log.With(zap.String("web_app_id", webAppID.String()))

			token, err := s.getEndUserBotToken(ctx, webAppID)
			if err != nil {
				log.Error("s.getEndUserBotToken()", zap.Error(err))
				return nil, err
			}
			err = initdata.Validate(xInitData, token, initDataTTL)
			if err != nil {
				log.Error("initdata.Validate()", zap.Error(err))
				return nil, ErrorInitDataIsInvalid
			}

			parsedInitData, err := initdata.Parse(xInitData)
			if err != nil {
				log.Error("initdata.Parse()", zap.Error(err))
				return nil, ErrorInitDataIsInvalid
			}

			tgUser := parsedInitData.User
			ctx = PutUserToContext(ctx, User{
				ID:           uuid.UUID{}, // TODO: Get the actual user ID, or maybe not
				ExternalId:   tgUser.ID,
				IsBot:        tgUser.IsBot,
				FirstName:    tgUser.FirstName,
				LastName:     tgUser.LastName,
				Username:     tgUser.Username,
				LanguageCode: tgUser.LanguageCode,
				IsPremium:    tgUser.IsPremium,
				AllowsPm:     tgUser.AllowsWriteToPm,
			})

			return next(ctx, request)
		}
	}
}
