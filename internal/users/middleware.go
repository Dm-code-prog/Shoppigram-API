package telegram_users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// makeTelegramAuthUserMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func makeTelegramRequestValidationMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// TODO: Implement the authentication middleware
			return next(ctx, request)
		}
	}
}
