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
		ctx                  context.Context
		orderProcessingTimer int
	}
)

var (
	ErrorUserNotFound   = errors.New("user not found")
	ErrorWebAppNotFound = errors.New("web app id not found")
	ErrorInternal       = errors.New("internal server error")
)

// New creates a new user service
func New(repo Repository, log *zap.Logger, ctx context.Context, orderProcessingTimer int) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if orderProcessingTimer == 0 {
		orderProcessingTimer = 300
	}

	return &Service{
		repo:                 repo,
		log:                  log,
		ctx:                  ctx,
		orderProcessingTimer: orderProcessingTimer,
	}
}

func (s *Service) getOrderNotifications(cur Cursor) ([]OrderNotification, error) {
	return nil, nil
}

func (s *Service) sendOrderNotifications(orderNotifications []OrderNotification) (Cursor, error) {
	return Cursor{}, nil
}

func (s *Service) notifyIteration() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, "defaultCursor")

	orderNotifications, err := s.getOrderNotifications(cursor)
	if err != nil {
		return errors.Wrap(err, "s.getOrderNotifications")
	}

	if len(orderNotifications) == 0 {
		// FIXME: Log warning
		return nil
	}

	updCursor, err := s.sendOrderNotifications(orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendOrderNotifications")
	}

	err = s.repo.UpdateNotifierCursor(s.ctx, updCursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

func (s *Service) Run() error {
	ticker := time.NewTicker(time.Duration(s.orderProcessingTimer) * time.Second)

	for {
		select {
		case <-ticker.C:
			err := s.notifyIteration()
			if err != nil {
				return errors.Wrap(err, "s.notifyIteration")
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

func (s *Service) Shutdown(cancel context.CancelFunc) error {
	cancel()
	return nil
}
