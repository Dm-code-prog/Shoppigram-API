package admins

import (
	"context"

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
func NewPg(db *pgxpool.Pool, encryptionKey string) *Pg {
	return &Pg{gen: generated.New(db)}
}

// GetMarketplacesByUserID gets all user-related marketplaces
func (p *Pg) GetMarketplacesByUserID(ctx context.Context, userID int32) (interface{}, error) {
	val, err := p.gen.GetMarketplacesByUserID(ctx, pgtype.Int4{
		Int32: int32(userID),
		Valid: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetMarketplacesByUserID")
	}

	return val, nil
}
