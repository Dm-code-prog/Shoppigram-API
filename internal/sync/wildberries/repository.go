package wildberries

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/sync/wildberries/generated"
	"math/big"
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
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	qtx := pg.gen.WithTx(tx)

	err = pg.syncProductsWithTx(ctx, qtx, params)
	if err != nil {
		return errors.Wrap(err, "pg.syncProductsWithTx")
	}

	return tx.Commit(ctx)
}

func (pg *Pg) syncProductsWithTx(ctx context.Context, qtx *generated.Queries, params SetProductsParams) error {
	currentExternalIDs, err := pg.getCurrentExternalIDs(ctx, qtx, params)
	if err != nil {
		return err
	}

	productsToDelete := pg.identifyProductsToDelete(currentExternalIDs, params)

	// Delete products not in the new list
	if err := pg.deleteProducts(ctx, qtx, productsToDelete); err != nil {
		return err
	}

	// Insert or update products
	if err := pg.insertOrUpdateProducts(ctx, qtx, params); err != nil {
		return err
	}

	// Get internal product IDs after insertion/update
	intProducts, err := pg.getInternalProductIDs(ctx, qtx, params)
	if err != nil {
		return err
	}

	// Create a map of external ID to internal ID
	productIDsMap := pg.createProductIDsMap(intProducts)

	// Update product variants
	//
	// Note:
	// In this implementation we first delete all the product variants
	// and then insert the new ones. I am not sure if this is the way to go.
	//
	// One thing to keep in mind is that after each sync, the IDs of variants will be reset.
	if err := pg.updateProductVariants(ctx, qtx, params, intProducts, productIDsMap); err != nil {
		return err
	}

	// Delete and insert external links
	if err := pg.updateExternalLinks(ctx, qtx, params, intProducts, productIDsMap); err != nil {
		return err
	}

	// Delete and insert photos
	if err := pg.updatePhotos(ctx, qtx, params, intProducts, productIDsMap); err != nil {
		return err
	}

	return nil
}

