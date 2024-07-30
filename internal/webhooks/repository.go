package webhooks

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/webhooks/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen *generated.Queries
}

// NewPg takes the reference to pgxpool and returns the new instance of repository interface
func NewPg(db *pgxpool.Pool) *Pg {
	return &Pg{
		gen: generated.New(db),
	}
}

// GetOrder takes the id of an order and returns this order's data
func (p *Pg) GetOrder(ctx context.Context, id string) (Order, error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return Order{}, errors.Wrap(err, "uuid.Parse(id). id = "+id)
	}
	order, err := p.gen.GetOrder(ctx, idParsed)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Order{}, ErrorOrderDoesntExist
		}
		return Order{}, errors.Wrap(err, "p.gen.GetOrder(ctx, idParsed)")
	}
	parsedTime := order.UpdatedAt.Time

	return Order{
		ID:        order.ID,
		UpdatedAt: parsedTime,
		Sum:       order.OrderSum,
		Currency:  order.PriceCurrency,
	}, nil
}
