package telegram_users

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// makeTelegramAuthUserMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func makeTelegramRequestValidationMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var telegramAuthUserRequest TelegramAuthUserRequest

			// TODO: Get token string here
			token := ""

			initData, ok := request.(string)
			if !ok {
				return nil, ErrorBadRequest
			}

			err := initdata.Validate(initData, token, telegramAuthUserRequestExpireTime)
			if err != nil {
				// ASK: Should we have proper logging here?
				return nil, errors.Wrap(err, "s.TelegramRequestValidation")
			}

			if !json.Valid([]byte(initData)) {
				return nil, ErrorInvalidJSON
			}

			err = json.Unmarshal([]byte(initData), &telegramAuthUserRequest)
			if err != nil {
				return nil, ErrorInternal
			}

			if telegramAuthUserRequest.User.ExternalId == 0 {
				return nil, ErrorBadRequest
			}

			ctx = context.WithValue(ctx, "user", telegramAuthUserRequest.User)

			return next(ctx, request)
		}
	}
}
