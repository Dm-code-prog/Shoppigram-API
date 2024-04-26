package telegram_users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// userToContext puts User data into context
func putUserToContext(ctx context.Context, usr User) context.Context {
	// ASK: Shall we put it together with getUserFromContext in one place?
	ctx = context.WithValue(ctx, "user", usr)
	return ctx
}

// makeCreateOrUpdateTelegramUserMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func makeValidateTelegramUserMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var usr User

			// TODO: Get token string here
			token := ""

			initData, ok := request.(string)
			if !ok {
				return nil, ErrorBadRequest
			}

			err := initdata.Validate(initData, token, createOrUpdateTelegramUserRequestExpireTime)
			if err != nil {
				// ASK: Should we have proper logging here?
				return nil, errors.Wrap(err, "s.ValidateTelegramUser")
			}

			req, ok := request.(*http.Request)
			if !ok {
				return nil, ErrorBadRequest
			}

			usrRaw := chi.URLParam(req, "user")

			err = json.Unmarshal([]byte(usrRaw), &usr)
			if err != nil {
				return nil, ErrorInvalidJSON
			}

			if usr.ExternalId == 0 {
				return nil, ErrorBadRequest
			}

			ctx = putUserToContext(ctx, usr)

			return next(ctx, request)
		}
	}
}
