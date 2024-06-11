package products

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
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
	}

	// Product is a marketplace product
	// that is identified by the ID and quantity
	ProductOrder struct {
		ID       uuid.UUID `json:"id"`
		Quantity int32     `json:"quantity"`
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

	// CreateOrderRequest specifies the products
	// of a web app marketplace that make up
	// the order and user information
	CreateOrderRequest struct {
		WebAppID uuid.UUID
		Products []ProductOrder `json:"products"`
	}

	// CreateOrderResponse returns the ID of the newly created order
	CreateOrderResponse struct {
		ReadableID int `json:"readable_id"`
	}

	// SaveOrderRequest is a request to save order info
	// to the storage
	SaveOrderRequest struct {
		WebAppID       uuid.UUID
		Products       []ProductOrder
		ExternalUserID int
	}

	// SaveOrderResponse is the response to SaveOrderRequest
	//
	// It contains the readable order ID
	SaveOrderResponse struct {
		ReadableID int
	}

	// Repository provides access to the product storage
	Repository interface {
		GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error)
		CreateOrder(context.Context, SaveOrderRequest) (SaveOrderResponse, error)
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
	ErrorBadRequest             = errors.New("the request to create an order is malformed")
	ErrorInvalidWebAppID        = errors.New("invalid web app id")
	ErrorInvalidProductQuantity = errors.New("the product quantity must be greater than zero")
	ErrorProductsNotFound       = errors.New("products not found")
	ErrorInternal               = errors.New("internal server error")
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

// CreateOrder saves an order to the database
// and notifies the clients, that own the marketplace web app
// about a new order
func (s *Service) CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	u, err := telegramusers.GetUserFromContext(ctx)
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "telegramusers.GetUserFromContext")
	}

	res, err := s.repo.CreateOrder(ctx, SaveOrderRequest{
		WebAppID:       req.WebAppID,
		Products:       req.Products,
		ExternalUserID: int(u.ExternalId),
	})
	if err != nil {
		s.log.
			With(zap.String("web_app_id", req.WebAppID.String())).
			Error("repository.CreateOrder", logging.SilentError(err))
		return CreateOrderResponse{}, errors.Wrap(err, "s.repo.CreateOrder")
	}
	return CreateOrderResponse(res), nil
}
