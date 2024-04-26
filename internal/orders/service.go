package orders

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/products"
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
		WebAppID     uuid.UUID
		UserID       uuid.UUID
		UserNickname string
		Products     []Product `json:"products"`
	}

	// CreateOrderResponse returns the ID of the newly created order
	CreateOrderResponse struct {
		ID uuid.UUID `json:"id"`
	}

	// Repository is the storage interface for orders
	Repository interface {
		CreateOrder(context.Context, CreateOrderRequest) (CreateOrderResponse, error)
	}

	NotificationProduct struct {
		ID       uuid.UUID
		Quantity int32
		Name     string
	}

	// OrderNotification makes up the information
	// about an order that should be delivered
	// to the seller of the products
	OrderNotification struct {
		ID           uuid.UUID
		UserNickname string
		Products     []NotificationProduct
	}

	// Notifier sends message notifications
	// on events
	Notifier interface {
		NotifySeller(context.Context, OrderNotification) error
	}
)

type (
	// Service is the order service
	Service struct {
		logger     *zap.Logger
		repo       Repository
		notifier   Notifier
		productSvc *products.Service
	}
)

var (
	// ErrorBadRequest is returned to the end user when the request is malformed
	ErrorBadRequest = errors.New("the request to create an order is malformed")
)

// New returns new instance of Service
func New(logger *zap.Logger, repo Repository, n Notifier) *Service {
	return &Service{
		logger:   logger,
		repo:     repo,
		notifier: n,
	}
}

// CreateOrder saves an order to the database
// and notifies the clients, that own the marketplace web app
// about a new order
func (s *Service) CreateOrder(ctx context.Context, req CreateOrderRequest) (CreateOrderResponse, error) {
	res, err := s.repo.CreateOrder(ctx, req)
	if err != nil {
		return CreateOrderResponse{}, errors.Wrap(err, "s.repo.CreateOrder()")
	}

	go s.notifySeller(ctx, req, res)
	return res, nil
}

func (s *Service) notifySeller(ctx context.Context, req CreateOrderRequest, resp CreateOrderResponse) {
	webAppProducts, err := s.productSvc.GetProducts(ctx, products.GetProductsRequest{WebAppID: req.WebAppID})
	if err != nil {
		s.logger.Error("s.productSvc.GetProducts()", zap.Error(err))
		return
	}

	var webAppProductsMap map[string]products.Product
	for _, v := range webAppProducts.Products {
		webAppProductsMap[v.ID.String()] = v
	}

	// enrich the product with the name
	// for the notification message
	var orderProducts []NotificationProduct
	for _, product := range req.Products {
		if WaP, ok := webAppProductsMap[product.ID.String()]; ok {
			orderProducts = append(orderProducts, NotificationProduct{
				ID:       product.ID,
				Quantity: product.Quantity,
				Name:     WaP.Name,
			})
		}
	}

	err = s.notifier.NotifySeller(ctx, OrderNotification{
		ID:           resp.ID,
		UserNickname: req.UserNickname,
		Products:     orderProducts,
	})
	if err != nil {
		s.logger.Error("s.notifier.NotifySeller()", zap.Error(err))
	}
}
