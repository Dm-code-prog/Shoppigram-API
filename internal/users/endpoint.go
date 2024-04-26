package telegram_users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// makeCreateOrUpdateTgUserEndpoint constructs a CreateOrUpdateTgUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/telegram_auth
func makeCreateOrUpdateTgUserEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateOrUpdateTgUserRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		v0, err := s.CreateOrUpdateTgUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return v0, nil
	}
}
