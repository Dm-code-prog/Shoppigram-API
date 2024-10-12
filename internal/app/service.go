package app

import (
	"context"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/auth"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	// DefaultService implements the Service interface
	// and provides the business logic for shops and orders
	// at the consumer level
	DefaultService struct {
		repo  Repository
		cache *ristretto.Cache[string, GetShopResponse]
	}
)

const (
	getProductsCacheKeyBase = "products.GetShop:"

	orderTypeP2P    orderType = "p2p"
	orderTypeOnline orderType = "online"
)

// New creates a new Service
func New(repo Repository, maxCacheSize int64) *DefaultService {
	cache, err := ristretto.NewCache[string, GetShopResponse](&ristretto.Config[string, GetShopResponse]{
		NumCounters: 100000,
		MaxCost:     maxCacheSize,
		BufferItems: 64,
	})
	if err != nil {
		log.Fatal("failed to create productsCache", logger.SilentError(err))
	}

	return &DefaultService{
		repo:  repo,
		cache: cache,
	}
}

// GetShop returns the products of a marketplace along with the shop info
func (s *DefaultService) GetShop(ctx context.Context, request GetShopRequest) (GetShopResponse, error) {
	cacheKey := makeProductsCacheKey(request.WebAppID, request.WebAppShortName)
	if res, ok := s.cache.Get(cacheKey); ok {
		return res, nil
	}

	res, err := s.repo.GetShop(ctx, request)
	if err != nil {
		return GetShopResponse{}, errors.Wrap(err, "s.repo.GetShop")
	}

	// Cache the response
	s.cache.SetWithTTL(cacheKey, res, 0, 15*time.Minute)

	return res, nil
}

// CreateOrder saves an order to the database
// and notifies the clients, that own the marketplace web app
// about a new order
func (s *DefaultService) CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	u, err := telegramusers.GetUserFromContext(ctx)
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "telegramusers.GetUserFromContext")
	}

	// for backward compatibility
	if req.Type == "" {
		req.Type = orderTypeP2P
	}

	res, err := s.repo.CreateOrder(ctx, SaveOrderParams{
		WebAppID:       req.WebAppID,
		Products:       req.Products,
		ExternalUserID: u.ExternalId,
		Type:           req.Type,
	})
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "s.repo.CreateOrder")
	}
	return CreateOrderResponse(res), nil
}

// GetOrder gets the products in order
func (s *DefaultService) GetOrder(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error) {
	resp, err := s.repo.GetOrder(ctx, req.OrderId, req.ExternalUserId)
	if err != nil {
		return GetOrderResponse{}, errors.Wrap(err, "s.repo.GetOrder")
	}

	return resp, nil
}

// InvalidateShopCache invalidates the cache for the given shop
// The next time GetShop is called with the same web app id, the cache will be missed
// and request will hit the database
func (s *DefaultService) InvalidateShopCache(_ context.Context, req InvalidateShopCacheRequest) {
	s.cache.Del(
		makeProductsCacheKey(req.WebAppID, req.WebAppShortName),
	)
}

// makeProductsCacheKey creates a cache key for the GetShop endpoint
// The key is based on the web app id or the web app short name
func makeProductsCacheKey(webAppID uuid.UUID, webAppShortname string) string {
	if webAppID != uuid.Nil {
		return getProductsCacheKeyBase + webAppID.String()
	}
	return getProductsCacheKeyBase + webAppShortname
}
