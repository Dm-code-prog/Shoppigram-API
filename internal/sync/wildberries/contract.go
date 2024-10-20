package wildberries

import (
	"context"
	"github.com/google/uuid"
	"time"
)

const (
	syncInterval         = time.Hour
	failureRetryInterval = time.Hour * 3
)

type (
	// ExternalLink is a link to a product
	// on an external website
	//
	// Used for panels
	ExternalLink struct {
		URL   string
		Label string
	}

	// ExternalProduct defines the structure for a product
	ExternalProduct struct {
		ExternalID    string
		Name          string
		Description   string
		Category      string
		Price         float64
		ExternalLinks []ExternalLink
	}

	// Product is a product in the database
	Product struct {
		ID          uuid.UUID
		ExternalID  string
		Name        string
		Price       float64
		Category    string
		Description string
	}
)

type (
	// SetProductsParams is the stucture that holds params for
	// Repository.SetExternalProducts method
	SetProductsParams struct {
		ShopID           uuid.UUID
		ExternalProvider string
		Products         []ExternalProduct
	}

	// GetProductsParams is the structure that holds params for
	// Repository.GetProducts method
	GetProductsParams struct {
		ShopID           uuid.UUID
		ExternalProvider string
	}

	// SetSyncSuccessParams is the structure that holds params for
	// Repository.SetSyncSuccess method
	SetSyncSuccessParams struct {
		JobID uuid.UUID
	}

	// SetSyncFailureParams is the structure that holds params for
	// Repository.SetSyncFailure method
	SetSyncFailureParams struct {
		JobID     uuid.UUID
		LastError string
	}

	// NextShop is the structure that holds the next shop to sync
	NextShop struct {
		ShopID    uuid.UUID
		SyncJobID uuid.UUID
		APIKey    string
	}

	Repository interface {
		// SetExternalProducts replaces all the products of a shop from a specific external provider
		// with a new list of products
		SetExternalProducts(context.Context, SetProductsParams) error

		// GetProducts returns all the products of a shop from a specific external provider
		GetProducts(context.Context, GetProductsParams) ([]Product, error)

		// GetNextShop returns the next shop to sync
		GetNextShop(context.Context) (NextShop, error)

		// SetSyncSuccess marks a sync job as successful
		SetSyncSuccess(context.Context, SetSyncSuccessParams) error

		// SetSyncFailure marks a sync job as failed
		SetSyncFailure(context.Context, SetSyncFailureParams) error
	}
)
