package telegram_users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// userToContext puts User data into context
func putUserToContext(ctx context.Context, usr User) context.Context {
	// ASK: Shall we put it together with getUserFromContext in one place?
	ctx = context.WithValue(ctx, "user", usr)
	return ctx
}

// makeCreateOrUpdateTgUserMiddleware constructs a middleware which is responsible for
// Telegram user auth.
func makeValidateTgUserMiddleware(s *Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var createOrUpdateTgUserRequest CreateOrUpdateTgUserRequest
			var usr User

			req, ok := request.(*http.Request)
			if !ok {
				return nil, ErrorBadRequest
			}

			initData := req.Header.Get("X-Init-Data")
			// ASK: Will JSON data be a JSON in that case?

			err := json.Unmarshal([]byte(initData), &createOrUpdateTgUserRequest)
			if err != nil {
				return nil, ErrorInvalidJSON
			}

			token, err := s.GetEndUserBotToken(ctx, createOrUpdateTgUserRequest)
			if err != nil {
				return nil, errors.Wrap(err, "s.GetEndUserBotToken")
			}

			err = initdata.Validate(initData, token, createOrUpdateTgUserRequestExpireTime)
			if err != nil {
				// ASK: Should we have proper logging here?
				return nil, errors.Wrap(err, "s.ValidateTgUser")
			}

			if createOrUpdateTgUserRequest.User.ExternalId == 0 {
				return nil, ErrorBadRequest
			}

			ctx = putUserToContext(ctx, usr)

			return next(ctx, request)
		}
	}
}
