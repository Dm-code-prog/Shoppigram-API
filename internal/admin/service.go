package admin

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

type (
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
	maxShops    = 8
	maxProducts = 128
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

// GetShops returns all shops created by user
func (s *DefaultService) GetShops(ctx context.Context, req GetShopsRequest) (GetShopsResponse, error) {
	shops, err := s.repo.GetShops(ctx, req)
	if err != nil {
		return GetShopsResponse{}, errors.Wrap(err, "s.repo.GetShops")
	}

	return shops, nil
}

// CreateShop creates a new Shop
func (s *DefaultService) CreateShop(ctx context.Context, req CreateShopRequest) (CreateShopResponse, error) {
	if !isNameValid(req.Name) {
		return CreateShopResponse{}, ErrorInvalidName
	}

	if !isShortNameValid(req.ShortName) {
		return CreateShopResponse{}, ErrorInvalidShortName
	}

	res, err := s.repo.CreateShop(ctx, req)
	if err != nil {
		return CreateShopResponse{}, errors.Wrap(err, "s.repo.CreateShop")
	}

	err = s.notifier.AddUserToNewOrderNotifications(ctx, AddUserToNewOrderNotificationsParams{
		WebAppID:    res.ID,
		AdminChatID: req.ExternalUserID,
	})

	return res, err
}

// UpdateShop updates a Shop
func (s *DefaultService) UpdateShop(ctx context.Context, req UpdateShopRequest) error {
	if !isNameValid(req.Name) {
		return ErrorInvalidName
	}

	err := s.repo.UpdateShop(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateShop")
	}

	return nil
}

// DeleteShop deletes a shop
func (s *DefaultService) DeleteShop(ctx context.Context, req DeleteShopRequest) error {
	ok, err := s.repo.IsShopOwner(ctx, req.ExternalUserID, req.WebAppId)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsShopOwner")
	}

	if !ok {
		return ErrorOpNotAllowed
	}

	err = s.repo.SoftDeleteShop(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.SoftDeleteShop")
	}

	return nil
}

// CreateProduct creates a new product in a marketplace
func (s *DefaultService) CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error) {
	ok, err := s.repo.IsShopOwner(ctx, req.ExternalUserID, req.WebAppID)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "s.repo.IsShopOwner")
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
	if ok, err := s.repo.IsProductOwner(ctx, req.ExternalUserID, req.ID); err != nil {
		return errors.Wrap(err, "s.repo.IsProductOwner")
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
	if ok, err := s.repo.IsProductOwner(ctx, req.ExternalUserID, req.ID); err != nil {
		return errors.Wrap(err, "s.repo.IsProductOwner")
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
	if ok, err := s.repo.IsProductOwner(ctx, request.ExternalUserID, request.ProductID); err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "s.repo.IsProductOwner")
	} else if !ok {
		return CreateProductImageUploadURLResponse{}, ErrorOpNotAllowed
	}

	// validate extension
	if !isValidImageExtension(request.Extension) {
		return CreateProductImageUploadURLResponse{}, ErrorInvalidImageExtension
	}

	shortName, err := s.repo.GetShortName(ctx, request.WebAppID)
	if err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "s.repo.GetShortName")
	}

	if shortName == "" {
		return CreateProductImageUploadURLResponse{}, errors.New("s.repo.GetShortName: short name is empty")
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

// CreateShopLogoUploadURL creates a new upload URL for a shop logo
func (s *DefaultService) CreateShopLogoUploadURL(ctx context.Context, request CreateShopLogoUploadURLRequest) (CreateShopLogoUploadURLResponse, error) {
	if ok, err := s.repo.IsShopOwner(ctx, request.ExternalUserID, request.WebAppID); err != nil {
		return CreateShopLogoUploadURLResponse{}, errors.Wrap(err, "s.repo.IsShopOwner")
	} else if !ok {
		return CreateShopLogoUploadURLResponse{}, ErrorOpNotAllowed
	}

	// validate extension
	if !isValidImageExtension(request.Extension) {
		return CreateShopLogoUploadURLResponse{}, ErrorInvalidImageExtension
	}

	shortName, err := s.repo.GetShortName(ctx, request.WebAppID)
	if err != nil {
		return CreateShopLogoUploadURLResponse{}, errors.Wrap(err, "s.repo.GetShortName")
	}

	if shortName == "" {
		return CreateShopLogoUploadURLResponse{}, errors.New("s.repo.GetShortName: short name is empty")
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
		return CreateShopLogoUploadURLResponse{}, errors.Wrap(err, "req.Presign")
	}

	return CreateShopLogoUploadURLResponse{
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

// PublishShopBannerToChannel publishes a banner to a Telegram channel
func (s *DefaultService) PublishShopBannerToChannel(ctx context.Context, req PublishShopBannerToChannelRequest) error {
	if req.Message == "" {
		return ErrorBadRequest
	}

	ok, err := s.repo.IsTelegramChannelOwner(ctx, req.ExternalUserID, req.ExternalChannelID)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsTelegramChannelOwner")
	}
	if !ok {
		return ErrorOpNotAllowed
	}

	ok, err = s.repo.IsShopOwner(ctx, req.ExternalUserID, req.WebAppID)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsShopOwner")
	}
	if !ok {
		return ErrorOpNotAllowed
	}

	shortName, err := s.repo.GetShortName(ctx, req.WebAppID)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetShortName")
	}

	messageID, err := s.notifier.SendMarketplaceBanner(ctx, SendShopBannerParams{
		WebAppLink:    makeShopURL(s.botName, shortName),
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

func makeShopURL(botName, shortName string) string {
	return fmt.Sprintf("https://t.me/%s/shop?startapp=shop_%s", botName, shortName)
}
