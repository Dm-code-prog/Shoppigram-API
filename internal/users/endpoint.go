package telegram_users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// makeAuthUserEndpoint constructs a AuthUser endpoint wrapping the service.
//
// Path: PUT /api/v1/public/auth
func makeAuthUserEndpoint(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(AuthUserRequest)
		if !ok {
			return nil, ErrorBadRequest
		}
		v0, err := s.AuthUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return v0, nil
	}
}
