package orders

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"github.com/shoppigram-com/marketplace-api/internal/products"
	telegramusers "github.com/shoppigram-com/marketplace-api/internal/users"
	"go.uber.org/zap"
)

type (
	// Product is a marketplace product
	// that is identified by the ID and quantity
	Product struct {
		ID       uuid.UUID `json:"id"`
		Quantity int32     `json:"quantity"`
	}

	// CreateOrderRequest specifies the products
	// of a web app marketplace that make up
	// the order and user information
	CreateOrderRequest struct {
		WebAppID uuid.UUID
		Products []Product `json:"products"`
	}

	// CreateOrderResponse returns the ID of the newly created order
	CreateOrderResponse struct {
		ReadableID int `json:"readable_id"`
	}

	// SaveOrderRequest is a request to save order info
	// to the storage
	SaveOrderRequest struct {
		WebAppID       uuid.UUID
		Products       []Product
		ExternalUserID int
	}

	// SaveOrderResponse is the response to SaveOrderRequest
	//
	// It contains the readable order ID
	SaveOrderResponse struct {
		ReadableID int
	}

	// Repository is the storage interface for orders
	Repository interface {
		CreateOrder(context.Context, SaveOrderRequest) (SaveOrderResponse, error)
	}
)

type (
	// Service is the order service
	Service struct {
		log        *zap.Logger
		repo       Repository
		productSvc *products.Service
	}
)

var (
	// ErrorBadRequest is returned to the end user when the request is malformed
	ErrorBadRequest = errors.New("the request to create an order is malformed")

	ErrorInvalidProductQuantity = errors.New("the product quantity must be greater than zero")

	// ErrorInternal is a server error
	ErrorInternal = errors.New("internal error, try again later")
)

// New returns new instance of Service
func New(repo Repository, log *zap.Logger) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

// CreateOrder saves an order to the database
// and notifies the clients, that own the marketplace web app
// about a new order
func (s *Service) CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	u, err := telegramusers.GetUserFromContext(ctx)
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "telegramusers.GetUserFromContext")
	}

	res, err := s.repo.CreateOrder(ctx, SaveOrderRequest{
		WebAppID:       req.WebAppID,
		Products:       req.Products,
		ExternalUserID: int(u.ExternalId),
	})
	if err != nil {
		s.log.
			With(zap.String("web_app_id", req.WebAppID.String())).
			Error("repository.CreateOrder()", logging.SilentError(err))
		return CreateOrderResponse{}, errors.Wrap(err, "s.repo.CreateOrder")
	}
	return CreateOrderResponse(res), nil
}
