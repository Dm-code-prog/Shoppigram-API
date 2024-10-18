package wildberries

import (
	"context"
	"github.com/google/uuid"
	"time"
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

	// GetCursorParams is the structure that holds params for
	// Repository.GetCursor method
	GetCursorParams struct {
		Name string
	}

	// ResetCursorParams is the structure that holds params for
	// Repository.ResetCursor method
	ResetCursorParams struct {
		Name string
	}

	// Cursor points to a specific record in a database
	// used for tracking state
	Cursor struct {
		Name      string
		ID        uuid.UUID
		Timestamp time.Time
	}

	// NextShop is the structure that holds the next shop to sync
	NextShop struct {
		ID              uuid.UUID
		CursorTimestamp time.Time
		APIKey          string
	}

	Repository interface {
		// SetExternalProducts replaces all the products of a shop from a specific external provider
		// with a new list of products
		SetExternalProducts(context.Context, SetProductsParams) error

		// GetProducts returns all the products of a shop from a specific external provider
		GetProducts(context.Context, GetProductsParams) ([]Product, error)

		// GetCursor returns a Cursor by name
		GetCursor(context.Context, GetCursorParams) (*Cursor, error)

		// SetCursor sets a Cursor by name
		SetCursor(context.Context, Cursor) error

		// ResetCursor resets a Cursor by name
		ResetCursor(context.Context, ResetCursorParams) error

		// GetNextShop returns the next shop to sync
		GetNextShop(context.Context, Cursor) (NextShop, error)
	}
)
