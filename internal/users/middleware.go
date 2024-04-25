package telegram_users

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

// makeTelegramAuthUserMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func makeTelegramRequestValidationMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var telegramAuthUserRequest TelegramAuthUserRequest

			initData, ok := request.(string)
			if !ok {
				// ASK: Is TelegramAuthUserResponse a correct type to return?
				return TelegramAuthUserResponse{}, ErrorBadRequest
			}

			err := s.TelegramRequestValidation(ctx, initData)
			if err != nil {
				// ASK: Should we have proper logging here?
				return TelegramAuthUserResponse{}, errors.Wrap(err, "s.TelegramRequestValidation")
			}

			if !json.Valid([]byte(initData)) {
				return TelegramAuthUserRequest{}, ErrorInvalidJSON
			}

			err = json.Unmarshal([]byte(initData), &telegramAuthUserRequest)
			if err != nil {
				return TelegramAuthUserRequest{}, ErrorInternal
			}

			// ASK: TelegramAuthUserRequest is not a small structure, maybe we don't need something in it?
			if telegramAuthUserRequest.User.ExternalId == 0 {
				return TelegramAuthUserRequest{}, ErrorBadRequest
			}

			ctx = context.WithValue(ctx, "user", telegramAuthUserRequest.User)

			return next(ctx, request)
		}
	}
}
