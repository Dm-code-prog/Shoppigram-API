package admins

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"regexp"
	"strconv"
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
		WebAppID      uuid.UUID
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		PriceCurrency string  `json:"price_currency"`
		Category      string  `json:"category,omitempty"`
	}

	// CreateProductResponse returns the ID of the created product
	CreateProductResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// UpdateProductRequest specifies the new information about a product
	// in a marketplace
	UpdateProductRequest struct {
		ID            uuid.UUID `json:"id"`
		WebAppID      uuid.UUID
		Name          string  `json:"name"`
		Description   string  `json:"description"`
		Price         float64 `json:"price"`
		PriceCurrency string  `json:"price_currency"`
		Category      string  `json:"category,omitempty"`
	}

	// DeleteProductRequest specifies a product in a marketplace that needs to be deleted
	DeleteProductRequest struct {
		WebAppID uuid.UUID
		ID       uuid.UUID `json:"id"`
	}
)

type (
	// Repository provides access to the admin storage
	Repository interface {
		GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error)
		CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error)
		UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error
	}

	// Service provides admin operations
	Service struct {
		repo Repository
		log  *zap.Logger
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
func New(repo Repository, log *zap.Logger) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}

	return &Service{
		repo: repo,
		log:  log,
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
}

// UpdateProduct updates a product of a marketplace
func (s *Service) UpdateProduct(ctx context.Context, req UpdateProductRequest) error {}

func (s *Service) DeleteProduct(ctx context.Context, req DeleteProductRequest) error {}

// isUserTheOwnerOfMarketplace checks if a user owns the marketplace
// and therefore has the right to edit it and it's products
func (s *Service) isUserTheOwnerOfMarketplace(ctx context.Context, externalUserID int64) (bool, error) {
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
