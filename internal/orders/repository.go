package orders

import (
	"context"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/orders/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen  *generated.Queries
	pool *pgxpool.Pool
}

// NewPg creates a new Pg
func NewPg(pool *pgxpool.Pool) *Pg {
	return &Pg{gen: generated.New(pool), pool: pool}
}

// CreateOrder adds a new order to the database
func (pg *Pg) CreateOrder(ctx context.Context, req SaveOrderRequest) (SaveOrderResponse, error) {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		return SaveOrderResponse{}, errors.Wrap(err, "pg.pool.Begin")
	}
	defer tx.Rollback(ctx)
	qtx := pg.gen.WithTx(tx)

	res, err := qtx.CreateOrder(ctx, generated.CreateOrderParams{
		WebAppID: pgtype.UUID{
			Bytes: req.WebAppID,
			Valid: true,
		},
		ExternalUserID: pgtype.Int4{
			Int32: int32(req.ExternalUserID),
			Valid: true,
		},
	})
	if err != nil {
		return SaveOrderResponse{}, errors.Wrap(err, "qtx.CreateOrder")
	}

	var orderProducts []generated.SetOrderProductsParams
	for _, product := range req.Products {
		orderProducts = append(orderProducts, generated.SetOrderProductsParams{
			OrderID:   pgtype.UUID{Bytes: res.ID, Valid: true},
			ProductID: pgtype.UUID{Bytes: product.ID, Valid: true},
			Quantity:  product.Quantity,
		})
	}

	var batchErr error
	br := qtx.SetOrderProducts(ctx, orderProducts)
	br.Exec(func(i int, err error) {
		if err != nil {
			if strings.Contains(err.Error(), pgerrcode.CheckViolation) {
				batchErr = ErrorInvalidProductQuantity
				return
			}

			batchErr = errors.Wrap(err, "br.Exec")
		}
	})
	if batchErr != nil {
		return SaveOrderResponse{}, batchErr
	}

	err = tx.Commit(ctx)
	if err != nil {
		return SaveOrderResponse{}, errors.Wrap(err, "tx.Commit")
	}

	return SaveOrderResponse{ReadableID: int(res.ReadableID.Int64)}, nil
}
