package products

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Product defines the structure for a Marketplace product
	Product struct {
		ID            uuid.UUID `json:"id"`
		Name          string    `json:"name"`
		Description   string    `json:"description,omitempty"`
		Category      string    `json:"category,omitempty"`
		Price         float64   `json:"price"`
		PriceCurrency string    `json:"price_currency"`
		ImageURL      string    `json:"image_url,omitempty"`
	}

	// GetProductsRequest defines the request for the GetProducts endpoint
	// Products are queried based on the WebAppID
	GetProductsRequest struct {
		WebAppID uuid.UUID
	}

	// GetProductsResponse defines the response body for the GetProducts endpoint
	GetProductsResponse struct {
		WebAppID         uuid.UUID `json:"web_app_id,omitempty"`
		WebAppName       string    `json:"web_app_name,omitempty"`
		WebAppShortName  string    `json:"web_app_short_name,omitempty"`
		WebAppIsVerified bool      `json:"web_app_is_verified,omitempty"`
		Products         []Product `json:"products,omitempty"`
	}

	// InvalidateProductsCacheRequest defines the request for the InvalidateProductsCache endpoint
	InvalidateProductsCacheRequest struct {
		WebAppID uuid.UUID
	}

	// Repository provides access to the product storage
	Repository interface {
		GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error)
	}

	// Service provides product operations
	Service struct {
		repo  Repository
		log   *zap.Logger
		cache *ristretto.Cache
	}
)

const (
	getProductsCacheKeyBase = "products.GetProducts:"
)

var (
	ErrorProductsNotFound = errors.New("products not found")
	ErrorInternal         = errors.New("internal server error")
	ErrorInvalidWebAppID  = errors.New("invalid web app id")
)

// New creates a new product service
func New(repo Repository, log *zap.Logger, cache *ristretto.Cache) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if cache == nil {
		log.Fatal("cache *ristretto.Cache is nil, fatal")
	}

	return &Service{
		repo:  repo,
		log:   log,
		cache: cache,
	}
}

// GetProducts returns a list of products
func (s *Service) GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error) {
	// Check if the request is cached
	key := getProductsCacheKeyBase + request.WebAppID.String()
	if res, ok := s.cache.Get(key); ok {
		return res.(GetProductsResponse), nil
	}

	res, err := s.repo.GetProducts(ctx, request)
	if err != nil {
		if !errors.Is(err, ErrorProductsNotFound) {
			s.log.With(
				zap.String("method", "s.repo.GetProducts"),
				zap.String("web_app_id", request.WebAppID.String()),
			).Error(err.Error())
		}
		return GetProductsResponse{}, errors.Wrap(err, "s.repo.GetProducts")
	}

	// Cache the response
	s.cache.SetWithTTL(key, res, 0, 15*time.Minute)

	return res, nil
}

// InvalidateProductsCache invalidates the cache for the given web app id
// The next time GetProducts is called with the same web app id, the cache will be missed
// and request will hit the database
func (s *Service) InvalidateProductsCache(ctx context.Context, req InvalidateProductsCacheRequest) {
	key := makeProductsCacheKey(req.WebAppID)
	s.cache.Del(key)
}

func makeProductsCacheKey(webAppID uuid.UUID) string {
	return getProductsCacheKeyBase + webAppID.String()
}
