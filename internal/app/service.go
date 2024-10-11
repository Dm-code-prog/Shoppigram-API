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

	// GetProductsRequest defines the request for the GetShop endpoint
	// Products are queried based on the WebAppID
	// or the WebAppShortName
	GetProductsRequest struct {
		WebAppID        uuid.UUID
		WebAppShortName string
	}

	// GetProductsResponse defines the response body for the GetShop endpoint
	GetProductsResponse struct {
		WebAppID              uuid.UUID `json:"web_app_id,omitempty"`
		WebAppName            string    `json:"web_app_name,omitempty"`
		WebAppShortName       string    `json:"web_app_short_name,omitempty"`
		WebAppIsVerified      bool      `json:"web_app_is_verified,omitempty"`
		Currency              string    `json:"currency"`
		OnlinePaymentsEnabled bool      `json:"online_payments_enabled"`
		Products              []Product `json:"products,omitempty"`
	}

	// InvalidateProductsCacheRequest defines the request for the InvdlidateShopCache endpoint
	InvalidateProductsCacheRequest struct {
		WebAppID        uuid.UUID
		WebAppShortName string
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
	// SaveOrderParams is a request to save order info
	// to the storage
	SaveOrderParams struct {
		WebAppID uuid.UUID
		Products []ProductItem
		// p2p or online for now
		Type           string
		ExternalUserID int64
	}

	// SaveOrderResult is the response to SaveOrderParams
	//
	// It contains the readable order ID
	SaveOrderResult struct {
		ID         uuid.UUID
		ReadableID int
	}

	// Repository provides access to the product storage
	Repository interface {
		GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error)
		CreateOrder(context.Context, SaveOrderParams) (SaveOrderResult, error)
		GetOrder(ctx context.Context, orderID uuid.UUID, externalUserId int64) (GetOrderResponse, error)
	}

	Service interface {
		GetShop(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error)
		InvdlidateShopCache(ctx context.Context, req InvalidateProductsCacheRequest)
		CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error)
		GetOrder(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error)
	}

	// DefaultService provides product operations
	DefaultService struct {
		repo  Repository
		cache *ristretto.Cache[string, GetProductsResponse]
	}
)

const (
	getProductsCacheKeyBase = "products.GetShop:"

	orderTypeP2P    = "p2p"
	orderTypeOnline = "online"
)

// New creates a new product service
func New(repo Repository, maxCacheSize int64) *DefaultService {
	cache, err := ristretto.NewCache[string, GetProductsResponse](&ristretto.Config[string, GetProductsResponse]{
		NumCounters: 1e7,          // number of keys to track frequency of (10M).
		MaxCost:     maxCacheSize, // maximum cost of the cache
		BufferItems: 64,           // number of keys per Get buffer.
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
func (s *DefaultService) GetShop(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error) {
	cacheKey := makeProductsCacheKey(request.WebAppID, request.WebAppShortName)
	if res, ok := s.cache.Get(cacheKey); ok {
		return res, nil
	}

	res, err := s.repo.GetProducts(ctx, request)
	if err != nil {
		return GetProductsResponse{}, errors.Wrap(err, "s.repo.GetShop")
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

// InvdlidateShopCache invalidates the cache for the given shop
// The next time GetProducts is called with the same web app id, the cache will be missed
// and request will hit the database
func (s *DefaultService) InvdlidateShopCache(_ context.Context, req InvalidateProductsCacheRequest) {
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
