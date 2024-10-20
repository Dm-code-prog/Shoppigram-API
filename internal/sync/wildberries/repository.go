package wildberries

import (
	"context"
	"github.com/google/uuid"
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

// SyncProducts replaces all the products of a shop from a specific external provider
func (pg *Pg) SyncProducts(ctx context.Context, params SetProductsParams) error {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.pool.Begin")
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)
	qtx := pg.gen.WithTx(tx)

	existingProductIDs, err := qtx.GetExternalIds(ctx, generated.GetExternalIdsParams{
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
	var productsToDelete []pgtype.Text
	for _, id := range existingProductIDs {
		if _, ok := productsIDsLookUpMap[id.String]; !ok {
			productsToDelete = append(productsToDelete, id)
		}
	}

	var batchDeleteProductsErr error
	brProductsDel := qtx.MarkProductAsDeleted(ctx, productsToDelete)
	brProductsDel.Exec(func(i int, err error) {
		if err != nil {
			batchDeleteProductsErr = err
		}
	})
	if batchDeleteProductsErr != nil {
		return errors.Wrap(batchDeleteProductsErr, "brProductsDel.Exec")
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

	var batchInsertProductsErr error
	brInsert := qtx.CreateOrUpdateProducts(ctx, insertParams)
	brInsert.Exec(func(i int, err error) {
		if err != nil {
			batchInsertProductsErr = err
		}
	})
	if batchInsertProductsErr != nil {
		return errors.Wrap(batchInsertProductsErr, "brInsert.Exec")
	}

	intProducts, err := qtx.GetProducts(ctx, generated.GetProductsParams{
		WebAppID: pgtype.UUID{
			Bytes: params.ShopID,
			Valid: true,
		},
		ExternalProvider: generated.NullExternalProvider{ExternalProvider: externalProvider, Valid: true},
	})
	if err != nil {
		return errors.Wrap(err, "qtx.GetProducts")
	}

	productIDsMap := make(map[string]uuid.UUID)
	for _, p := range intProducts {
		productIDsMap[p.ExternalID.String] = p.ID
	}

	// set external links
	var insertLinksParams []generated.CreateOrUpdateExternalLinksParams
	for _, p := range params.Products {
		for _, l := range p.ExternalLinks {
			if id, ok := productIDsMap[p.ExternalID]; ok {
				insertLinksParams = append(insertLinksParams, generated.CreateOrUpdateExternalLinksParams{
					ProductID: id,
					Url:       l.URL,
					Label:     l.Label,
				})
			}
		}
	}

	var batchInsertLinksErr error
	brInsertLinks := qtx.CreateOrUpdateExternalLinks(ctx, insertLinksParams)
	brInsertLinks.Exec(func(i int, err error) {
		if err != nil {
			batchInsertLinksErr = err
		}
	})
	if batchInsertLinksErr != nil {
		return errors.Wrap(batchInsertLinksErr, "brInsertLinks.Exec")
	}

	// set photos
	// delete photos of all products in the batch
	var toDelete []uuid.UUID
	for _, p := range intProducts {
		toDelete = append(toDelete, p.ID)
	}

	var batchDelPhotosErr error
	brPhotosDel := qtx.DeletePhotos(ctx, toDelete)
	brPhotosDel.Exec(func(i int, err error) {
		if err != nil {
			batchDelPhotosErr = err
		}
	})
	if batchDelPhotosErr != nil {
		return errors.Wrap(batchDeleteProductsErr, "brProductsDel.Exec")
	}

	var insertPhotosParams []generated.CreateOrUpdatePhotosParams
	for _, p := range params.Products {
		for _, photo := range p.Photos {
			if id, ok := productIDsMap[p.ExternalID]; ok {
				insertPhotosParams = append(insertPhotosParams, generated.CreateOrUpdatePhotosParams{
					ProductID: id,
					Url:       photo.URL,
				})
			}
		}
	}

	var batchInsertPhotosErr error
	brPhotosInsert := qtx.CreateOrUpdatePhotos(ctx, insertPhotosParams)
	brPhotosInsert.Exec(func(i int, err error) {
		if err != nil {
			batchInsertPhotosErr = err
		}
	})
	if batchInsertPhotosErr != nil {
		return errors.Wrap(batchInsertPhotosErr, "brPhotosInsert.Exec")
	}

	return tx.Commit(ctx)
}

// GetNextSyncJob returns the next shop to sync
func (pg *Pg) GetNextSyncJob(ctx context.Context) (*Job, error) {
	ns, err := pg.gen.GetNextShop(ctx)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return &Job{}, errors.Wrap(err, "pg.gen.GetNextSyncJob")
	}

	return &Job{
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
