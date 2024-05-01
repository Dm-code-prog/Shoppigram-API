package adminbot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dgraph-io/ristretto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/products"
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
		//Products       []orders.Product
		//BuyerUsername  string
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
		cache                *ristretto.Cache
		ctx                  context.Context
		orderProcessingTimer int
	}
)

const (
	defaultCursor = "orders_cursor"
	template      = "Hello! User %d want to buy products %d with a total price of %d. Order ID: `%d` ."
)

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorWebAppNotFound    = errors.New("web app id not found")
	ErrorInternal          = errors.New("internal server error")
	ErrorEmptyProductsList = errors.New("empty products list")
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

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e2,       // number of keys to track frequency of (100).
		MaxCost:     2_000_000, // maximum cost of cache (200 MB).
		BufferItems: 10,        // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal("cache *ristretto.Cache is nil, fatal")
	}

	return &Service{
		repo:                 repo,
		log:                  log,
		ctx:                  ctx,
		cache:                cache,
		orderProcessingTimer: orderProcessingTimer,
	}
}

func (s *Service) getOrderNotifications(cur Cursor) ([]OrderNotification, error) {
	adminsList, err := s.repo.GetAdminsNotificationList(s.ctx, webAppID)
	if err != nil {
		return nil, errors.Wrap(err, "s.repo.GetAdminsNotificationList")
	}

	return nil, nil
}

func (s *Service) buildMessage(orderId int, username string, products []products.Product, totalPrice float32) (string, error) {
	if len(products) == 0 {
		return "", ErrorEmptyProductsList
	}

	return fmt.Sprintf(
		template,
		username,
		strings.Join(products, ", "),
		totalPrice,
		orderId,
	), nil
}

func (s *Service) sendOrderNotifications(orderNotifications []OrderNotification) (Cursor, error) {
	var bot *tgbotapi.BotAPI

	for _, a := range orderNotifications {
		if a.WebAppID == uuid.Nil {
			// FIXME: It should not be like that, log
			continue
		}

		val, ok := s.cache.Get(a.WebAppID)
		if ok {
			bot = val.(*tgbotapi.BotAPI)
		} else {
			s.log.With(
				zap.String("web_app_id", a.WebAppID.String()),
			).Info("cache miss")

			token, err := s.repo.GetAdminBotToken(s.ctx, a.WebAppID)
			if err != nil {
				return Cursor{}, errors.Wrap(err, "s.repo.GetAdminBotToken")
			}

			bot, err = tgbotapi.NewBotAPI(token)
			if err != nil {
				return Cursor{}, errors.Wrap(err, "tgbotapi.NewBotAPI")
			}

			// Cache the bot structure
			ok = s.cache.SetWithTTL(a.WebAppID, bot, 0, 10*time.Minute)
			if !ok {
				// FIXME: Probably something is wrong, log
			}
		}

		msg := tgbotapi.NewMessage(a, "Test message")
		_, err := bot.Send(msg)
		if err != nil {
			return Cursor{}, errors.Wrap(err, "bot.Send")
		}
	}

	// FIXME: Return updated cursor after sending the messages
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
