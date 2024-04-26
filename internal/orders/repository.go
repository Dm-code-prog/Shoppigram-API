package orders

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/orders/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen *generated.Queries
	db  *pgx.Conn
}

// NewPg creates a new Pg
func NewPg(gen *generated.Queries) *Pg {
	return &Pg{gen: gen}
}

// CreateOrder adds a new order to the database
func (pg *Pg) CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	tx, err := pg.db.Begin(ctx)
	if err != nil {
		return CreateOrderResponse{}, err
	}
	defer tx.Rollback(ctx)
	qtx := pg.gen.WithTx(tx)

	orderID, err := qtx.CreateOrder(ctx, generated.CreateOrderParams{
		WebAppID: pgtype.UUID{
			Bytes: req.WebAppID,
			Valid: true,
		},
		TelegramUserID: pgtype.UUID{
			Bytes: req.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "qtx.CreateOrder")
	}

	var orderProducts []generated.SetOrderProductsParams
	for _, product := range req.Products {
		orderProducts = append(orderProducts, generated.SetOrderProductsParams{
			OrderID:   pgtype.UUID{Bytes: orderID, Valid: true},
			ProductID: pgtype.UUID{Bytes: product.ID, Valid: true},
			Quantity:  product.Quantity,
		})
	}

	br := qtx.SetOrderProducts(ctx, orderProducts)
	br.Exec(nil)
	err = tx.Commit(ctx)
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "qtx.SetOrderProducts")
	}

	return CreateOrderResponse{ID: orderID}, nil
}
