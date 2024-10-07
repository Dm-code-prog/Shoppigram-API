package admin

import (
	"context"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	// Marketplace defines the structure for a Marketplace
	Marketplace struct {
		ID                    uuid.UUID `json:"id"`
		Name                  string    `json:"name"`
		LogoURL               string    `json:"logo_url"`
		ShortName             string    `json:"short_name"`
		IsVerified            bool      `json:"is_verified"`
		OnlinePaymentsEnabled bool      `json:"online_payments_enabled"`
	}

	// GetMarketplacesRequest defines the request for the GetMarketplaces endpoint
	GetMarketplacesRequest struct {
		ExternalUserID int64
	}
	// GetMarketplacesResponse defines the response for the GetMarketplaces endpoint
	GetMarketplacesResponse struct {
		Marketplaces []Marketplace `json:"marketplaces"`
	}

	// CreateMarketplaceRequest creates a new marketplace
	// for a client with a given name and shortname
	CreateMarketplaceRequest struct {
		ShortName      string `json:"short_name"`
		Name           string `json:"name"`
		Currency       string `json:"currency"`
		ExternalUserID int64
	}

	// CreateMarketplaceResponse returns the ID of the created marketplace
	CreateMarketplaceResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// UpdateMarketplaceRequest allows editing the name
	// of the marketplace
	UpdateMarketplaceRequest struct {
		ID             uuid.UUID
		Name           string `json:"name"`
		ExternalUserID int64
	}
	// DeleteMarketplaceRequest specifies a marketplace that needs to be deleted
	DeleteMarketplaceRequest struct {
		WebAppId       uuid.UUID
		ExternalUserID int64
	}

	// CreateProductRequest specifies the information about a product
	CreateProductRequest struct {
		WebAppID       uuid.UUID
		ExternalUserID int64
		Name           string  `json:"name"`
		Description    string  `json:"description"`
		Price          float64 `json:"price"`
		Category       string  `json:"category,omitempty"`
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
		Name           string  `json:"name"`
		Description    string  `json:"description"`
		Price          float64 `json:"price"`
		Category       string  `json:"category,omitempty"`
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

	// CreateMarketplaceLogoUploadURLRequest specifies the request for creating a new marketplace logo upload URL
	CreateMarketplaceLogoUploadURLRequest struct {
		WebAppID       uuid.UUID
		Extension      string `json:"extension"`
		ExternalUserID int64
	}

	// CreateMarketplaceLogoUploadURLResponse specifies the response for creating a new marketplace logo upload URL
	CreateMarketplaceLogoUploadURLResponse struct {
		UploadURL string `json:"upload_url"`
		Key       string `json:"key"`
	}

	// AddUserToNewOrderNotificationsParams mirrors a corresponding struct
	// in notifications module to reduce coupling
	AddUserToNewOrderNotificationsParams struct {
		WebAppID    uuid.UUID
		AdminChatID int64
	}

	// SendMarketplaceBannerParams is a struct for request params to send a marketplace banner to a Telegram channel
	// with a TWA link button markup
	SendMarketplaceBannerParams struct {
		WebAppLink    string
		Message       string
		ChannelChatID int64
	}

	// PinNotificationParams is a struct for request params to pin a message in a Telegram channel
	PinNotificationParams struct {
		ChatID    int64
		MessageID int64
	}

	// CreateOrUpdateTelegramChannelRequest contains the data about a Telegram channel, Shoppigram bot is added to
	CreateOrUpdateTelegramChannelRequest struct {
		ExternalID      int64
		Title           string
		Name            string
		OwnerExternalID int64
		IsPublic        bool
	}

	// GetTelegramChannelOwnerRequest contains a channel id
	GetTelegramChannelOwnerRequest struct {
		ChannelChatId int64
	}

	// GetTelegramChannelsResponse contains the data about Telegram channels owned by a specific user
	GetTelegramChannelOwnerResponse struct {
		ChatId int64
	}

	// PublishMarketplaceBannerToChannelRequest contains the data about a banner to be published to a Telegram channel
	PublishMarketplaceBannerToChannelRequest struct {
		WebAppID          uuid.UUID
		ExternalUserID    int64
		ExternalChannelID int64  `json:"channel_id"`
		Message           string `json:"message"`
		PinMessage        bool   `json:"pin_message"`
	}

	// TelegramChannel contains the data about a Telegram channel
	TelegramChannel struct {
		ID         uuid.UUID `json:"id"`
		ExternalID int64     `json:"external_id"`
		Name       string    `json:"name"`
		Title      string    `json:"title"`
	}

	// GetTelegramChannelsResponse contains the data about Telegram channels owned by a specific user
	GetTelegramChannelsResponse struct {
		Channels []TelegramChannel `json:"channels"`
	}

	// GetOrdersRequest is a filter for getting orders
	GetOrdersRequest struct {
		ExternalUserID int64
		State          string
		MarketplaceID  uuid.UUID
		Limit          int
		Offset         int
	}

	// Product represents a product in a marketplace
	Product struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Quantity int       `json:"quantity"`
		Price    float64   `json:"price"`
	}

	// Order represents an order in a marketplace
	Order struct {
		ID            uuid.UUID `json:"id"`
		MarketplaceID uuid.UUID `json:"marketplace_id"`
		State         string    `json:"state"`
		Type          string    `json:"type"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		ReadableID    int       `json:"readable_id"`
		TotalPrice    float64   `json:"total_price"`
		Currency      string    `json:"currency"`
		BuyerUsername string    `json:"buyer_username"`
		Products      []Product `json:"products"`
	}

	GetOrdersResponse struct {
		Orders []Order `json:"orders"`
	}

	// GetBalanceRequest is a request to get the balance of a user
	// Which is calculated as the sum of all online orders minus the commission
	GetBalanceRequest struct {
		ExternalUserID int64
	}

	// Balance represents the balance of a user
	// in a currency
	Balance struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	}

	// GetBalanceResponse is a response to GetBalanceRequest
	GetBalanceResponse struct {
		Balances []Balance `json:"balances"`
	}
)

type (
	// Repository provides access to the admin storage
	Repository interface {
		GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error)
		GetMarketplaceShortName(ctx context.Context, id uuid.UUID) (string, error)
		CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error)
		UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error
		DeleteMarketplace(ctx context.Context, req DeleteMarketplaceRequest) error

		CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest) error
		DeleteProduct(ctx context.Context, req DeleteProductRequest) error

		GetOrders(ctx context.Context, request GetOrdersRequest) (GetOrdersResponse, error)
		GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error)

		IsUserTheOwnerOfMarketplace(ctx context.Context, externalUserID int64, webAppID uuid.UUID) (bool, error)
		IsUserTheOwnerOfProduct(ctx context.Context, externalUserID int64, productID uuid.UUID) (bool, error)
		IsUserTheOwnerOfTelegramChannel(ctx context.Context, externalUserID, channelID int64) (bool, error)

		GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error)
	}

	// DOSpacesConfig holds the credentials for the S3 bucket
	DOSpacesConfig struct {
		Endpoint string
		ID       string
		Secret   string
		Bucket   string
	}

	Notifier interface {
		AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsParams) error
		SendMarketplaceBanner(ctx context.Context, req SendMarketplaceBannerParams) (messageID int64, err error)
		PinNotification(ctx context.Context, req PinNotificationParams) error
	}

	Service interface {
		GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error)
		CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error)
		UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error
		DeleteMarketplace(ctx context.Context, req DeleteMarketplaceRequest) error

		CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest) error
		DeleteProduct(ctx context.Context, req DeleteProductRequest) error

		GetOrders(ctx context.Context, req GetOrdersRequest) (GetOrdersResponse, error)
		GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error)

		CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error)
		CreateMarketplaceLogoUploadURL(ctx context.Context, request CreateMarketplaceLogoUploadURLRequest) (CreateMarketplaceLogoUploadURLResponse, error)

		GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error)
		PublishMarketplaceBannerToChannel(ctx context.Context, req PublishMarketplaceBannerToChannelRequest) error
	}

	// DefaultService provides admin operations
	DefaultService struct {
		repo     Repository
		spaces   *s3.S3
		bucket   string
		notifier Notifier
		botName  string
	}
)

var (
	shortNameRegex = regexp.MustCompile("^[a-z-0-9]{5,}$")
)

const (
	// possibly make it configurable
	maxMarketplacesThreshold = 8
	maxMarketplaceProducts   = 128
)

// New creates a new admin service
func New(repo Repository, conf DOSpacesConfig, notifier Notifier, botName string) *DefaultService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("fra1"),
		Credentials: credentials.NewStaticCredentials(
			conf.ID,
			conf.Secret,
			"",
		),
		Endpoint:         aws.String(conf.Endpoint),
		S3ForcePathStyle: aws.Bool(false),
	}))

	return &DefaultService{
		repo:     repo,
		spaces:   s3.New(sess),
		bucket:   conf.Bucket,
		notifier: notifier,
		botName:  botName,
	}
}

// GetMarketplaces gets all marketplaces created by user
func (s *DefaultService) GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error) {
	marketplaces, err := s.repo.GetMarketplaces(ctx, req)
	if err != nil {
		return GetMarketplacesResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return marketplaces, nil
}

// CreateMarketplace creates and saves a new marketplace
func (s *DefaultService) CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error) {
	if !isMarketplaceNameValid(req.Name) {
		return CreateMarketplaceResponse{}, ErrorInvalidName
	}

	if !isMarketplaceShortNameValid(req.ShortName) {
		return CreateMarketplaceResponse{}, ErrorInvalidShortName
	}

	res, err := s.repo.CreateMarketplace(ctx, req)
	if err != nil {
		return CreateMarketplaceResponse{}, errors.Wrap(err, "s.repo.CreateMarketplace")
	}

	err = s.notifier.AddUserToNewOrderNotifications(ctx, AddUserToNewOrderNotificationsParams{
		WebAppID:    res.ID,
		AdminChatID: req.ExternalUserID,
	})

	return res, err
}

// UpdateMarketplace edits the name of an existing marketplace
func (s *DefaultService) UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error {
	if !isMarketplaceNameValid(req.Name) {
		return ErrorInvalidName
	}

	err := s.repo.UpdateMarketplace(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateMarketplace")
	}

	return nil
}

// DeleteMarketplace deletes a marketplace
func (s *DefaultService) DeleteMarketplace(ctx context.Context, req DeleteMarketplaceRequest) error {
	ok, err := s.repo.IsUserTheOwnerOfMarketplace(ctx, req.ExternalUserID, req.WebAppId)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfMarketplace")
	}

	if !ok {
		return ErrorOpNotAllowed
	}

	err = s.repo.DeleteMarketplace(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.DeleteMarketplace")
	}

	return nil
}

// CreateProduct creates a new product in a marketplace
func (s *DefaultService) CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error) {
	ok, err := s.repo.IsUserTheOwnerOfMarketplace(ctx, req.ExternalUserID, req.WebAppID)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "s.repo.IsUserTheOwnerOfMarketplace")
	}

	if !ok {
		return CreateProductResponse{}, ErrorOpNotAllowed
	}

	if !isProductNameValid(req.Name) {
		return CreateProductResponse{}, ErrorInvalidName
	}

	res, err := s.repo.CreateProduct(ctx, req)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "s.repo.CreateProduct")
	}

	return res, err
}

// UpdateProduct updates a product of a marketplace
func (s *DefaultService) UpdateProduct(ctx context.Context, req UpdateProductRequest) error {
	if ok, err := s.repo.IsUserTheOwnerOfProduct(ctx, req.ExternalUserID, req.ID); err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfProduct")
	} else if !ok {
		return ErrorOpNotAllowed
	}

	if !isProductNameValid(req.Name) {
		return ErrorInvalidName
	}

	err := s.repo.UpdateProduct(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateProduct")
	}

	return nil
}

func (s *DefaultService) DeleteProduct(ctx context.Context, req DeleteProductRequest) error {
	if ok, err := s.repo.IsUserTheOwnerOfProduct(ctx, req.ExternalUserID, req.ID); err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfProduct")
	} else if !ok {
		return ErrorOpNotAllowed
	}

	err := s.repo.DeleteProduct(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.DeleteProduct")
	}

	return nil
}

// GetOrders gets a list of orders
func (s *DefaultService) GetOrders(ctx context.Context, req GetOrdersRequest) (GetOrdersResponse, error) {
	orders, err := s.repo.GetOrders(ctx, req)
	if err != nil {
		return GetOrdersResponse{}, errors.Wrap(err, "s.repo.GetOrders")
	}

	return orders, nil
}

func (s *DefaultService) GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error) {
	balances, err := s.repo.GetBalance(ctx, req)
	if err != nil {
		return GetBalanceResponse{}, errors.Wrap(err, "s.repo.GetBalance")
	}

	return balances, nil
}

// CreateProductImageUploadURL creates a new upload URL for a product image
func (s *DefaultService) CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error) {
	if ok, err := s.repo.IsUserTheOwnerOfProduct(ctx, request.ExternalUserID, request.ProductID); err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "s.repo.IsUserTheOwnerOfProduct")
	} else if !ok {
		return CreateProductImageUploadURLResponse{}, ErrorOpNotAllowed
	}

	// validate extension
	if !isValidImageExtension(request.Extension) {
		return CreateProductImageUploadURLResponse{}, ErrorInvalidImageExtension
	}

	shortName, err := s.repo.GetMarketplaceShortName(ctx, request.WebAppID)
	if err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "s.repo.GetMarketplaceShortName")
	}

	if shortName == "" {
		return CreateProductImageUploadURLResponse{}, errors.New("s.repo.GetMarketplaceShortName: short name is empty")
	}

	key := shortName + "/" + request.ProductID.String()
	req, _ := s.spaces.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/" + request.Extension),
	})

	url, err := req.Presign(time.Minute)
	if err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "req.Presign")
	}

	return CreateProductImageUploadURLResponse{
		UploadURL: url,
		Key:       key,
	}, nil
}

// CreateMarketplaceLogoUploadURL creates a new upload URL for a marketplace logo
func (s *DefaultService) CreateMarketplaceLogoUploadURL(ctx context.Context, request CreateMarketplaceLogoUploadURLRequest) (CreateMarketplaceLogoUploadURLResponse, error) {
	if ok, err := s.repo.IsUserTheOwnerOfMarketplace(ctx, request.ExternalUserID, request.WebAppID); err != nil {
		return CreateMarketplaceLogoUploadURLResponse{}, errors.Wrap(err, "s.repo.IsUserTheOwnerOfMarketplace")
	} else if !ok {
		return CreateMarketplaceLogoUploadURLResponse{}, ErrorOpNotAllowed
	}

	// validate extension
	if !isValidImageExtension(request.Extension) {
		return CreateMarketplaceLogoUploadURLResponse{}, ErrorInvalidImageExtension
	}

	shortName, err := s.repo.GetMarketplaceShortName(ctx, request.WebAppID)
	if err != nil {
		return CreateMarketplaceLogoUploadURLResponse{}, errors.Wrap(err, "s.repo.GetMarketplaceShortName")
	}

	if shortName == "" {
		return CreateMarketplaceLogoUploadURLResponse{}, errors.New("s.repo.GetMarketplaceShortName: short name is empty")
	}

	key := shortName + "/logo"
	req, _ := s.spaces.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/" + request.Extension),
	})

	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		return CreateMarketplaceLogoUploadURLResponse{}, errors.Wrap(err, "req.Presign")
	}

	return CreateMarketplaceLogoUploadURLResponse{
		UploadURL: url,
		Key:       key,
	}, nil
}

// GetTelegramChannels gets a list of Telegram channels owned by a specific user
func (s *DefaultService) GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error) {
	res, err := s.repo.GetTelegramChannels(ctx, ownerExternalID)
	if err != nil {
		return GetTelegramChannelsResponse{}, errors.Wrap(err, "s.repo.GetTelegramChannels")
	}

	return res, nil
}

// PublishMarketplaceBannerToChannel publishes a banner to a Telegram channel
func (s *DefaultService) PublishMarketplaceBannerToChannel(ctx context.Context, req PublishMarketplaceBannerToChannelRequest) error {
	if req.Message == "" {
		return ErrorBadRequest
	}

	ok, err := s.repo.IsUserTheOwnerOfTelegramChannel(ctx, req.ExternalUserID, req.ExternalChannelID)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfTelegramChannel")
	}
	if !ok {
		return ErrorOpNotAllowed
	}

	ok, err = s.repo.IsUserTheOwnerOfMarketplace(ctx, req.ExternalUserID, req.WebAppID)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfMarketplace")
	}
	if !ok {
		return ErrorOpNotAllowed
	}

	shortName, err := s.repo.GetMarketplaceShortName(ctx, req.WebAppID)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetMarketplaceShortName")
	}

	messageID, err := s.notifier.SendMarketplaceBanner(ctx, SendMarketplaceBannerParams{
		WebAppLink:    "https://t.me/" + s.botName + "/" + shortName,
		Message:       req.Message,
		ChannelChatID: req.ExternalChannelID,
	})
	if err != nil {
		return errors.Wrap(err, "s.notifier.SendMarketplaceBanner")
	}

	if req.PinMessage {
		err = s.notifier.PinNotification(ctx, PinNotificationParams{
			ChatID:    req.ExternalChannelID,
			MessageID: messageID,
		})
		if err != nil {
			return errors.Wrap(err, "s.notifier.PinNotification")
		}
	}

	return nil
}

func isMarketplaceShortNameValid(shortName string) bool {
	return shortNameRegex.MatchString(shortName)
}

func isMarketplaceNameValid(name string) bool {
	return len(name) >= 3
}

func isProductNameValid(name string) bool {
	return len(name) >= 3 && len(name) <= 75
}

func isValidImageExtension(ext string) bool {
	switch ext {
	case "png", "jpg", "jpeg", "webp":
		return true
	}

	return false
}
