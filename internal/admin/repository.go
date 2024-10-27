package admin

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/admin/generated"
)

const (
	defaultLimit = 50
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	pool *pgxpool.Pool
	gen  *generated.Queries
}

// NewPg creates a new Pg
func NewPg(db *pgxpool.Pool) *Pg {
	return &Pg{
		gen:  generated.New(db),
		pool: db,
	}
}

// GetShops returns user's shops
func (p *Pg) GetShops(ctx context.Context, req GetShopsRequest) (GetShopsResponse, error) {
	shops := make([]Shop, 0)

	rows, err := p.gen.GetShops(ctx, pgtype.Int8{
		Int64: req.ExternalUserID,
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetShopsResponse{
				Shops: shops,
			}, nil
		}
		return GetShopsResponse{}, errors.Wrap(err, "p.gen.GetShops")
	}

	for _, v := range rows {
		var syncDetails *SyncDetails
		if v.SyncProvider.Valid {
			syncDetails = &SyncDetails{
				ExternalProvider: string(v.SyncProvider.ExternalProvider),
				IsActive:         v.SyncIsActive.Bool,
				LastSyncedAt:     v.LastSyncAt.Time,
				LastStatus:       string(v.LastSyncStatus.ExtenalSyncStatus),
			}
		}

		shops = append(shops, Shop{
			ID:          v.ID,
			Name:        v.Name,
			IsVerified:  v.IsVerified.Bool,
			ShortName:   v.ShortName,
			Type:        shopType(v.Type),
			Currency:    string(v.Currency),
			SyncDetails: syncDetails,
		})
	}

	return GetShopsResponse{
		Shops: shops,
	}, nil
}

// GetShop returns a shop by ID
func (p *Pg) GetShop(ctx context.Context, req GetShopRequest) (GetShopResponse, error) {
	shop, err := p.gen.GetShop(ctx, req.WebAppID)
	if err != nil {
		return GetShopResponse{}, errors.Wrap(err, "p.gen.GetShop")
	}

	var syncDetails *SyncDetails
	if shop.SyncProvider.Valid {
		syncDetails = &SyncDetails{
			ExternalProvider: string(shop.SyncProvider.ExternalProvider),
			IsActive:         shop.SyncIsActive.Bool,
			LastSyncedAt:     shop.LastSyncAt.Time,
			LastStatus:       string(shop.LastSyncStatus.ExtenalSyncStatus),
		}
	}

	return GetShopResponse{
		ID:          shop.ID,
		Name:        shop.Name,
		IsVerified:  shop.IsVerified.Bool,
		ShortName:   shop.ShortName,
		Type:        shopType(shop.Type),
		Currency:    string(shop.Currency),
		SyncDetails: syncDetails,
	}, nil
}

// CreateShop creates a new shop in database
func (p *Pg) CreateShop(ctx context.Context, req CreateShopRequest) (CreateShopResponse, error) {
	if count, err := p.gen.CountUserShops(ctx, pgtype.Int8{
		Int64: req.ExternalUserID,
		Valid: true,
	}); err != nil {
		return CreateShopResponse{}, errors.Wrap(err, "p.gen.CountUserMarketplaces")
	} else if count > maxShops {
		return CreateShopResponse{}, MaxShopsLimitExceeded
	}

	id, err := p.gen.CreateShop(ctx, generated.CreateShopParams{
		Name:            req.Name,
		ShortName:       req.ShortName,
		OwnerExternalID: pgtype.Int8{Int64: req.ExternalUserID, Valid: true},
		Currency:        generated.ProductCurrency(req.Currency),
		Type:            generated.WebAppType(req.Type),
	})
	if err != nil {
		if strings.Contains(err.Error(), pgerrcode.UniqueViolation) {
			return CreateShopResponse{}, ErrorNotUniqueShortName
		}
		return CreateShopResponse{}, errors.Wrap(err, "p.gen.CreateShop")
	}

	return CreateShopResponse{ID: id}, err
}

