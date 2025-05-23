package app

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/app/generated"
)

const (
	stateDone = "done"
)

// Pg implements the Repository interface
// using PostgresSQL as the backing store.
type Pg struct {
	gen  *generated.Queries
	pool *pgxpool.Pool
}

// NewPg creates a new Pg
func NewPg(pool *pgxpool.Pool) *Pg {
	return &Pg{gen: generated.New(pool), pool: pool}
}

// GetShop returns a list of products
func (pg *Pg) GetShop(ctx context.Context, request GetShopRequest) (GetShopResponse, error) {
	m, err := pg.gen.GetMarketplaceWithProducts(ctx, generated.GetMarketplaceWithProductsParams{
		WebAppID:  request.WebAppID,
		ShortName: request.WebAppShortName,
	})
	if err != nil {
		return GetShopResponse{}, errors.Wrap(err, "pg.gen.GetMarketplaceWithProducts")
	}

	var products []Product
	if m.Products != nil {
		err = json.Unmarshal(m.Products, &products)
		if err != nil {
			return GetShopResponse{}, errors.Wrap(err, "json.Unmarshal")
		}
	}

	return GetShopResponse{
		WebAppID:              m.ID,
		WebAppName:            m.Name,
		WebAppShortName:       m.ShortName,
		WebAppType:            string(m.Type),
		WebAppIsVerified:      m.IsVerified.Bool,
		OnlinePaymentsEnabled: m.OnlinePaymentsEnabled,
		Products:              products,
		Currency:              string(m.Currency),
	}, nil
}

// CreateOrder adds a new order to the database
func (pg *Pg) CreateOrder(ctx context.Context, req SaveOrderParams) (SaveOrderResult, error) {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		return SaveOrderResult{}, errors.Wrap(err, "pg.pool.Begin")
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)
	qtx := pg.gen.WithTx(tx)

	ok, err := qtx.DoesWebAppSupportOrders(ctx, req.WebAppID)
	if err != nil {
		return SaveOrderResult{}, errors.Wrap(err, "qtx.DoesWebAppSupportOrders")
	}

	if !ok {
		return SaveOrderResult{}, ErrorOrderIsNotSupported
	}
	var (
		id         uuid.UUID
		readableID int
	)

	if req.Type == orderTypeOnline {
		res, err := qtx.CreateOnlineOrder(ctx, generated.CreateOnlineOrderParams{
			WebAppID: pgtype.UUID{
				Bytes: req.WebAppID,
				Valid: true,
			},
			ExternalUserID: pgtype.Int8{
				Int64: req.ExternalUserID,
				Valid: true,
			},
		})

		if err != nil {
			return SaveOrderResult{}, errors.Wrap(err, "qtx.CreateOnlineOrder")
		}

		id = res.ID
		readableID = int(res.ReadableID.Int64)
	} else if req.Type == orderTypeP2P {
		res, err := qtx.CreateP2POrder(ctx, generated.CreateP2POrderParams{
			WebAppID: pgtype.UUID{
				Bytes: req.WebAppID,
				Valid: true,
			},
			ExternalUserID: pgtype.Int8{
				Int64: req.ExternalUserID,
				Valid: true,
			},
		})
		if err != nil {
			return SaveOrderResult{}, errors.Wrap(err, "qtx.CreateOrder")
		}

		id = res.ID
		readableID = int(res.ReadableID.Int64)
	} else {
		return SaveOrderResult{}, ErrorInvalidOrderType
	}

	var orderProducts []generated.SetOrderProductsParams
	for _, product := range req.Products {
		orderProducts = append(orderProducts, generated.SetOrderProductsParams{
			OrderID:   pgtype.UUID{Bytes: id, Valid: true},
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
		return SaveOrderResult{}, batchErr
	}

	orderAmount, err := qtx.GetOrderAmount(ctx, id)
	if err != nil {
		return SaveOrderResult{}, errors.Wrap(err, "pg.gen.GetOrderAmount")
	}

	if orderAmount == 0 {
		err = qtx.UpdateOrderState(ctx, generated.UpdateOrderStateParams{
			ID:    id,
			State: stateDone,
		})
		if err != nil {
			return SaveOrderResult{}, errors.Wrap(err, "pg.gen.UpdateOrderState")
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return SaveOrderResult{}, errors.Wrap(err, "tx.Commit")
	}

	return SaveOrderResult{ReadableID: readableID, ID: id}, nil
}

// GetOrder gets a list of products in order
func (pg *Pg) GetOrder(ctx context.Context, orderId uuid.UUID, userId int64) (GetOrderResponse, error) {
	rows, err := pg.gen.GetOrder(ctx, generated.GetOrderParams{
		ID:             orderId,
		ExternalUserID: pgtype.Int8{Int64: userId, Valid: userId != 0},
	})

	if err != nil {
		return GetOrderResponse{}, errors.Wrap(err, "pg.gen.GetOrder")
	}

	products := make([]Product, len(rows))

	if len(rows) == 0 {
		return GetOrderResponse{}, ErrorGetOrderNotPermitted
	}
	WebAppName := rows[0].WebAppName
	WebAppShortName := rows[0].WebAppShortName

	var (
		totalPrice     float64
		readableID     int
		sellerUsername string
		currency       string
	)

	for i, v := range rows {
		products[i] = Product{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description.String,
			Category:    v.Category.String,
			BasePrice:   v.Price,
			Quantity:    v.Quantity,
		}

		totalPrice += v.Price * float64(v.Quantity)
		readableID = int(v.ReadableID.Int64)
		sellerUsername = v.SellerUsername.String
		currency = v.PriceCurrency
	}

	return GetOrderResponse{
		Products:        products,
		WebAppName:      WebAppName,
		WebAppShortName: WebAppShortName,
		TotalPrice:      totalPrice,
		ReadableOrderID: readableID,
		SellerUsername:  sellerUsername,
		Currency:        currency,
	}, nil
}
