package admins

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/admins/generated"
)

const (
	defaultLimit = 50
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen *generated.Queries
}

// NewPg creates a new Pg
func NewPg(db *pgxpool.Pool) *Pg {
	return &Pg{gen: generated.New(db)}
}

// GetMarketplaces gets all marketplaces created by user
func (p *Pg) GetMarketplaces(ctx context.Context, req GetMarketplacesRequest) (GetMarketplacesResponse, error) {
	rows, err := p.gen.GetMarketplaces(ctx, pgtype.Int4{
		Int32: int32(req.ExternalUserID),
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetMarketplacesResponse{}, errors.Wrap(ErrorAdminNotFound, "p.gen.GetMarketplaces")
		}
		return GetMarketplacesResponse{}, errors.Wrap(err, "p.gen.GetMarketplaces")
	}

	marketplaces := make([]Marketplace, len(rows))

	for i, v := range rows {
		marketplaces[i] = Marketplace{
			ID:         v.ID,
			Name:       v.Name,
			LogoURL:    v.LogoUrl.String,
			IsVerified: v.IsVerified.Bool,
			ShortName:  v.ShortName,
		}
	}

	return GetMarketplacesResponse{
		Marketplaces: marketplaces,
	}, nil
}

// CreateMarketplace stores the marketplace information in the database
// It creates the ID for the marketplace and returns it
func (p *Pg) CreateMarketplace(ctx context.Context, req CreateMarketplaceRequest) (CreateMarketplaceResponse, error) {
	count, err := p.gen.CountUserMarketplaces(ctx, pgtype.Int4{
		Int32: int32(req.ExternalUserID),
		Valid: true,
	})
	if err != nil {
		return CreateMarketplaceResponse{}, errors.Wrap(err, "p.gen.CountUserMarketplaces")
	}

	if count > maxMarketplacesThreshold {
		return CreateMarketplaceResponse{}, ErrorMaxMarketplacesExceeded
	}

	id, err := p.gen.CreateMarketplace(ctx, generated.CreateMarketplaceParams{
		Name:            req.Name,
		ShortName:       req.ShortName,
		OwnerExternalID: pgtype.Int4{Int32: int32(req.ExternalUserID), Valid: true},
	})
	if err != nil {
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			return CreateMarketplaceResponse{}, ErrorNotUniqueShortName
		}
		return CreateMarketplaceResponse{}, errors.Wrap(err, "p.gen.CreateMarketplace")
	}

	return CreateMarketplaceResponse{ID: id}, err
}

// UpdateMarketplace updates the name of the marketplace in the database
func (p *Pg) UpdateMarketplace(ctx context.Context, req UpdateMarketplaceRequest) error {
	execRes, err := p.gen.UpdateMarketplace(ctx, generated.UpdateMarketplaceParams{
		ID:              req.ID,
		Name:            req.Name,
		OwnerExternalID: pgtype.Int4{Int32: int32(req.ExternalUserID), Valid: true},
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.UpdateMarketplace")
	}

	if execRes.RowsAffected() == 0 {
		return ErrorOpNotAllowed
	}

	return nil
}

// DeleteMarketplace soft deletes marketplace
func (p *Pg) DeleteMarketplace(ctx context.Context, req DeleteMarketplaceRequest) error {
	err := p.gen.SoftDeleteMarketplace(ctx, req.WebAppId)
	if err != nil {
		return errors.Wrap(err, "p.gen.SoftDeleteMarketplace")
	}

	return nil
}

// CreateProduct saves a product to the database
// and returns the ID that it assigned to it
func (p *Pg) CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error) {
	count, err := p.gen.CountMarketplaceProducts(ctx, req.WebAppID)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "p.gen.CountMarketplaceProducts")
	}

	if count > maxMarketplaceProducts {
		return CreateProductResponse{}, ErrorMaxProductsExceeded
	}

	id, err := p.gen.CreateProduct(ctx, generated.CreateProductParams{
		WebAppID:    req.WebAppID,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
	})
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "p.gen.CreateProduct")
	}

	return CreateProductResponse{ID: id}, err
}

