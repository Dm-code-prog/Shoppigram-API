package webhooks

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
func (p *Pg) GetOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	order, err := p.gen.GetOrder(ctx, id)
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
		State:     string(order.State),
	}, nil
}

// SetOrderStateConfirmed sets the state of an order to 'confirmed'
func (p *Pg) SetOrderStateConfirmed(ctx context.Context, id uuid.UUID) error {
	err := p.gen.SetOrderStateConfirmed(ctx, id)
	if err != nil {
		return errors.Wrap(err, "p.gen.SetOrderStateConfirmed")
	}
	return nil
}

// SavePaymentExtraInfo saves the extra info about a payment
func (p *Pg) SavePaymentExtraInfo(ctx context.Context, params SavePaymentExtraInfoParams) error {
	err := p.gen.SavePaymentExtraInfo(ctx, generated.SavePaymentExtraInfoParams{
		OrderID:            pgtype.UUID{Bytes: params.OrderID, Valid: true},
		Provider:           generated.PaymentProviders(params.Provider),
		OrderStateSnapshot: generated.OrderState(params.OrderStateSnapshot),
		EventType:          generated.PaymentsEventType(params.EventType),
		ExtraInfo:          params.ExtraInfo,
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.SavePaymentExtraInfo(ctx, params)")
	}
	return nil
}
