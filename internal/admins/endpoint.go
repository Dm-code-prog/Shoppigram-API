package admins

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// makeGetMarketplaces constructs a GetMarketplaces endpoint wrapping the service.
//
// Path: GET /api/v1/private/marketplaces/{user_id}
func makeGetMarketplaces(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(int64)
		if !ok {
			return GetMarketplacesResponse{}, ErrorInvalidUserID
		}

		v0, err := s.GetMarketplaces(ctx, req)
		if err != nil {
			s.log.With(
				zap.String("method", "s.GetMarketplaces"),
				zap.String("external_id", strconv.FormatInt(req, 10)),
			).Error(err.Error())
			return nil, err
		}
		return v0, nil
	}
}
