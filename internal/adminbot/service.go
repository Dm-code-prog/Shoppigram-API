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
	"go.uber.org/zap"
)

type (
	// Cursor defines the structure for a notify list cursor
	Cursor struct {
		Name                   string
		LastProcessedCreatedAt time.Time
		LastProcessedID        uuid.UUID
	}

	// Product is a marketplace product
	Product struct {
		Name     string
		Quantity int
		Price    float64
	}

	// OrderNotification defines the structure of order notification
	OrderNotification struct {
		ID              uuid.UUID
		ReadableOrderID int64
		CreatedAt       time.Time
		UserNickname    string
		WebAppID        uuid.UUID
		Products        []Product
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
		cancel               context.CancelFunc
		orderProcessingTimer time.Duration
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

// String creates a notification message for a new order
func (o *OrderNotification) String() string {
	var productList strings.Builder
	for _, p := range o.Products {
		productList.WriteString(fmt.Sprintf("%d x %s at $%.2f each; ", p.Quantity, p.Name, p.Price))
	}

	return fmt.Sprintf("У вас новый заказ. Номер заказа: #%d, создан: %s, пользователь: %s, Продукты в заказе: [%s]",
		o.ReadableOrderID,
		o.CreatedAt.Format("Jan 2, 2006 at 3:04pm (MST)"),
		o.UserNickname,
		strings.TrimRight(productList.String(), "; "),
	)
}

// New creates a new user service
func New(repo Repository, log *zap.Logger, orderProcessingTimer time.Duration) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if orderProcessingTimer == 0 {
		orderProcessingTimer = 300
	}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e2,       // number of keys to track frequency of (100).
		MaxCost:     2_000_000, // maximum cost of cache (2MB).
		BufferItems: 10,        // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal("cache *ristretto.Cache is nil, fatal")
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Service{
		repo:                 repo,
		log:                  log,
		ctx:                  ctx,
		cancel:               cancel,
		cache:                cache,
		orderProcessingTimer: orderProcessingTimer,
	}
}

func (s *Service) Run() error {
	ticker := time.NewTicker(s.orderProcessingTimer * time.Minute)

	for {
		select {
		case <-ticker.C:
			err := s.runOnce()
			if err != nil {
				return errors.Wrap(err, "s.runOnce")
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

func (s *Service) Shutdown() error {
	s.cancel()
	return nil
}

func (s *Service) sendOrderNotifications(orderNotifications []OrderNotification) error {
	var bot *tgbotapi.BotAPI

	for _, a := range orderNotifications {
		val, ok := s.cache.Get(a.WebAppID)
		if ok {
			bot = val.(*tgbotapi.BotAPI)
		} else {
			token, err := s.repo.GetAdminBotToken(s.ctx, a.WebAppID)
			if err != nil {
				return errors.Wrap(err, "s.repo.GetAdminBotToken")
			}

			bot, err = tgbotapi.NewBotAPI(token)
			if err != nil {
				return errors.Wrap(err, "tgbotapi.NewBotAPI")
			}

			s.cache.SetWithTTL(a.WebAppID, bot, 0, 10*time.Minute)
		}

		nl, err := s.repo.GetAdminsNotificationList(s.ctx, a.WebAppID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetAdminsNotificationList")
		}

		// need to get chat id's of users, who we are going to send messages
		for _, v := range nl {
			msg := tgbotapi.NewMessage(v, a.String())
			_, err := bot.Send(msg)
			if err != nil {
				return errors.Wrap(err, "bot.Send")
			}
		}

	}

	return nil
}

func (s *Service) runOnce() error {
	cursor, err := s.repo.GetNotifierCursor(s.ctx, "defaultCursor")

	orderNotifications, err := s.repo.GetNotificationsForOrdersAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.getOrderNotifications")
	}

	err = s.sendOrderNotifications(orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendOrderNotifications")
	}

	lastElem := orderNotifications[len(orderNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		LastProcessedCreatedAt: lastElem.CreatedAt,
		LastProcessedID:        lastElem.ID,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}
