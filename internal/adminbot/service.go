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
		repo Repository
		log  *zap.Logger
	}
)

var (
	ErrorUserNotFound   = errors.New("user not found")
	ErrorWebAppNotFound = errors.New("web app id not found")
	ErrorInternal       = errors.New("internal server error")
)

// New creates a new user service
func New(repo Repository, log *zap.Logger) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}

	return &Service{
		repo: repo,
		log:  log,
	}
}

func (s *Service) notifyIteration(ctx context.Context) error {
	return nil
}

func (s *Service) Run(ctx context.Context, done <-chan interface{}) error {
	ticker := time.NewTicker(300 * time.Second)

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
