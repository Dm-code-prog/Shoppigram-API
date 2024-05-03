package notifications

import (
	"context"
	"embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/ristretto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//go:embed message.md
var messageTemplate embed.FS

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
	notifierName = "order_notifications"
)

// BuildMessage creates a notification message for a new order
func (o *OrderNotification) BuildMessage() (string, error) {
	var productList strings.Builder
	for _, p := range o.Products {
		productList.WriteString(fmt.Sprintf(`\- %d x %s по цене %d
`, p.Quantity, p.Name, int(p.Price)))
	}

	data, err := messageTemplate.ReadFile("message.md")
	if err != nil {
		return "", errors.Wrap(err, "messageTemplate.ReadFile")
	}

	return fmt.Sprintf(string(data),
		o.ReadableOrderID,
		o.UserNickname,
		strings.TrimRight(productList.String(), "; "),
	), nil
}

// New creates a new user service
func New(repo Repository, log *zap.Logger, orderProcessingTimer time.Duration) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if orderProcessingTimer == 0 {
		log.Fatal("order processing timer is not specified")
	}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e2,     // number of keys to track frequency of (100).
		MaxCost:     100_000, // maximum cost of cache (100KB).
		BufferItems: 10,      // number of keys per Get buffer.
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

// Run starts a job that batch loads new orders
// and sends notifications to the owners of marketplaces
func (s *Service) Run() error {
	ticker := time.NewTicker(s.orderProcessingTimer)

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

func (s *Service) runOnce() error {
	defer s.cache.Clear()
	cursor, err := s.repo.GetNotifierCursor(s.ctx, notifierName)

	orderNotifications, err := s.repo.GetNotificationsForOrdersAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.getOrderNotifications")
	}

	if len(orderNotifications) == 0 {
		s.log.Info("no new orders found, skipping")
		return nil
	}

	s.log.With(zap.String("count", strconv.Itoa(len(orderNotifications)))).Info("sending notifications for new orders")
	err = s.sendOrderNotifications(orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendOrderNotifications")
	}

	lastElem := orderNotifications[len(orderNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		LastProcessedCreatedAt: lastElem.CreatedAt,
		LastProcessedID:        lastElem.ID,
		Name:                   notifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// Shutdown stops the job
func (s *Service) Shutdown() error {
	s.cancel()
	<-s.ctx.Done()
	return nil
}

func (s *Service) sendOrderNotifications(orderNotifications []OrderNotification) error {
	var bot *tgbotapi.BotAPI

	for _, a := range orderNotifications {
		val, ok := s.cache.Get(a.WebAppID.String())
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

			s.cache.SetWithTTL(a.WebAppID.String(), bot, 0, 10*time.Minute)
		}

		nl, err := s.repo.GetAdminsNotificationList(s.ctx, a.WebAppID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetAdminsNotificationList")
		}

		// need to get chat id's of users, who we are going to send messages
		for _, v := range nl {
			fmt.Println(a.BuildMessage())
			msgTxt, err := a.BuildMessage()
			if err != nil {
				return errors.Wrap(err, "a.BuildMessage")
			}
			msg := tgbotapi.NewMessage(v, msgTxt)
			msg.ParseMode = tgbotapi.ModeMarkdownV2
			_, err = bot.Send(msg)
			if err != nil {
				return errors.Wrap(err, "bot.Send")
			}
		}

	}

	return nil
}
