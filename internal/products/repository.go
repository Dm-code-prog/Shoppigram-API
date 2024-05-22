package products

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/products/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen *generated.Queries
}

// NewPg creates a new Pg
func NewPg(gen *generated.Queries) *Pg {
	return &Pg{gen: gen}
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
