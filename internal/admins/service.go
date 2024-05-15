package admins

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"time"
)

type (
	// Marketplace defines the structure for a Marketplace
	Marketplace struct {
		ID         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		LogoURL    string    `json:"logo_url"`
		IsVerified bool      `json:"is_verified"`
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

	// CreateProductRequest specifies the information about a product
	CreateProductRequest struct {
		WebAppID       uuid.UUID
		ExternalUserID int64
		Name           string  `json:"name"`
		Description    string  `json:"description"`
		Price          float64 `json:"price"`
		PriceCurrency  string  `json:"price_currency"`
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
		PriceCurrency  string  `json:"price_currency"`
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
)

type (
	// Repository provides access to the admin storage
	Repository interface {
		GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error)
		GetMarketplaceShortName(ctx context.Context, id uuid.UUID) (string, error)
		CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error)
		UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error

		CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error)
		UpdateProduct(ctx context.Context, req UpdateProductRequest) error
		DeleteProduct(ctx context.Context, req DeleteProductRequest) error

		IsUserTheOwnerOfMarketplace(ctx context.Context, externalUserID int64, webAppID uuid.UUID) (bool, error)
		IsUserTheOwnerOfProduct(ctx context.Context, externalUserID int64, productID uuid.UUID) (bool, error)
	}

	// DOSpacesConfig holds the credentials for the S3 bucket
	DOSpacesConfig struct {
		Endpoint string
		ID       string
		Secret   string
		Bucket   string
	}

	// Service provides admin operations
	Service struct {
		repo   Repository
		spaces *s3.S3
		log    *zap.Logger
		bucket string
	}
)

var (
	ErrorAdminNotFound = errors.New("admin not found")
	ErrorInternal      = errors.New("internal server error")
	ErrorBadRequest    = errors.New("bad request")

	ErrorNotUniqueShortName      = errors.New("not unique short name")
	ErrorInvalidShortName        = errors.New("invalid short name")
	ErrorInvalidName             = errors.New("invalid name")
	ErrorMaxMarketplacesExceeded = errors.New("max marketplaces exceeded")

	ErrorMaxProductsExceeded    = errors.New("max products exceeded")
	ErrorInvalidProductCurrency = errors.New("invalid product currency")

	ErrorOpNotAllowed = errors.New("operation not allowed")

	ErrorInvalidImageExtension = errors.New("invalid image extension, only png, jpg, jpeg are allowed")
)

var (
	shortNameRegex = regexp.MustCompile("^[a-z]{5,}$")
)

const (
	// possibly make it configurable
	maxMarketplacesThreshold = 8
	maxMarketplaceProducts   = 128
)

// New creates a new admin service
func New(repo Repository, log *zap.Logger, conf DOSpacesConfig) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}

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

	return &Service{
		repo:   repo,
		log:    log,
		spaces: s3.New(sess),
		bucket: conf.Bucket,
	}
}

// GetMarketplaces gets all marketplaces created by user
func (s *Service) GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error) {
	marketplaces, err := s.repo.GetMarketplaces(ctx, req)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.GetProducts"),
			zap.String("user_id", strconv.FormatInt(req.ExternalUserID, 10)),
		).Error(err.Error())
		return GetMarketplacesResponse{}, errors.Wrap(err, "s.repo.CreateOrUpdateTgUser")
	}

	return marketplaces, nil
}

// CreateMarketplace creates and saves a new marketplace
func (s *Service) CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error) {
	if !isMarketplaceNameValid(req.Name) {
		return CreateMarketplaceResponse{}, ErrorInvalidName
	}

	if !isMarketplaceShortNameValid(req.ShortName) {
		return CreateMarketplaceResponse{}, ErrorInvalidShortName
	}

	res, err := s.repo.CreateMarketplace(ctx, req)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.CreateProducts"),
			zap.String("user_id", strconv.FormatInt(req.ExternalUserID, 10))).Error(err.Error())
		return CreateMarketplaceResponse{}, errors.Wrap(err, "s.repo.CreateMarketplace")
	}

	return res, err
}

// UpdateMarketplace edits the name of an existing marketplace
func (s *Service) UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error {
	if !isMarketplaceNameValid(req.Name) {
		return ErrorInvalidName
	}

	err := s.repo.UpdateMarketplace(ctx, req)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.UpdateProducts"),
			zap.String("user_id", strconv.FormatInt(req.ExternalUserID, 10)),
		).Error(err.Error())
		return errors.Wrap(err, "s.repo.UpdateMarketplace")
	}

	return nil
}

