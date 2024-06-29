package marketplaces

import (
	"context"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/marketplaces/generated"
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

// GetProducts returns a list of products
func (p *Pg) GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error) {
	prod, err := p.gen.GetProducts(ctx, request.WebAppID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetProductsResponse{}, errors.Wrap(ErrorProductsNotFound, "p.gen.GetProducts")
		}
		return GetProductsResponse{}, errors.Wrap(err, "p.gen.GetProducts")
	}

	var id uuid.UUID
	var name string
	var shortName string
	var isVerified bool
	var products []Product
	for _, p := range prod {
		products = append(products, Product{
			ID:            p.ID,
			Name:          p.Name,
			Description:   p.Description.String,
			Category:      p.Category.String,
			Price:         p.Price,
			PriceCurrency: p.PriceCurrency,
		})
		id = p.WebAppID
		name = p.WebAppName
		shortName = p.WebAppShortName
		isVerified = p.WebAppIsVerified.Bool
	}

	return GetProductsResponse{
		WebAppID:         id,
		WebAppName:       name,
		WebAppShortName:  shortName,
		WebAppIsVerified: isVerified,
		Products:         products,
	}, nil
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