// UpdateShop updates the shop in the database
func (p *Pg) UpdateShop(ctx context.Context, req UpdateShopRequest) error {
	execRes, err := p.gen.UpdateShop(ctx, generated.UpdateShopParams{
		ID:              req.ID,
		Name:            req.Name,
		OwnerExternalID: pgtype.Int8{Int64: req.ExternalUserID, Valid: true},
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.UpdateShop")
	}

	if execRes.RowsAffected() == 0 {
		return ErrorOpNotAllowed
	}

	return nil
}

// SoftDeleteShop marks the shop as deleted
func (p *Pg) SoftDeleteShop(ctx context.Context, req DeleteShopRequest) error {
	err := p.gen.SoftDeleteShop(ctx, req.WebAppId)
	if err != nil {
		return errors.Wrap(err, "p.gen.SoftDeleteMarketplace")
	}

	return nil
}

// ConfigureShopSync enables the shop synchronization
func (p *Pg) ConfigureShopSync(ctx context.Context, req ConfigureShopSyncRequest) error {
	ok, err := p.gen.IsShopSyncSupported(ctx, req.WebAppID)
	if err != nil {
		return errors.Wrap(err, "p.gen.IsShopSyncSupported")
	}
	if !ok {
		return ErrorShopSyncNotSupported
	}

	err = p.gen.EnableShopSync(ctx, generated.EnableShopSyncParams{
		WebAppID:         req.WebAppID,
		ApiKey:           req.APIKey,
		ExternalProvider: generated.ExternalProvider(req.ExternalProvider),
		IsActive:         req.IsActive,
	})
	if err != nil {
		return errors.Wrap(err, "pg.gen.ConfigureShopSync")
	}

	return nil
}

// CreateProduct saves a product to the database
// and returns the ShopID that it assigned to it
func (p *Pg) CreateProduct(ctx context.Context, req CreateProductRequest) (CreateProductResponse, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "pg.pool.Begin")
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)
	qtx := p.gen.WithTx(tx)

	count, err := qtx.CountMarketplaceProducts(ctx, req.WebAppID)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "p.gen.CountMarketplaceProducts")
	}

	if count > maxProducts {
		return CreateProductResponse{}, ErrorMaxProductsExceeded
	}

	id, err := qtx.CreateProduct(ctx, generated.CreateProductParams{
		WebAppID:    req.WebAppID,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
	})
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "p.gen.CreateProduct")
	}

	err = setProductExternalLinks(ctx, qtx, id, req.ExternalLinks)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "setProductExternalLinks")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return CreateProductResponse{}, errors.Wrap(err, "tx.Commit")
	}

	return CreateProductResponse{ID: id}, err
}

// UpdateProduct updates the product of a marketplace in the database
func (p *Pg) UpdateProduct(ctx context.Context, req UpdateProductRequest) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.pool.Begin")
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)
	qtx := p.gen.WithTx(tx)

	execRes, err := qtx.UpdateProduct(ctx, generated.UpdateProductParams{
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

	err = setProductExternalLinks(ctx, qtx, req.ID, req.ExternalLinks)
	if err != nil {
		return errors.Wrap(err, "setProductExternalLinks")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.Wrap(err, "tx.Commit")
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
		OwnerExternalID: pgtype.Int8{Int64: req.ExternalUserID, Valid: true},
		MarketplaceID:   req.ShopID,
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
		products := make([]OrderProduct, 0)
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
			Currency:      string(v.Currency),
			Type:          string(v.Type),
		}
	}

	return GetOrdersResponse{Orders: orders}, nil
}

// GetBalance returns balances in all currencies
func (p *Pg) GetBalance(ctx context.Context, req GetBalanceRequest) (GetBalanceResponse, error) {
	rows, err := p.gen.GetBalance(ctx, pgtype.Int8{Int64: req.ExternalUserID, Valid: true})
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

// IsShopOwner checks if the user is the owner of the shop
func (p *Pg) IsShopOwner(ctx context.Context, userID int64, webAppID uuid.UUID) (bool, error) {
	ok, err := p.gen.IsUserTheOwnerOfWebApp(ctx, generated.IsUserTheOwnerOfWebAppParams{
		OwnerExternalID: pgtype.Int8{Int64: userID, Valid: true},
		ID:              webAppID,
	})
	if err != nil {
		return false, errors.Wrap(err, "p.gen.IsUserTheOwnerOfWebApp")
	}

	return ok, nil
}

// IsTelegramChannelOwner checks if the user is the owner of the Telegram channel
func (p *Pg) IsTelegramChannelOwner(ctx context.Context, externalUserID, channelID int64) (bool, error) {
	ok, err := p.gen.IsUserTheOwnerOfTelegramChannel(ctx, generated.IsUserTheOwnerOfTelegramChannelParams{
		OwnerExternalID: externalUserID,
		ExternalID:      channelID,
	})
	if err != nil {
		return false, ErrorOpNotAllowed
	}

	return ok, nil
}

// GetTelegramChannels gets a user's Telegram channels
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

func setProductExternalLinks(
	ctx context.Context,
	qtx *generated.Queries,
	productID uuid.UUID,
	externalLinks []ProductExternalLink,
) error {
	// Remove existing external links for the product
	err := qtx.RemoveProductExternalLinks(ctx, productID)
	if err != nil {
		return errors.Wrap(err, "qtx.RemoveProductExternalLinks")
	}

	// Prepare new external links
	var extLinks []generated.SetProductExternalLinksParams
	for _, link := range externalLinks {
		extLinks = append(extLinks, generated.SetProductExternalLinksParams{
			ProductID: productID,
			Url:       link.URL,
			Label:     link.Label,
		})
	}

	// Insert new external links using batch execution
	var batchErr error
	br := qtx.SetProductExternalLinks(ctx, extLinks)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = errors.Wrap(err, "qtx.SetProductExternalLinks")
		}
	})

	if batchErr != nil {
		return batchErr
	}

	return nil
}