// UpdateProduct updates the product of a marketplace in the database
func (p *Pg) UpdateProduct(ctx context.Context, req UpdateProductRequest) error {
	execRes, err := p.gen.UpdateProduct(ctx, generated.UpdateProductParams{
		ID:          req.ID,
		WebAppID:    req.WebAppID,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.UpdateProduct")
	}

	if execRes.RowsAffected() == 0 {
		return ErrorOpNotAllowed
	}

	return nil
}

// DeleteProduct deletes a product from a marketplace
func (p *Pg) DeleteProduct(ctx context.Context, req DeleteProductRequest) error {
	execRes, err := p.gen.DeleteProduct(ctx, generated.DeleteProductParams{WebAppID: req.WebAppID, ID: req.ID})
	if err != nil {
		return errors.Wrap(err, "p.gen.DeleteProduct")
	}

	if execRes.RowsAffected() == 0 {
		return ErrorOpNotAllowed
	}

	return nil
}

// GetOrders gets a list of orders and allows filtering by marketplace and state
func (p *Pg) GetOrders(ctx context.Context, req GetOrdersRequest) (GetOrdersResponse, error) {
	params := generated.GetOrdersParams{
		Limit:           int32(req.Limit),
		Offset:          int32(req.Offset),
		OwnerExternalID: int32(req.ExternalUserID),
		MarketplaceID:   req.MarketplaceID,
		State:           req.State,
	}
	if req.Limit == 0 {
		params.Limit = defaultLimit
	}

	rows, err := p.gen.GetOrders(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetOrdersResponse{Orders: make([]Order, 0)}, err
		}

		return GetOrdersResponse{}, errors.Wrap(err, "p.gen.GetOrders")
	}

	orders := make([]Order, len(rows))

	for i, v := range rows {
		products := make([]Product, 0)
		err := json.Unmarshal(v.Products, &products)
		if err != nil {
			return GetOrdersResponse{}, errors.Wrap(err, "json.Unmarshal")
		}

		orders[i] = Order{
			ID:            v.ID,
			MarketplaceID: v.MarketplaceID.Bytes,
			ReadableID:    int(v.ReadableID.Int64),
			TotalPrice:    float64(v.TotalPrice),
			BuyerUsername: v.BuyerUsername.String,
			Products:      products,
			State:         string(v.State),
			CreatedAt:     v.CreatedAt.Time,
			UpdatedAt:     v.UpdatedAt.Time,
		}
	}

	return GetOrdersResponse{Orders: orders}, nil
}

// GetBalance returns balances in all currencies
func (p *Pg) GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error) {
	rows, err := p.gen.GetBalance(ctx, pgtype.Int4{Int32: int32(req.ExternalUserID), Valid: true})
	if err != nil {
		return GetBalanceResponse{}, errors.Wrap(err, "p.gen.GetBalance")
	}

	balances := make([]Balance, 0, len(rows))
	for _, r := range rows {
		balances = append(balances, Balance{
			Currency: string(r.Currency),
			Balance:  float64(r.Balance),
		})
	}

	return GetBalanceResponse{Balances: balances}, nil
}

// IsUserTheOwnerOfMarketplace checks if the user is the owner of the marketplace
func (p *Pg) IsUserTheOwnerOfMarketplace(ctx context.Context, userID int64, webAppID uuid.UUID) (bool, error) {
	ok, err := p.gen.IsUserTheOwnerOfWebApp(ctx, generated.IsUserTheOwnerOfWebAppParams{
		OwnerExternalID: pgtype.Int4{Int32: int32(userID), Valid: true},
		ID:              webAppID,
	})
	if err != nil {
		return false, errors.Wrap(err, "p.gen.IsUserTheOwnerOfWebApp")
	}

	return ok, nil
}

// IsUserTheOwnerOfProduct checks if the user is the owner of the product
func (p *Pg) IsUserTheOwnerOfProduct(ctx context.Context, userID int64, productID uuid.UUID) (bool, error) {
	ok, err := p.gen.IsUserTheOwnerOfProduct(ctx, generated.IsUserTheOwnerOfProductParams{
		OwnerExternalID: pgtype.Int4{Int32: int32(userID), Valid: true},
		ID:              productID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, ErrorOpNotAllowed
		}
		return false, errors.Wrap(err, "p.gen.IsUserTheOwnerOfProduct")
	}

	return ok, nil
}

func (p *Pg) IsUserTheOwnerOfTelegramChannel(ctx context.Context, externalUserID, channelID int64) (bool, error) {
	ok, err := p.gen.IsUserTheOwnerOfTelegramChannel(ctx, generated.IsUserTheOwnerOfTelegramChannelParams{
		OwnerExternalID: externalUserID,
		ExternalID:      channelID,
	})
	if err != nil {
		return false, ErrorOpNotAllowed
	}

	return ok, nil
}

func (p *Pg) GetMarketplaceShortName(ctx context.Context, id uuid.UUID) (string, error) {
	return p.gen.GetMarketplaceShortName(ctx, id)
}

// CreateOrUpdateTelegramChannel creates or updates a Telegram channel
// that was integrated with Shoppigram
func (p *Pg) CreateOrUpdateTelegramChannel(ctx context.Context, req CreateOrUpdateTelegramChannelRequest) error {
	err := p.gen.CreateOrUpdateTelegramChannel(ctx, generated.CreateOrUpdateTelegramChannelParams{
		ExternalID:      req.ExternalID,
		Name:            pgtype.Text{String: req.Name, Valid: req.Name != ""},
		Title:           req.Title,
		IsPublic:        req.IsPublic,
		OwnerExternalID: req.OwnerExternalID,
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.CreateOrUpdateTelegramChannel")
	}

	return nil
}

// GetTelegramChannels gets a list of Telegram channels owned by a specific user
func (p *Pg) GetTelegramChannels(ctx context.Context, ownerExternalID int64) (GetTelegramChannelsResponse, error) {
	rows, err := p.gen.GetTelegramChannels(ctx, ownerExternalID)
	if err != nil {
		return GetTelegramChannelsResponse{}, errors.Wrap(err, "p.gen.GetTelegramChannels")
	}

	channels := make([]TelegramChannel, len(rows))

	for i, v := range rows {
		channels[i] = TelegramChannel{
			ID:         v.ID,
			Name:       v.Name.String,
			Title:      v.Title,
			ExternalID: v.ExternalID,
		}
	}

	return GetTelegramChannelsResponse{Channels: channels}, nil
}
