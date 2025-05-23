package admin

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

type (
	// DefaultService provides admin operations
	DefaultService struct {
		repo     Repository
		s3       *s3.S3
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
func New(repo Repository, notifier Notifier, s3Instance *s3.S3, botName, s3Bucket string) *DefaultService {
	return &DefaultService{
		repo:     repo,
		s3:       s3Instance,
		bucket:   s3Bucket,
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

// GetShop returns a shop by ID
func (s *DefaultService) GetShop(ctx context.Context, req GetShopRequest) (GetShopResponse, error) {
	if err := s.verifyAccessToShop(ctx, req.ExternalUserID, req.WebAppID); err != nil {
		return GetShopResponse{}, err
	}

	shop, err := s.repo.GetShop(ctx, req)
	if err != nil {
		return GetShopResponse{}, errors.Wrap(err, "s.repo.GetShop")
	}
	return shop, nil
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
	if err := s.verifyAccessToShop(ctx, req.ExternalUserID, req.WebAppId); err != nil {
		return err
	}

	err := s.repo.SoftDeleteShop(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.SoftDeleteShop")
	}
	return nil
}

// ConfigureShopSync enables shop synchronization
func (s *DefaultService) ConfigureShopSync(ctx context.Context, request ConfigureShopSyncRequest) error {
	if err := s.verifyAccessToShop(ctx, request.ExternalUserID, request.WebAppID); err != nil {
		return err
	}

	err := s.repo.ConfigureShopSync(ctx, request)
	if err != nil {
		return errors.Wrap(err, "s.repo.ConfigureShopSync")
	}
	return nil
}

// CreateProduct creates a new product in a marketplace
func (s *DefaultService) CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error) {
	if !isProductNameValid(req.Name) {
		return CreateProductResponse{}, ErrorInvalidName
	}

	if err := s.verifyAccessToShop(ctx, req.ExternalUserID, req.WebAppID); err != nil {
		return CreateProductResponse{}, err
	}

	res, err := s.repo.CreateProduct(ctx, req)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "s.repo.CreateProduct")
	}
	return res, err
}

// UpdateProduct updates a product of a marketplace
func (s *DefaultService) UpdateProduct(ctx context.Context, req UpdateProductRequest) error {
	if !isProductNameValid(req.Name) {
		return ErrorInvalidName
	}

	if err := s.verifyAccessToShop(ctx, req.ExternalUserID, req.WebAppID); err != nil {
		return err
	}

	err := s.repo.UpdateProduct(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateProduct")
	}
	return nil
}

// DeleteProduct marks a product as deleted
func (s *DefaultService) DeleteProduct(ctx context.Context, req DeleteProductRequest) error {
	if err := s.verifyAccessToShop(ctx, req.ExternalUserID, req.WebAppID); err != nil {
		return err
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

// GetBalance TODO: fix this shit
func (s *DefaultService) GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error) {
	balances, err := s.repo.GetBalance(ctx, req)
	if err != nil {
		return GetBalanceResponse{}, errors.Wrap(err, "s.repo.GetBalance")
	}
	return balances, nil
}

// CreateProductImageUploadURL creates a new upload URL for a product image
func (s *DefaultService) CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error) {
	if err := s.verifyAccessToShop(ctx, request.ExternalUserID, request.WebAppID); err != nil {
		return CreateProductImageUploadURLResponse{}, err
	}

	shop, err := s.repo.GetShop(ctx, GetShopRequest{WebAppID: request.WebAppID})
	if err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "s.repo.GetShortName")
	}

	key := fmt.Sprintf("%s/%s", shop.ShortName, request.ProductID.String())
	url, err := s.presignURL(key, request.Extension, time.Minute)
	if err != nil {
		return CreateProductImageUploadURLResponse{}, errors.Wrap(err, "s.presignURL")
	}
	return CreateProductImageUploadURLResponse{
		UploadURL: url,
		Key:       key,
	}, nil
}

// CreateShopLogoUploadURL creates a new upload URL for a shop logo
func (s *DefaultService) CreateShopLogoUploadURL(ctx context.Context, request CreateShopLogoUploadURLRequest) (CreateShopLogoUploadURLResponse, error) {
	if err := s.verifyAccessToShop(ctx, request.ExternalUserID, request.WebAppID); err != nil {
		return CreateShopLogoUploadURLResponse{}, err
	}

	shop, err := s.repo.GetShop(ctx, GetShopRequest{WebAppID: request.WebAppID})
	if err != nil {
		return CreateShopLogoUploadURLResponse{}, errors.Wrap(err, "s.repo.GetShortName")
	}

	key := shop.ShortName + "/logo"
	url, err := s.presignURL(key, request.Extension, time.Minute)
	if err != nil {
		return CreateShopLogoUploadURLResponse{}, errors.Wrap(err, "s.presignURL")
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

	if ok, err := s.repo.IsTelegramChannelOwner(ctx, req.ExternalUserID, req.ExternalChannelID); err != nil {
		return errors.Wrap(err, "s.repo.IsTelegramChannelOwner")
	} else if !ok {
		return ErrorOpNotAllowed
	}

	shop, err := s.repo.GetShop(ctx, GetShopRequest{
		ExternalUserID: req.ExternalUserID,
		WebAppID:       req.WebAppID,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.GetShortName")
	}

	messageID, err := s.notifier.SendMarketplaceBanner(ctx, SendShopBannerParams{
		WebAppLink:    makeShopURL(s.botName, shop.ShortName),
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

func (s *DefaultService) presignURL(key, extension string, ttl time.Duration) (string, error) {
	if !isValidImageExtension(extension) {
		return "", ErrorInvalidImageExtension
	}

	req, _ := s.s3.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ACL:         aws.String("public-read"),
		ContentType: aws.String("image/" + extension),
	})

	url, err := req.Presign(ttl)
	if err != nil {
		return "", errors.Wrap(err, "req.Presign")
	}

	return url, nil
}

func (s *DefaultService) verifyAccessToShop(ctx context.Context, externalUserID int64, webAppID uuid.UUID) error {
	ok, err := s.repo.IsShopOwner(ctx, externalUserID, webAppID)
	if err != nil {
		return errors.Wrap(err, "s.repo.IsShopOwner")
	}
	if !ok {
		return ErrorOpNotAllowed
	}

	return nil
}

func makeShopURL(botName, shortName string) string {
	return fmt.Sprintf("https://t.me/%s/shop?startapp=shop_%s", botName, shortName)
}
