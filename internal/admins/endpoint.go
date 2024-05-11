package admins

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"go.uber.org/zap"
)

// makeGetMarketplaces constructs a GetMarketplaces endpoint wrapping the service.
//
// Path: GET /api/v1/private/marketplaces
func makeGetMarketplaces(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		usr, err := telegramusers.GetUserFromContext(ctx)
		if err != nil {
			s.log.With(
				zap.String("method", "GetUserFromContext"),
				zap.String("external_id", strconv.FormatInt(usr.ExternalId, 10)),
			).Error(err.Error())
			return nil, err
		}

		v0, err := s.GetMarketplaces(ctx, GetMarketplacesRequest{
			ExternalUserID: usr.ExternalId,
		})
		if err != nil {
			s.log.With(
				zap.String("method", "s.GetMarketplaces"),
				zap.String("external_id", strconv.FormatInt(usr.ExternalId, 10)),
			).Error(err.Error())
			return nil, err
		}
		return v0, nil
	}
}
