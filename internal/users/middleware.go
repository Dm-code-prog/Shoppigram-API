package telegram_users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"time"
)

type (
	contextKeyT string
)

const (
	userKey     contextKeyT = "user_key"
	initDataKey contextKeyT = "init_data_key"
	webAppIDKey contextKeyT = "web_app_id"

	initDataTTL = time.Second * 30
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
func MakeAuthMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			xInitData, err := GetInitDataFromContext(ctx)
			if err != nil {
				return nil, err
			}

			webAppID, err := GetWebAppIDFromContext(ctx)
			if err != nil {
				return nil, err
			}

			token, err := s.getEndUserBotToken(ctx, webAppID)
			err = initdata.Validate(xInitData, token, initDataTTL)
			if err != nil {
				return nil, err
			}

			parsedInitData, err := initdata.Parse(xInitData)
			if err != nil {
				return nil, err
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
