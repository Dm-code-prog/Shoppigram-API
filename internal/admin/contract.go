package admin

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type (
	shopType string

	// Shop stores the information about a shop application
	Shop struct {
		ID                    uuid.UUID `json:"id"`
		Name                  string    `json:"name"`
		ShortName             string    `json:"short_name"`
		IsVerified            bool      `json:"is_verified"`
		Type                  shopType  `json:"type"`
		Currency              string    `json:"currency"`
		OnlinePaymentsEnabled bool      `json:"online_payments_enabled"`
	}

	// ProductExternalLink is a link to a product
	// on an external website
	//
	// Used for panels
	ProductExternalLink struct {
		URL string `json:"url"`
	}

	// TelegramChannel contains the data about a Telegram channel
	TelegramChannel struct {
		ID         uuid.UUID `json:"id"`
		ExternalID int64     `json:"external_id"`
		Name       string    `json:"name"`
		Title      string    `json:"title"`
	}

	// Balance represents the balance of a user
	// in a currency
	Balance struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	}

	// OrderProduct represents a product in a marketplace
	OrderProduct struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Quantity int       `json:"quantity"`
		Price    float64   `json:"price"`
	}

	// Order represents an order in a shop
	Order struct {
		ID            uuid.UUID      `json:"id"`
		MarketplaceID uuid.UUID      `json:"marketplace_id"`
		State         string         `json:"state"`
		Type          string         `json:"type"`
		CreatedAt     time.Time      `json:"created_at"`
		UpdatedAt     time.Time      `json:"updated_at"`
		ReadableID    int            `json:"readable_id"`
		TotalPrice    float64        `json:"total_price"`
		Currency      string         `json:"currency"`
		BuyerUsername string         `json:"buyer_username"`
		Products      []OrderProduct `json:"products"`
	}
)

// Types for requests and responses
type (
	// GetShopsRequest defines the request for the GetShops endpoint
	GetShopsRequest struct {
		ExternalUserID int64
	}

	// GetShopsResponse defines the response for the GetShops endpoint
	GetShopsResponse struct {
		Shops []Shop `json:"marketplaces"`
	}

	// CreateShopRequest creates a new marketplace
	// for a client with a given name and shortname
	CreateShopRequest struct {
		ShortName      string   `json:"short_name"`
		Name           string   `json:"name"`
		Currency       string   `json:"currency"`
		Type           shopType `json:"type"`
		ExternalUserID int64
	}

	// CreateShopResponse returns the ID of the created marketplace
	CreateShopResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// UpdateShopRequest allows editing the name
	// of the marketplace
	UpdateShopRequest struct {
		ID             uuid.UUID
		Name           string `json:"name"`
		ExternalUserID int64
	}
	// DeleteShopRequest specifies a marketplace that needs to be deleted
	DeleteShopRequest struct {
		WebAppId       uuid.UUID
		ExternalUserID int64
	}

	// CreateProductRequest specifies the information about a product
	CreateProductRequest struct {
		WebAppID       uuid.UUID
		ExternalUserID int64
		Name           string                `json:"name"`
		Description    string                `json:"description"`
		Price          float64               `json:"price"`
		Category       string                `json:"category,omitempty"`
		ExternalLinks  []ProductExternalLink `json:"external_links"`
	}

	// CreateProductResponse returns the ID of the created product
	CreateProductResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// UpdateProductRequest specifies the new information about a product
	// in a marketplace
	UpdateProductRequest struct {
		ID             uuid.UUID `json:"id"`
		WebAppID       uuid.UUID
		ExternalUserID int64
		Name           string                `json:"name"`
		Description    string                `json:"description"`
		Price          float64               `json:"price"`
		Category       string                `json:"category,omitempty"`
		ExternalLinks  []ProductExternalLink `json:"external_links"`
	}

	// DeleteProductRequest specifies a product in a marketplace that needs to be deleted
	DeleteProductRequest struct {
		WebAppID       uuid.UUID
		ID             uuid.UUID `json:"id"`
		ExternalUserID int64
	}

	// CreateProductImageUploadURLRequest specifies the request for creating a new product image upload URL
	// for a product in a marketplace
	//
	// The user will be able to upload an image directly to the DigitalOcean Space
	CreateProductImageUploadURLRequest struct {
		WebAppID       uuid.UUID
		ProductID      uuid.UUID `json:"product_id"`
		Extension      string    `json:"extension"`
		ExternalUserID int64
	}

	// CreateProductImageUploadURLResponse specifies the response for creating a new product image upload URL
	CreateProductImageUploadURLResponse struct {
		UploadURL string `json:"upload_url"`
		Key       string `json:"key"`
	}

	// CreateShopLogoUploadURLRequest specifies the request for creating a new marketplace logo upload URL
	CreateShopLogoUploadURLRequest struct {
		WebAppID       uuid.UUID
		Extension      string `json:"extension"`
		ExternalUserID int64
	}

	// CreateShopLogoUploadURLResponse specifies the response for creating a new marketplace logo upload URL
	CreateShopLogoUploadURLResponse struct {
		UploadURL string `json:"upload_url"`
		Key       string `json:"key"`
	}

	// PublishShopBannerToChannelRequest contains the data about a banner to be published to a Telegram channel
	PublishShopBannerToChannelRequest struct {
		WebAppID          uuid.UUID
		ExternalUserID    int64
		ExternalChannelID int64  `json:"channel_id"`
		Message           string `json:"message"`
		PinMessage        bool   `json:"pin_message"`
	}

	// GetTelegramChannelsResponse contains the data about Telegram channels owned by a specific user
	GetTelegramChannelsResponse struct {
		Channels []TelegramChannel `json:"channels"`
	}

	// GetBalanceRequest is a request to get the balance of a user
	// Which is calculated as the sum of all online orders minus the commission
	GetBalanceRequest struct {
		ExternalUserID int64
	}

	// GetBalanceResponse is a response to GetBalanceRequest
	GetBalanceResponse struct {
		Balances []Balance `json:"balances"`
	}

	// GetTelegramChannelOwnerResponse contains the data about Telegram channels owned by a specific user
	GetTelegramChannelOwnerResponse struct {
		ChatId int64
	}

	// GetOrdersRequest is a filter for getting orders
	GetOrdersRequest struct {
		ExternalUserID int64
		State          string
		MarketplaceID  uuid.UUID
		Limit          int
		Offset         int
	}

	GetOrdersResponse struct {
		Orders []Order `json:"orders"`
	}
)

