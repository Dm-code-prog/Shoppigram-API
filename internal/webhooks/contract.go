package webhooks

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Repository provides access to the webhooks storage
type (
	Repository interface {
		GetOrder(ctx context.Context, id uuid.UUID) (Order, error)
		SetOrderStateConfirmed(ctx context.Context, id uuid.UUID) error
		SavePaymentExtraInfo(ctx context.Context, params SavePaymentExtraInfoParams) error
		CreateOrUpdateTelegramChannel(ctx context.Context, req CreateOrUpdateTelegramChannelRequest) error
		DeleteTelegramChannel(ctx context.Context, req DeleteTelegramChannelRequest) error
	}

	// Notifier is the service for notifications
	// The interface requires a method for notifying a user about a successful
	// channel integration with Shoppigram
	Notifier interface {
		NotifyChannelIntegrationSuccess(ctx context.Context, request NotifyChannelIntegrationSuccessRequest) error
		NotifyChannelIntegrationFailure(ctx context.Context, request NotifyChannelIntegrationFailureRequest) error
		NotifyGreetings(ctx context.Context, request NotifyGreetingsRequest) error
	}
)

var (
	ErrorInternalServerError = errors.New("internal server error")
	ErrorOrderDoesntExist    = errors.New("order does not exist")
	ErrorBadRequest          = errors.New("bad request")
)
