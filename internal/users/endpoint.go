package telegram_users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// makeCreateOrUpdateTelegramUserEndpoint constructs a CreateOrUpdateTelegramUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/telegram_auth
func makeCreateOrUpdateTelegramUserEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateOrUpdateTelegramUserRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		v0, err := s.CreateOrUpdateTelegramUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return v0, nil
	}
}
