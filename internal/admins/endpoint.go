package admins

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

// makeGetMarketplacesByUserID constructs a GetMarketplacesByUserID endpoint wrapping the service.
//
// Path: GET /api/v1/private/marketplaces/{user_id}
func makeGetMarketplacesByUserID(s *Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetMarketplacesByUserIDRequest)
		if !ok {
			return GetMarketplacesByUserIDResponse{}, ErrorInvalidUserID
		}

		v0, err := s.GetMarketplacesByUserID(ctx, req)
		if err != nil {
			s.log.With(
				zap.String("method", "s.GetMarketplacesByUserID"),
				zap.String("external_id", strconv.FormatInt(req.ExternalID, 10)),
			).Error(err.Error())
			return nil, err
		}
		return v0, nil
	}
}
