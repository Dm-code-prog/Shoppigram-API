package wildberries

import (
	"context"
	"github.com/google/uuid"
	"time"
)

const (
	externalProvider  = "wildberries"
	supportedCurrency = "rub"

	// default configuration for the sync process
	timeout = time.Second * 10
)

var (
	fetchLimit int32 = 100
)

type (
	Photo struct {
		URL string
	}

	// ExternalLink is a link to a product
	// on an external website
	//
	// Used for panels
	ExternalLink struct {
		URL   string
		Label string
	}

	// Variant is a product variant
	Variant struct {
		ID              uuid.UUID
		Dimensions      map[string]string
		Price           float64
		DiscountedPrice float64
	}

	// Product defines the structure for a product
	Product struct {
		ExternalID    string
		Name          string
		Description   string
		Category      string
		Price         float64
		Variants      []Variant
		ExternalLinks []ExternalLink
		Photos        []Photo
	}
)

type (
	// SetProductsParams is the stucture that holds params for
	// Repository.SetExternalProducts method
	SetProductsParams struct {
		ShopID           uuid.UUID
		ExternalProvider string
		Products         []Product
	}

	// ProductPhotos is the structure that holds the product id and
	// the photos of the product
	ProductPhotos struct {
		ProductID uuid.UUID
		Photos    []Photo
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

	// Job is the structure that holds the next shop to sync
	Job struct {
		ShopID    uuid.UUID
		SyncJobID uuid.UUID
		APIKey    string
	}

	Repository interface {
		// GetNextSyncJob returns the next shop to sync
		GetNextSyncJob(context.Context) (*Job, error)

		// SyncProducts replaces all the products of a shop from a specific external provider
		// with a new list of products
		SyncProducts(context.Context, SetProductsParams) error

		// SetSyncSuccess marks a sync job as successful
		SetSyncSuccess(context.Context, SetSyncSuccessParams) error

		// SetSyncFailure marks a sync job as failed
		SetSyncFailure(context.Context, SetSyncFailureParams) error
	}
)
