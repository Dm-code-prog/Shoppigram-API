package telegram_users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

// makeCreateOrUpdateTgUserEndpoint constructs a CreateOrUpdateTgUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/auth/telegram
func makeCreateOrUpdateTgUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		usr, err := GetUserFromContext(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "GetUserFromContext")
		}

		v0, err := s.CreateOrUpdateTgUser(ctx, CreateOrUpdateTgUserRequest(usr))
		if err != nil {
			return nil, errors.Wrap(err, "s.CreateOrUpdateTgUser")
		}
		return v0, nil
	}
}
