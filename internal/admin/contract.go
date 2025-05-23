package admin

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type (
	shopType string

	// SyncDetails represents the synchronization status of a shop
	// with an external provider
	SyncDetails struct {
		ExternalProvider string    `json:"external_provider"`
		IsActive         bool      `json:"is_active"`
		LastSyncedAt     time.Time `json:"last_synced_at"`
		LastStatus       string    `json:"last_status"`
	}

	// Shop stores the information about a shop application
	Shop struct {
		ID                    uuid.UUID    `json:"id"`
		Name                  string       `json:"name"`
		ShortName             string       `json:"short_name"`
		IsVerified            bool         `json:"is_verified"`
		Type                  shopType     `json:"type"`
		Currency              string       `json:"currency"`
		OnlinePaymentsEnabled bool         `json:"online_payments_enabled"`
		SyncDetails           *SyncDetails `json:"sync_details,omitempty"`
	}

	// ProductExternalLink is a link to a product
	// on an external website
	//
	// Used for panels
	ProductExternalLink struct {
		URL   string `json:"url"`
		Label string `json:"label"`
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

	// OrderProduct represents a product in a shop
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
		Shops []Shop `json:"shops"`
	}

	// GetShopRequest specifies the information required to get a shop
	GetShopRequest struct {
		ExternalUserID int64
		WebAppID       uuid.UUID
	}

	// GetShopResponse specifies the information returned
	// about a shop
	GetShopResponse Shop

	// CreateShopRequest creates a new shop
	// for a client with a given name and shortname
	CreateShopRequest struct {
		ShortName      string   `json:"short_name"`
		Name           string   `json:"name"`
		Currency       string   `json:"currency"`
		Type           shopType `json:"type"`
		ExternalUserID int64
	}

	// CreateShopResponse returns the ID of the created shop
	CreateShopResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// UpdateShopRequest allows editing the name
	// of the shop
	UpdateShopRequest struct {
		ID             uuid.UUID
		Name           string `json:"name"`
		ExternalUserID int64
	}
	// DeleteShopRequest specifies a shop that needs to be deleted
	DeleteShopRequest struct {
		WebAppId       uuid.UUID
		ExternalUserID int64
	}

	// ConfigureShopSyncRequest specifies the information required to enable
	// the synchronization of a shop with an external provider
	ConfigureShopSyncRequest struct {
		WebAppID         uuid.UUID
		ExternalUserID   int64
		IsActive         bool   `json:"is_active"`
		ExternalProvider string `json:"external_provider"`
		APIKey           string `json:"api_key"`
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
	// in a shop
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

	// DeleteProductRequest specifies a product in a shop that needs to be deleted
	DeleteProductRequest struct {
		WebAppID       uuid.UUID
		ID             uuid.UUID `json:"id"`
		ExternalUserID int64
	}

	// CreateProductImageUploadURLRequest specifies the request for creating a new product image upload URL
	// for a product in a shop
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

	// CreateShopLogoUploadURLRequest specifies the request for creating a new shop logo upload URL
	CreateShopLogoUploadURLRequest struct {
		WebAppID       uuid.UUID
		Extension      string `json:"extension"`
		ExternalUserID int64
	}

	// CreateShopLogoUploadURLResponse specifies the response for creating a new shop logo upload URL
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

	// GetOrdersRequest is a filter for getting orders
	GetOrdersRequest struct {
		ExternalUserID int64
		State          string    `json:"state"`
		ShopID         uuid.UUID `json:"shop_id"`
		Limit          int       `json:"limit"`
		Offset         int       `json:"offset"`
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

	// SendShopBannerParams is a struct for request params to send a shop banner to a Telegram channel
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
)

// interfaces
type (
	// Repository provides access to the admin storage
	Repository interface {
		GetShops(ctx context.Context, req GetShopsRequest) (GetShopsResponse, error)

		GetShop(ctx context.Context, req GetShopRequest) (GetShopResponse, error)

		CreateShop(ctx context.Context, req CreateShopRequest) (CreateShopResponse, error)

		UpdateShop(ctx context.Context, req UpdateShopRequest) error

		SoftDeleteShop(ctx context.Context, req DeleteShopRequest) error

		ConfigureShopSync(ctx context.Context, req ConfigureShopSyncRequest) error

		CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest) error
		DeleteProduct(ctx context.Context, req DeleteProductRequest) error

		GetOrders(ctx context.Context, request GetOrdersRequest) (GetOrdersResponse, error)
		GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error)

		IsShopOwner(ctx context.Context, externalUserID int64, webAppID uuid.UUID) (bool, error)
		IsTelegramChannelOwner(ctx context.Context, externalUserID, channelID int64) (bool, error)

		GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error)
	}

	Notifier interface {
		AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsParams) error
		SendMarketplaceBanner(ctx context.Context, req SendShopBannerParams) (messageID int64, err error)
		PinNotification(ctx context.Context, req PinNotificationParams) error
	}

	Service interface {
		// GetShops returns a list of shops owned by a user
		GetShops(context.Context, GetShopsRequest) (GetShopsResponse, error)

		// GetShop returns the information about a shop
		GetShop(context.Context, GetShopRequest) (GetShopResponse, error)

		// CreateShop creates a new shop for a user
		CreateShop(context.Context, CreateShopRequest) (CreateShopResponse, error)

		// UpdateShop updates a shop
		UpdateShop(context.Context, UpdateShopRequest) error

		// DeleteShop deletes a shop
		DeleteShop(context.Context, DeleteShopRequest) error

		// ConfigureShopSync enables or disables the synchronization of a shop with an external provider
		ConfigureShopSync(context.Context, ConfigureShopSyncRequest) error

		// CreateProduct creates a new product in a shop
		CreateProduct(context.Context, CreateProductRequest) (CreateProductResponse, error)

		// UpdateProduct updates a product in a shop
		UpdateProduct(context.Context, UpdateProductRequest) error

		// DeleteProduct deletes a product in a shop
		DeleteProduct(context.Context, DeleteProductRequest) error

		// GetOrders returns a list of orders for a user
		// can be filtered by shop, state, etc.
		GetOrders(context.Context, GetOrdersRequest) (GetOrdersResponse, error)

		// GetBalance returns the balance of a user
		// across all shops
		//
		// Not production-ready, don't use it
		GetBalance(context.Context, GetBalanceRequest) (GetBalanceResponse, error)

		// CreateProductImageUploadURL returns a URL for uploading a product image
		CreateProductImageUploadURL(context.Context, CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error)

		// CreateShopLogoUploadURL returns a URL for uploading a shop logo
		CreateShopLogoUploadURL(context.Context, CreateShopLogoUploadURLRequest) (CreateShopLogoUploadURLResponse, error)

		// GetTelegramChannels returns a list of Telegram channels owned by a user
		GetTelegramChannels(context.Context, int64) (GetTelegramChannelsResponse, error)

		// PublishShopBannerToChannel publishes a banner to a Telegram channel
		PublishShopBannerToChannel(context.Context, PublishShopBannerToChannelRequest) error
	}
)