// CreateProduct creates a new product in a marketplace
func (s *Service) CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error) {
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

	if req.Price <= 0 {
		return CreateProductResponse{}, ErrorBadRequest
	}

	res, err := s.repo.CreateProduct(ctx, req)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.CreateProducts"),
			zap.String("web_app_id", req.WebAppID.String()),
		).Error(err.Error())
		return CreateProductResponse{}, errors.Wrap(err, "s.repo.CreateProduct")
	}

	return res, err
}

// UpdateProduct updates a product of a marketplace
func (s *Service) UpdateProduct(ctx context.Context, req UpdateProductRequest) error {
	if ok, err := s.repo.IsUserTheOwnerOfProduct(ctx, req.ExternalUserID, req.ID); err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfProduct")
	} else if !ok {
		return ErrorOpNotAllowed
	}

	if !isProductNameValid(req.Name) {
		return ErrorInvalidName
	}

	if req.Price <= 0 {
		return ErrorBadRequest
	}

	err := s.repo.UpdateProduct(ctx, req)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.UpdateProducts"),
			zap.String("web_app_id", req.WebAppID.String()),
			zap.String("product_id", req.ID.String()),
		).Error(err.Error())
		return errors.Wrap(err, "s.repo.UpdateProduct")
	}

	return nil
}

func (s *Service) DeleteProduct(ctx context.Context, req DeleteProductRequest) error {
	if ok, err := s.repo.IsUserTheOwnerOfProduct(ctx, req.ExternalUserID, req.ID); err != nil {
		return errors.Wrap(err, "s.repo.IsUserTheOwnerOfProduct")
	} else if !ok {
		return ErrorOpNotAllowed
	}

	err := s.repo.DeleteProduct(ctx, req)
	if err != nil {
		s.log.With(
			zap.String("method", "s.repo.DeleteProduct"),
			zap.String("web_app_id", req.WebAppID.String()),
			zap.String("product_id", req.ID.String()),
		).Error(err.Error())
		return errors.Wrap(err, "s.repo.DeleteProduct")
	}

	return nil
}

// CreateProductImageUploadURL creates a new upload URL for a product image
func (s *Service) CreateProductImageUploadURL(ctx context.Context, request CreateProductImageUploadURLRequest) (CreateProductImageUploadURLResponse, error) {
	if ok, err := s.repo.IsUserTheOwnerOfProduct(ctx, request.ExternalUserID, request.ProductID); err != nil {
		s.log.With(
			zap.String("method", "s.repo.IsUserTheOwnerOfProduct"),
			zap.String("user_id", strconv.FormatInt(request.ExternalUserID, 10)),
			zap.String("product_id", request.ProductID.String())).Error(err.Error())
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

	key := shortName + "/" + request.ProductID.String() + "." + request.Extension
	req, _ := s.spaces.PutObjectRequest(&s3.PutObjectInput{
		Bucket:   aws.String(s.bucket),
		Key:      aws.String(key),
		ACL:      aws.String("public"),
		Metadata: map[string]*string{},
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
func (s *Service) CreateMarketplaceLogoUploadURL(ctx context.Context, request CreateMarketplaceLogoUploadURLRequest) (CreateMarketplaceLogoUploadURLResponse, error) {
	if ok, err := s.repo.IsUserTheOwnerOfMarketplace(ctx, request.ExternalUserID, request.WebAppID); err != nil {
		s.log.With(
			zap.String("method", "s.repo.IsUserTheOwnerOfMarketplace"),
			zap.String("user_id", strconv.FormatInt(request.ExternalUserID, 10)),
			zap.String("web_app_id", request.WebAppID.String())).Error(err.Error())
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

	key := shortName + "/logo." + request.Extension
	req, _ := s.spaces.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		ACL:    aws.String("public"),
	})

	url, err := req.Presign(time.Minute)
	if err != nil {
		return CreateMarketplaceLogoUploadURLResponse{}, errors.Wrap(err, "req.Presign")
	}

	return CreateMarketplaceLogoUploadURLResponse{
		UploadURL: url,
		Key:       key,
	}, nil
}

func isMarketplaceShortNameValid(shortName string) bool {
	return shortNameRegex.MatchString(shortName)
}

func isMarketplaceNameValid(name string) bool {
	return len(name) >= 3
}

func isProductNameValid(name string) bool {
	return len(name) >= 3 && len(name) <= 30
}

func isValidImageExtension(ext string) bool {
	switch ext {
	case "png", "jpg", "jpeg":
		return true
	}

	return false
}
