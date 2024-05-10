package admins

import (
	"context"

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
	var marketplaces []Marketplace

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
