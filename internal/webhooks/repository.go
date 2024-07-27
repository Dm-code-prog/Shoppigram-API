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

func NewPg(db *pgxpool.Pool) *Pg {
	return &Pg{
		gen: generated.New(db),
	}
}

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
		return Order{}, errors.Wrap(err, "PG error")
	}

	return Order{
		ID:        order.ID,
		UpdatedAt: order.UpdatedAt,
		Sum:       order.Sum,
		Currency:  order.PriceCurrency,
	}, nil
}
