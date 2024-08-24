package marketplaces

import (
	"context"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Product defines the structure for a Marketplace product
	Product struct {
		ID                  uuid.UUID `json:"id"`
		Name                string    `json:"name"`
		Description         string    `json:"description,omitempty"`
		Quantity            int32     `json:"quantity,omitempty"`
		Category            string    `json:"category,omitempty"`
		Price               float64   `json:"price"`
		LegacyPriceCurrency string    `json:"price_currency"`
	}

	// GetProductsRequest defines the request for the GetProducts endpoint
	// Products are queried based on the WebAppID
	GetProductsRequest struct {
		WebAppID uuid.UUID
	}

	// GetProductsResponse defines the response body for the GetProducts endpoint
	GetProductsResponse struct {
		WebAppID              uuid.UUID `json:"web_app_id,omitempty"`
		WebAppName            string    `json:"web_app_name,omitempty"`
		WebAppShortName       string    `json:"web_app_short_name,omitempty"`
		WebAppIsVerified      bool      `json:"web_app_is_verified,omitempty"`
		Currency              string    `json:"currency"`
		OnlinePaymentsEnabled bool      `json:"online_payments_enabled"`
		Products              []Product `json:"products,omitempty"`
	}

	// InvalidateProductsCacheRequest defines the request for the InvalidateProductsCache endpoint
	InvalidateProductsCacheRequest struct {
		WebAppID uuid.UUID
	}

	// ProductItem is a marketplace product
	// that is identified by the ID and quantity
	ProductItem struct {
		ID       uuid.UUID `json:"id"`
		Quantity int32     `json:"quantity"`
	}

	// CreateOrderRequest specifies the products
	// of a web app marketplace that make up
	// the order and user information
	CreateOrderRequest struct {
		WebAppID uuid.UUID
		// p2p or online for now
		Type     string        `json:"type"`
		Products []ProductItem `json:"products"`
	}

	// CreateOrderResponse returns the ID of the newly created order
	CreateOrderResponse struct {
		ID         uuid.UUID `json:"id"`
		ReadableID int       `json:"readable_id"`
	}

	// SaveOrderRequest is a request to save order info
	// to the storage
	SaveOrderRequest struct {
		WebAppID uuid.UUID
		Products []ProductItem
		// p2p or online for now
		Type           string
		ExternalUserID int
	}

	// SaveOrderResponse is the response to SaveOrderRequest
	//
	// It contains the readable order ID
	SaveOrderResponse struct {
		ID         uuid.UUID
		ReadableID int
	}

	// GetOrderRequest defines the request for the GetOrder endpoint
	GetOrderRequest struct {
		OrderId        uuid.UUID
		ExternalUserId int64
	}

	// GetOrderResponse contains the data about all products in order
	GetOrderResponse struct {
		Products        []Product `json:"products"`
		TotalPrice      float64   `json:"total_price"`
		Currency        string    `json:"currency"`
		WebAppName      string    `json:"web_app_name"`
		WebAppShortName string    `json:"web_app_short_name"`
		ReadableOrderID int       `json:"readable_order_id"`
		SellerUsername  string    `json:"seller_username"`
	}
)

type (
	// Repository provides access to the product storage
	Repository interface {
		GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error)
		CreateOrder(context.Context, SaveOrderRequest) (SaveOrderResponse, error)
		GetOrder(ctx context.Context, orderID uuid.UUID, externalUserId int64) (GetOrderResponse, error)
	}

	Service interface {
		GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error)
		CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error)
		InvalidateProductsCache(ctx context.Context, req InvalidateProductsCacheRequest)
		GetOrder(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error)
	}

	// DefaultService provides product operations
	DefaultService struct {
		repo  Repository
		log   *zap.Logger
		cache *ristretto.Cache
	}
)

const (
	getProductsCacheKeyBase = "products.GetProducts:"

	orderTypeP2P    = "p2p"
	orderTypeOnline = "online"
)

// New creates a new product service
func New(repo Repository, cache *ristretto.Cache) *DefaultService {
	if cache == nil {
		log.Fatal("cache *ristretto.Cache is nil, fatal")
	}

	return &DefaultService{
		repo:  repo,
		cache: cache,
	}
}

// GetProducts returns a list of products
func (s *DefaultService) GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error) {
	// Check if the request is cached
	key := getProductsCacheKeyBase + request.WebAppID.String()
	if res, ok := s.cache.Get(key); ok {
		return res.(GetProductsResponse), nil
	}

	res, err := s.repo.GetProducts(ctx, request)
	if err != nil {
		return GetProductsResponse{}, errors.Wrap(err, "s.repo.GetProducts")
	}

	// Cache the response
	s.cache.SetWithTTL(key, res, 0, 15*time.Minute)

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

	res, err := s.repo.CreateOrder(ctx, SaveOrderRequest{
		WebAppID:       req.WebAppID,
		Products:       req.Products,
		ExternalUserID: int(u.ExternalId),
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

// InvalidateProductsCache invalidates the cache for the given web app id
// The next time GetProducts is called with the same web app id, the cache will be missed
// and request will hit the database
func (s *DefaultService) InvalidateProductsCache(_ context.Context, req InvalidateProductsCacheRequest) {
	key := makeProductsCacheKey(req.WebAppID)
	s.cache.Del(key)
}

func makeProductsCacheKey(webAppID uuid.UUID) string {
	return getProductsCacheKeyBase + webAppID.String()
}
