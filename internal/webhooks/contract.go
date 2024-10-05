package webhooks

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Repository provides access to the webhooks storage
type Repository interface {
	GetOrder(ctx context.Context, id uuid.UUID) (Order, error)
	SetOrderStateConfirmed(ctx context.Context, id uuid.UUID) error
	SavePaymentExtraInfo(ctx context.Context, params SavePaymentExtraInfoParams) error
	CreateOrUpdateTelegramChannel(ctx context.Context, req CreateOrUpdateTelegramChannelRequest) error
}

var (
	ErrorInternalServerError = errors.New("internal server error")
	ErrorOrderDoesntExist    = errors.New("order does not exist")
	ErrorBadRequest          = errors.New("bad request")
)
