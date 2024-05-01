package adminbot

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// Cursor defines the structure for a notify list cursor
	Cursor struct {
		Name                   string
		LastProcessedCreatedAt time.Time
		LastProcessedID        uuid.UUID
	}

	// OrderNotification defines the structure of order notification
	OrderNotification struct {
		ReadableID     int64
		WebAppID       uuid.UUID
		ExternalUserID int
	}

	// Repository provides access to the user storage
	Repository interface {
		GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error)
		GetAdminBotToken(ctx context.Context, webAppID uuid.UUID) (string, error)
		GetNotifierCursor(ctx context.Context, name string) (Cursor, error)
		UpdateNotifierCursor(ctx context.Context, cur Cursor) error
		GetNotificationsForOrdersAfterCursor(ctx context.Context, cur Cursor) ([]OrderNotification, error)
	}

	// Service provides user operations
	Service struct {
		repo                 Repository
		log                  *zap.Logger
		orderProcessingTimer int
	}
)

var (
	ErrorUserNotFound   = errors.New("user not found")
	ErrorWebAppNotFound = errors.New("web app id not found")
	ErrorInternal       = errors.New("internal server error")
)

// New creates a new user service
func New(repo Repository, log *zap.Logger, orderProcessingTimer int) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}

	return &Service{
		repo:                 repo,
		log:                  log,
		orderProcessingTimer: orderProcessingTimer,
	}
}

func (s *Service) getOrderNotifications(ctx context.Context) ([]OrderNotification, error) {
	return nil, nil
}

func (s *Service) sendOrderNotifications(ctx context.Context, orderNotifications []OrderNotification) error {
	return nil
}

func (s *Service) notifyIteration(ctx context.Context) error {
	// TODO: Fetch cursor here

	orderNotifications, err := s.getOrderNotifications(ctx)
	if err != nil {
		return errors.Wrap(err, "s.getOrderNotifications")
	}

	err = s.sendOrderNotifications(ctx, orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendOrderNotifications")
	}

	// TODO: Update cursor here

	return nil
}

func (s *Service) Run(ctx context.Context, done <-chan interface{}) error {
	ticker := time.NewTicker(time.Duration(s.orderProcessingTimer) * time.Second)

	for {
		select {
		case <-ticker.C:
			err := s.notifyIteration(ctx)
			if err != nil {
				return errors.Wrap(err, "s.notifyIteration")
			}
		case <-done:
			ticker.Stop()
			return nil
		}
	}
}

func (s *Service) Shutdown(ctx context.Context, done chan<- interface{}) error {
	done <- interface{}(nil)
	return nil
}