// Types for internal use
type (
	// AddUserToNewOrderNotificationsParams mirrors a corresponding struct
	// in notifications module to reduce coupling
	AddUserToNewOrderNotificationsParams struct {
		WebAppID    uuid.UUID
		AdminChatID int64
	}

	// SendShopBannerParams is a struct for request params to send a marketplace banner to a Telegram channel
	// with a TWA link button markup
	SendShopBannerParams struct {
		WebAppLink    string
		Message       string
		ChannelChatID int64
	}

	// PinNotificationParams is a struct for request params to pin a message in a Telegram channel
	PinNotificationParams struct {
		ChatID    int64
		MessageID int64
	}

	// DOSpacesConfig holds the credentials for the S3 bucket
	DOSpacesConfig struct {
		Endpoint string
		ID       string
		Secret   string
		Bucket   string
	}
)

// interfaces
type (
	// Repository provides access to the admin storage
	Repository interface {
		GetShops(ctx context.Context, req GetShopsRequest) (GetShopsResponse, error)
		GetShortName(ctx context.Context, id uuid.UUID) (string, error)
		CreateShop(ctx context.Context, req CreateShopRequest) (CreateShopResponse, error)
		UpdateShop(ctx context.Context, req UpdateShopRequest) error
		SoftDeleteShop(ctx context.Context, req DeleteShopRequest) error

		CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest) error
		DeleteProduct(ctx context.Context, req DeleteProductRequest) error

		GetOrders(ctx context.Context, request GetOrdersRequest) (GetOrdersResponse, error)
		GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error)

		IsShopOwner(ctx context.Context, externalUserID int64, webAppID uuid.UUID) (bool, error)
		IsProductOwner(ctx context.Context, externalUserID int64, productID uuid.UUID) (bool, error)
		IsTelegramChannelOwner(ctx context.Context, externalUserID, channelID int64) (bool, error)

		GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error)
	}

	Notifier interface {
		AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsParams) error
		SendMarketplaceBanner(ctx context.Context, req SendShopBannerParams) (messageID int64, err error)
		PinNotification(ctx context.Context, req PinNotificationParams) error
	}

	Service interface {
		GetShops(ctx context.Context, req GetShopsRequest) (GetShopsResponse, error)
		CreateShop(ctx context.Context, req CreateShopRequest) (CreateShopResponse, error)
		UpdateShop(ctx context.Context, req UpdateShopRequest) error
		DeleteShop(ctx context.Context, req DeleteShopRequest) error

		CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest) error
		DeleteProduct(ctx context.Context, req DeleteProductRequest) error

		GetOrders(ctx context.Context, req GetOrdersRequest) (GetOrdersResponse, error)
		GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error)

		CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error)
		CreateShopLogoUploadURL(ctx context.Context, request CreateShopLogoUploadURLRequest) (CreateShopLogoUploadURLResponse, error)

		GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error)
		PublishShopBannerToChannel(ctx context.Context, req PublishShopBannerToChannelRequest) error
	}
)
