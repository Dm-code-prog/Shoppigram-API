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

	// Start by getting the external IDs of products we have before the sync
	currentExternalIDs, err := qtx.GetExternalIds(ctx, generated.GetExternalIdsParams{
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

	// Create a map of current external IDs.
	// This way we can find out if a product is already present in the database
	// And find products that are not in the new list of products anymore
	// and therefore need to be deleted
	productsIDsLookUpMap := make(map[string]struct{})
	for _, p := range params.Products {
		productsIDsLookUpMap[p.ExternalID] = struct{}{}
	}

	// Products

	// If a product used to be in the database but is not in the new list of products
	// we need to delete it
	var productsToDelete []pgtype.Text
	for _, id := range currentExternalIDs {
		if _, ok := productsIDsLookUpMap[id.String]; !ok {
			productsToDelete = append(productsToDelete, id)
		}
	}

	// Delete products that are not in the new list of products
	// in a batch operation.
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
	// Finished deleting products

	// Now we need to insert the new products and update the existing ones
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

	// insert the new products in a batch operation
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
	// Finished inserting the new list of products

	// After the insert we get the internal IDs of all products that are to be found in the database
	// We need them for creating or updating matching records such as photos or external links.
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

	// We create a map with:
	// External ID as key
	// Internal ID as value
	productIDsMap := make(map[string]uuid.UUID)
	for _, p := range intProducts {
		productIDsMap[p.ExternalID.String] = p.ID
	}

	// External links

	// First we delete all external links of all products in the batch
	var linksToDelete []uuid.UUID
	for _, p := range intProducts {
		linksToDelete = append(linksToDelete, p.ID)
	}
	var batchDeleteLinksErr error
	brDeleteLinks := qtx.DeleteExternalLinks(ctx, linksToDelete)
	brDeleteLinks.Exec(func(i int, err error) {
		if err != nil {
			batchDeleteLinksErr = err
		}
	})
	if batchDeleteLinksErr != nil {
		return errors.Wrap(batchDeleteLinksErr, "brDeleteLinks.Exec")
	}

	// Now we insert the new external links
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

	// Insert the new external links in a batch operation
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
	// Finished inserting external links

	// Photos

	// The same as with external links
	// First we delete all photos of all products in the batch
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

	// Now we insert the new photos
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

	// If all the operations succeeded we commit the transaction
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
