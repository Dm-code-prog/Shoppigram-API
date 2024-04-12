package products

import (
	"context"
	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type (
	// Product defines the structure for a Marketplace product
	Product struct {
		ID            uuid.UUID `json:"id"`
		WebAppID      uuid.UUID
		Name          string  `json:"name"`
		Description   string  `json:"description,omitempty"`
		Price         float64 `json:"price"`
		PriceCurrency string  `json:"price_currency"`
		ImageURL      string  `json:"image_url,omitempty"`
	}

	// GetProductsRequest defines the request body for the GetProducts endpoint
	// Products are queried based on the WebAppID
	GetProductsRequest struct {
		WebAppID uuid.UUID `json:"web_app_id"`
	}

	// GetProductsResponse defines the response body for the GetProducts endpoint
	GetProductsResponse struct {
		WebAppName string    `json:"web_app_name,omitempty"`
		Products   []Product `json:"products,omitempty"`
	}

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
	getProductsCacheKeyBase = "products.GetProducts"
)

var (
	ErrorNotFound        = errors.New("products not found")
	ErrorInternal        = errors.New("internal server error")
	ErrorInvalidWebAppID = errors.New("invalid web app id")
)

// New creates a new product service
func New(repo Repository, log *zap.Logger, cache *ristretto.Cache) *Service {
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
	if s.cache != nil {
		if res, ok := s.cache.Get(key); ok {
			return res.(GetProductsResponse), nil
		} else {
			s.log.With(
				zap.String("web_app_id", request.WebAppID.String()),
			).Info("cache miss")
		}
	}

	res, err := s.repo.GetProducts(ctx, request)
	if err != nil {
		if !errors.Is(err, ErrorNotFound) {
			s.log.With(
				zap.String("method", "s.repo.GetProducts"),
				zap.String("web_app_id", request.WebAppID.String()),
			).Error(err.Error())
		}
		return GetProductsResponse{}, errors.Wrap(err, "s.repo.GetProducts")
	}

	// Cache the response
	if s.cache != nil {
		s.cache.SetWithTTL(key, res, 0, 15*time.Minute)
	}

	return res, nil
}
