package telegram_users

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// userToContext puts User data into context
func putUserToContext(ctx context.Context, usr User) context.Context {
	ctx = context.WithValue(ctx, "user", usr)
	return ctx
}

// makeCreateOrUpdateTelegramUserMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func makeValidateTelegramUserMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var createOrUpdateTelegramUserRequest CreateOrUpdateTelegramUserRequest

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

			if !json.Valid([]byte(initData)) {
				return nil, ErrorInvalidJSON
			}

			err = json.Unmarshal([]byte(initData), &createOrUpdateTelegramUserRequest)
			if err != nil {
				return nil, ErrorInternal
			}

			if createOrUpdateTelegramUserRequest.User.ExternalId == 0 {
				return nil, ErrorBadRequest
			}

			ctx = putUserToContext(ctx, createOrUpdateTelegramUserRequest.User)

			return next(ctx, request)
		}
	}
}
