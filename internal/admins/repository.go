package admins

import (
	"context"
	"github.com/jackc/pgerrcode"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/admins/generated"
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
	marketplaces := make([]Marketplace, 0)

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

	for _, v := range rows {
		marketplaces = append(marketplaces, Marketplace{
			ID:         v.ID,
			Name:       v.Name,
			LogoURL:    v.LogoUrl.String,
			IsVerified: v.IsVerified.Bool,
		})
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
		WebAppID:      req.WebAppID,
		Name:          req.Name,
		Price:         req.Price,
		PriceCurrency: req.PriceCurrency,
		Description:   req.Description,
		Category:      req.Category,
	})
	if err != nil {
		if strings.Contains(err.Error(), pgerrcode.InvalidTextRepresentation) {
			return CreateProductResponse{}, ErrorInvalidProductCurrency
		}

		return CreateProductResponse{}, errors.Wrap(err, "p.gen.CreateProduct")
	}

	return CreateProductResponse{ID: id}, err
}
