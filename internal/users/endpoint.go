package telegram_users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// makeTelegramAuthUserEndpoint constructs a TelegramAuthUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/telegram_auth
func makeTelegramAuthUserEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(TelegramAuthUserRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		v0, err := s.TelegramAuthUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return v0, nil
	}
}
