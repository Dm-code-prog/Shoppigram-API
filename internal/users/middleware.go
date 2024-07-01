package telegram_users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	initdata "github.com/telegram-mini-apps/init-data-golang"
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
		return usr, ErrorUserNotFound
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
		return initData, ErrorInitDataNotFound
	}

	if initData == "" {
		return initData, ErrorInitDataIsEmpty
	}

	return initData, nil
}

// PutWebAppIDToContext grabs the Web app ID from the path params and puts it into context
func PutWebAppIDToContext(ctx context.Context, webAppID string) context.Context {
	asUUID, err := uuid.Parse(webAppID)
	if err != nil {
		// ASK: Should we log it?
		return ctx
	}

	ctx = context.WithValue(ctx, webAppIDKey, asUUID)
	return ctx
}

// GetWebAppIDFromContext gets previously added to context Web app ID
func GetWebAppIDFromContext(ctx context.Context) (uuid.UUID, error) {
	webAppID, ok := ctx.Value(webAppIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrorWebAppNotFound
	}

	return webAppID, nil
}

// AuthServerBefore adds all values necessary to authenticate a Telegram user
// to context
//
// Add it to the ServerOption block of every Go kit server that needs authentication
// that ensures that the user actually came from Telegram.
var AuthServerBefore = []kithttp.ServerOption{
	kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
		xInitData := request.Header.Get("X-Init-Data")
		return PutInitDataToContext(ctx, xInitData)
	}),
	kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
		webAppID := chi.URLParam(request, "web_app_id")
		return PutWebAppIDToContext(ctx, webAppID)
	}),
}

// MakeAuthMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func MakeAuthMiddleware(botToken string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			xInitData, err := GetInitDataFromContext(ctx)
			if err != nil {
				return nil, err
			}

			err = initdata.Validate(xInitData, botToken, initDataTTL)
			if err != nil {
				return nil, ErrorInitDataIsInvalid
			}

			parsedInitData, err := initdata.Parse(xInitData)
			if err != nil {
				return nil, ErrorInitDataIsInvalid
			}

			tgUser := parsedInitData.User
			ctx = PutUserToContext(ctx, User{
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