func (pg *Pg) getCurrentExternalIDs(ctx context.Context, qtx *generated.Queries, params SetProductsParams) ([]pgtype.Text, error) {
	currentExternalIDs, err := qtx.GetExternalIds(ctx, generated.GetExternalIdsParams{
		WebAppID: pgtype.UUID{Bytes: params.ShopID, Valid: true},
		ExternalProvider: generated.NullExternalProvider{
			ExternalProvider: generated.ExternalProvider(params.ExternalProvider),
			Valid:            true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "qtx.GetExternalIds")
	}
	return currentExternalIDs, nil
}

func (pg *Pg) updateProductVariants(ctx context.Context, qtx *generated.Queries, params SetProductsParams, intProducts []generated.Product, productIDsMap map[string]uuid.UUID) error {
	// Delete existing product variants
	var productIDs []uuid.UUID
	for _, p := range intProducts {
		productIDs = append(productIDs, p.ID)
	}
	if err := pg.deleteProductVariants(ctx, qtx, productIDs); err != nil {
		return err
	}

	// Insert new product variants
	if err := pg.insertProductVariants(ctx, qtx, params, productIDsMap); err != nil {
		return err
	}

	return nil
}

func (pg *Pg) deleteProductVariants(ctx context.Context, qtx *generated.Queries, productIDs []uuid.UUID) error {
	if len(productIDs) == 0 {
		return nil
	}
	var batchErr error
	br := qtx.DeleteProductVariants(ctx, productIDs)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) insertProductVariants(ctx context.Context, qtx *generated.Queries, params SetProductsParams, productIDsMap map[string]uuid.UUID) error {
	var insertParams []generated.CreateOrUpdateProductVariantsParams
	for _, p := range params.Products {
		productID, ok := productIDsMap[p.ExternalID]
		if !ok {
			continue
		}
		for _, v := range p.Variants {
			var (
				dimensionsJSONb []byte
				err             error
			)
			if v.Dimensions != nil {
				dimensionsJSONb, err = json.Marshal(v.Dimensions)
				if err != nil {
					return errors.Wrap(err, "json.Marshal")
				}
			}

			insertParams = append(insertParams, generated.CreateOrUpdateProductVariantsParams{
				ProductID:       productID,
				Price:           pgtype.Numeric{Int: big.NewInt(int64(v.Price * 100)), Valid: true},
				DiscountedPrice: pgtype.Numeric{Int: big.NewInt(int64(v.DiscountedPrice * 100)), Valid: true},
				Dimensions:      dimensionsJSONb,
			})
		}
	}

	if len(insertParams) == 0 {
		return nil
	}

	var batchErr error
	br := qtx.CreateOrUpdateProductVariants(ctx, insertParams)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) identifyProductsToDelete(currentExternalIDs []pgtype.Text, params SetProductsParams) []pgtype.Text {
	productsIDsLookUpMap := make(map[string]struct{})
	for _, p := range params.Products {
		productsIDsLookUpMap[p.ExternalID] = struct{}{}
	}

	var productsToDelete []pgtype.Text
	for _, id := range currentExternalIDs {
		if _, ok := productsIDsLookUpMap[id.String]; !ok {
			productsToDelete = append(productsToDelete, id)
		}
	}
	return productsToDelete
}

func (pg *Pg) deleteProducts(ctx context.Context, qtx *generated.Queries, productsToDelete []pgtype.Text) error {
	if len(productsToDelete) == 0 {
		return nil
	}
	var batchErr error
	br := qtx.MarkProductAsDeleted(ctx, productsToDelete)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) insertOrUpdateProducts(ctx context.Context, qtx *generated.Queries, params SetProductsParams) error {
	var insertParams []generated.CreateOrUpdateProductsParams
	for _, p := range params.Products {
		insertParams = append(insertParams, generated.CreateOrUpdateProductsParams{
			WebAppID: pgtype.UUID{Bytes: params.ShopID, Valid: true},
			Name:     p.Name,
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

	var batchErr error
	br := qtx.CreateOrUpdateProducts(ctx, insertParams)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) getInternalProductIDs(ctx context.Context, qtx *generated.Queries, params SetProductsParams) ([]generated.Product, error) {
	intProducts, err := qtx.GetProductIDs(ctx, generated.GetProductIDsParams{
		WebAppID: pgtype.UUID{Bytes: params.ShopID, Valid: true},
		ExternalProvider: generated.NullExternalProvider{
			ExternalProvider: generated.ExternalProvider(params.ExternalProvider),
			Valid:            true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "qtx.GetProducts")
	}

	var products []generated.Product
	for _, p := range intProducts {
		products = append(products, generated.Product{
			ID:         p.ID,
			ExternalID: p.ExternalID,
		})
	}

	return products, nil
}

func (pg *Pg) createProductIDsMap(intProducts []generated.Product) map[string]uuid.UUID {
	productIDsMap := make(map[string]uuid.UUID)
	for _, p := range intProducts {
		productIDsMap[p.ExternalID.String] = p.ID
	}
	return productIDsMap
}

func (pg *Pg) updateExternalLinks(ctx context.Context, qtx *generated.Queries, params SetProductsParams, intProducts []generated.Product, productIDsMap map[string]uuid.UUID) error {
	// Delete existing external links
	var productIDs []uuid.UUID
	for _, p := range intProducts {
		productIDs = append(productIDs, p.ID)
	}
	if err := pg.deleteExternalLinks(ctx, qtx, productIDs); err != nil {
		return err
	}

	// Insert new external links
	if err := pg.insertExternalLinks(ctx, qtx, params, productIDsMap); err != nil {
		return err
	}

	return nil
}

func (pg *Pg) deleteExternalLinks(ctx context.Context, qtx *generated.Queries, productIDs []uuid.UUID) error {
	if len(productIDs) == 0 {
		return nil
	}
	var batchErr error
	br := qtx.DeleteExternalLinks(ctx, productIDs)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) insertExternalLinks(ctx context.Context, qtx *generated.Queries, params SetProductsParams, productIDsMap map[string]uuid.UUID) error {
	var insertParams []generated.CreateOrUpdateExternalLinksParams
	for _, p := range params.Products {
		productID, ok := productIDsMap[p.ExternalID]
		if !ok {
			continue
		}
		for _, l := range p.ExternalLinks {
			insertParams = append(insertParams, generated.CreateOrUpdateExternalLinksParams{
				ProductID: productID,
				Url:       l.URL,
				Label:     l.Label,
			})
		}
	}

	if len(insertParams) == 0 {
		return nil
	}

	var batchErr error
	br := qtx.CreateOrUpdateExternalLinks(ctx, insertParams)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) updatePhotos(ctx context.Context, qtx *generated.Queries, params SetProductsParams, intProducts []generated.Product, productIDsMap map[string]uuid.UUID) error {
	// Delete existing photos
	var productIDs []uuid.UUID
	for _, p := range intProducts {
		productIDs = append(productIDs, p.ID)
	}
	if err := pg.deletePhotos(ctx, qtx, productIDs); err != nil {
		return err
	}

	// Insert new photos
	if err := pg.insertPhotos(ctx, qtx, params, productIDsMap); err != nil {
		return err
	}

	return nil
}

func (pg *Pg) deletePhotos(ctx context.Context, qtx *generated.Queries, productIDs []uuid.UUID) error {
	if len(productIDs) == 0 {
		return nil
	}
	var batchErr error
	br := qtx.DeletePhotos(ctx, productIDs)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
}

func (pg *Pg) insertPhotos(ctx context.Context, qtx *generated.Queries, params SetProductsParams, productIDsMap map[string]uuid.UUID) error {
	var insertParams []generated.CreateOrUpdatePhotosParams
	for _, p := range params.Products {
		productID, ok := productIDsMap[p.ExternalID]
		if !ok {
			continue
		}
		for _, photo := range p.Photos {
			insertParams = append(insertParams, generated.CreateOrUpdatePhotosParams{
				ProductID: productID,
				Url:       photo.URL,
			})
		}
	}

	if len(insertParams) == 0 {
		return nil
	}

	var batchErr error
	br := qtx.CreateOrUpdatePhotos(ctx, insertParams)
	br.Exec(func(i int, err error) {
		if err != nil {
			batchErr = err
		}
	})
	if batchErr != nil {
		return errors.Wrap(batchErr, "br.Exec")
	}
	return nil
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
