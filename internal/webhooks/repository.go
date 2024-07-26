package webhooks

import (
	"context"

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

func (p *Pg) GetOrder(ctx context.Context, invoiceId string) (Order, error) {
	order, err := p.gen.GetOrder(ctx, invoiceId)
	if err != nil {
		return Order{}, errors.Wrap(err, "p.gen.GetOrder(id). InvoiceId = "+invoiceId)
	}

	return Order{
		ID:        order.ID,
		UpdatedAt: order.UpdatedAt,
		Sum:       order.Sum,
	}, nil
}
