package app

import (
	"context"
	"github.com/google/uuid"
)

type (
	orderType string

	// ProductExternalLink is a link to a product
	// on an external website
	//
	// Used for panels
	ProductExternalLink struct {
		URL string `json:"url"`
	}

	// Product defines the structure for a product
	Product struct {
		ID                  uuid.UUID             `json:"id"`
		Name                string                `json:"name"`
		Description         string                `json:"description,omitempty"`
		Quantity            int32                 `json:"quantity,omitempty"`
		Category            string                `json:"category,omitempty"`
		Price               float64               `json:"price"`
		LegacyPriceCurrency string                `json:"price_currency"`
		ExternalLinks       []ProductExternalLink `json:"external_links"`
	}

	// OrderProduct is a product in an order
	OrderProduct struct {
		ID       uuid.UUID `json:"id"`
		Quantity int32     `json:"quantity"`
	}
)

// Types for requests and responses
type (

	// GetShopRequest defines the request for the GetShops endpoint
	// Products are queried based on the WebAppID
	// or the WebAppShortName
	GetShopRequest struct {
		WebAppID        uuid.UUID
		WebAppShortName string
	}

	// GetShopResponse defines the response body for the GetShops endpoint
	GetShopResponse struct {
		WebAppID              uuid.UUID `json:"web_app_id,omitempty"`
		WebAppName            string    `json:"web_app_name,omitempty"`
		WebAppShortName       string    `json:"web_app_short_name,omitempty"`
		WebAppIsVerified      bool      `json:"web_app_is_verified,omitempty"`
		WebAppType            string    `json:"web_app_type"`
		Currency              string    `json:"currency"`
		OnlinePaymentsEnabled bool      `json:"online_payments_enabled"`
		Products              []Product `json:"products,omitempty"`
	}

	// InvalidateShopCacheRequest defines the request for the InvalidateShopCache endpoint
	InvalidateShopCacheRequest struct {
		WebAppID        uuid.UUID
		WebAppShortName string
	}

	// CreateOrderRequest specifies the products
	// of a web app marketplace that make up
	// the order and user information
	CreateOrderRequest struct {
		WebAppID uuid.UUID
		// p2p or online for now
		Type     orderType      `json:"type"`
		Products []OrderProduct `json:"products"`
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

// Types for repository and service
type (
	// SaveOrderParams is a request to save order info
	// to the storage
	SaveOrderParams struct {
		WebAppID uuid.UUID
		Products []OrderProduct
		// p2p or online for now
		Type           orderType
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
		GetShop(ctx context.Context, request GetShopRequest) (GetShopResponse, error)
		CreateOrder(context.Context, SaveOrderParams) (SaveOrderResult, error)
		GetOrder(ctx context.Context, orderID uuid.UUID, externalUserId int64) (GetOrderResponse, error)
	}

	Service interface {
		GetShop(ctx context.Context, request GetShopRequest) (GetShopResponse, error)
		InvalidateShopCache(ctx context.Context, req InvalidateShopCacheRequest)
		CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error)
		GetOrder(ctx context.Context, req GetOrderRequest) (GetOrderResponse, error)
	}
)
