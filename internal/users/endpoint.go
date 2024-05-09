package telegram_users

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// makeCreateOrUpdateTgUserEndpoint constructs a CreateOrUpdateTgUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/auth/telegram
func makeCreateOrUpdateTgUserEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		usr, err := GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
				zap.String("external_id", strconv.FormatInt(usr.ExternalId, 10)),
			).Error(err.Error())
			return nil, err
		}

		v0, err := s.CreateOrUpdateTgUser(ctx, CreateOrUpdateTgUserRequest(usr))
		if err != nil {
			s.log.With(
				zap.String("method", "s.CreateOrUpdateTgUser"),
				zap.String("external_id", strconv.FormatInt(usr.ExternalId, 10)),
			).Error(err.Error())
			return nil, err
		}
		return v0, nil
	}
}
