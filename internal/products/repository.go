package products

import (
	"context"
	"database/sql"
	"errors"
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
		if errors.Is(err, sql.ErrNoRows) {
			return GetProductsResponse{}, ErrorNotFound
		}
	}

	var products []Product
	for _, p := range prod {
		products = append(products, Product{
			ID:            p.ID,
			Name:          p.Name,
			Description:   p.Description.String,
			Price:         p.Price,
			PriceCurrency: p.PriceCurrency,
			ImageURL:      p.ImageUrl,
		})
	}

	return GetProductsResponse{
		WebAppName: prod[0].WebAppName,
		Products:   products,
	}, nil
}
