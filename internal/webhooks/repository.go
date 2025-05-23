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
		InvoiceID: pgtype.UUID{Bytes: params.InvoiceID, Valid: true},
		Provider:  generated.PaymentProviders(params.Provider),
		EventType: generated.PaymentsEventType(params.EventType),
		ExtraInfo: params.ExtraInfo,
		Response:  params.Response,
		Error:     pgtype.Text{String: params.Error, Valid: true},
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.SavePaymentExtraInfo(ctx, params)")
	}
	return nil
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

// DeleteTelegramChannel deletes a Telegram channel
func (p *Pg) DeleteTelegramChannel(ctx context.Context, req DeleteTelegramChannelRequest) error {
	err := p.gen.DeleteTelegramChannel(ctx, req.ExternalID)
	if err != nil {
		return errors.Wrap(err, "p.gen.DeleteTelegramChannel")
	}

	return nil
}
