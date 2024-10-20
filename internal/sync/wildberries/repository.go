package wildberries

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/sync/wildberries/generated"
)

// Pg implements the Repository interface
// using PostgresSQL as the backing store.
type Pg struct {
	gen  *generated.Queries
	pool *pgxpool.Pool
}

// NewPg returns new Pg
func NewPg(pool *pgxpool.Pool) *Pg {
	return &Pg{gen: generated.New(pool), pool: pool}
}

// SetExternalProducts replaces all the products of a shop from a specific external provider
func (pg *Pg) SetExternalProducts(ctx context.Context, params SetProductsParams) error {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.pool.Begin")
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)
	qtx := pg.gen.WithTx(tx)

	existingIDsDB, err := qtx.GetExternalIds(ctx, generated.GetExternalIdsParams{
		WebAppID: pgtype.UUID{
			Bytes: params.ShopID,
			Valid: true,
		},
		ExternalProvider: generated.NullExternalProvider{
			ExternalProvider: generated.ExternalProvider(params.ExternalProvider),
			Valid:            true,
		},
	})
	if err != nil {
		return errors.Wrap(err, "qtx.GetExternalIds")
	}

	productsIDsLookUpMap := make(map[string]struct{})
	for _, p := range params.Products {
		productsIDsLookUpMap[p.ExternalID] = struct{}{}
	}

	// If a product used to be in the database but is not in the new list of products
	// we need to delete it
	var toDelete []pgtype.Text
	for _, id := range existingIDsDB {
		if _, ok := productsIDsLookUpMap[id.String]; !ok {
			toDelete = append(toDelete, id)
		}
	}

	var batchDeleteErr error
	brDel := qtx.MarkProductAsDeleted(ctx, toDelete)
	brDel.Exec(func(i int, err error) {
		if err != nil {
			batchDeleteErr = err
		}
	})
	if batchDeleteErr != nil {
		return errors.Wrap(batchDeleteErr, "brDel.Exec")
	}

	var insertParams []generated.CreateOrUpdateProductsParams
	for _, p := range params.Products {
		insertParams = append(insertParams, generated.CreateOrUpdateProductsParams{
			WebAppID: pgtype.UUID{
				Bytes: params.ShopID,
				Valid: true,
			},
			Name: p.Name,
			Description: pgtype.Text{
				String: p.Description,
				Valid:  true,
			},
			Price: p.Price,
			Category: pgtype.Text{
				String: p.Category,
				Valid:  true,
			},
			ExternalProvider: generated.NullExternalProvider{
				ExternalProvider: generated.ExternalProvider(params.ExternalProvider),
				Valid:            true,
			},
			ExternalID: pgtype.Text{
				String: p.ExternalID,
				Valid:  true,
			},
		})
	}

	var batchInsertErr error
	brInsert := qtx.CreateOrUpdateProducts(ctx, insertParams)
	brInsert.Exec(func(i int, err error) {
		if err != nil {
			batchInsertErr = err
		}
	})
	if batchInsertErr != nil {
		return errors.Wrap(batchInsertErr, "brInsert.Exec")
	}

	return tx.Commit(ctx)
}

// GetProducts returns all the products of a shop from a specific external provider
func (pg *Pg) GetProducts(ctx context.Context, params GetProductsParams) ([]Product, error) {
	res, err := pg.gen.GetProducts(ctx, generated.GetProductsParams{
		WebAppID: pgtype.UUID{
			Bytes: params.ShopID,
			Valid: true,
		},
		ExternalProvider: generated.NullExternalProvider{
			ExternalProvider: generated.ExternalProvider(params.ExternalProvider),
			Valid:            true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "pg.gen.GetProducts")
	}

	var products []Product
	for _, p := range res {
		products = append(products, Product{
			ID:          p.ID,
			ExternalID:  p.ExternalID.String,
			Name:        p.Name,
			Price:       p.Price,
			Category:    p.Category.String,
			Description: p.Description.String,
		})
	}

	return products, nil
}

// GetNextShop returns the next shop to sync
func (pg *Pg) GetNextShop(ctx context.Context) (NextShop, error) {
	ns, err := pg.gen.GetNextShop(ctx, generated.GetNextShopParams{
		SyncInterval:         syncInterval,
		FailureRetryInterval: failureRetryInterval,
	})
	if err != nil {
		return NextShop{}, errors.Wrap(err, "pg.gen.GetNextShop")
	}

	return NextShop{
		ShopID:    ns.WebAppID,
		SyncJobID: ns.ID,
		APIKey:    ns.ApiKey,
	}, nil
}

// SetSyncSuccess marks the sync job as successful
func (pg *Pg) SetSyncSuccess(ctx context.Context, params SetSyncSuccessParams) error {
	err := pg.gen.SetSyncSuccess(ctx, params.JobID)
	if err != nil {
		return errors.Wrap(err, "pg.gen.SetSyncSuccess")
	}

	return nil
}

// SetSyncFailure marks the sync job as failed
func (pg *Pg) SetSyncFailure(ctx context.Context, params SetSyncFailureParams) error {
	err := pg.gen.SetSyncFailure(ctx, generated.SetSyncFailureParams{
		ID:        params.JobID,
		LastError: pgtype.Text{String: params.LastError, Valid: true},
	})
	if err != nil {
		return errors.Wrap(err, "pg.gen.SetSyncFailure")
	}

	return nil
}
