package telegram_users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// makeCreateOrUpdateTgUserEndpoint constructs a CreateOrUpdateTgUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/auth/telegram
func makeCreateOrUpdateTgUserEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		user, err := GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		v0, err := s.CreateOrUpdateTgUser(ctx, CreateOrUpdateTgUserRequest(user))
		if err != nil {
			return nil, err
		}
		return v0, nil
	}
}
