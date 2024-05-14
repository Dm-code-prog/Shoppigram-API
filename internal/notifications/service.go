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
	"github.com/shoppigram-com/marketplace-api/internal/logging"
	"go.uber.org/zap"
)

//go:embed message.md
var messageTemplate embed.FS

type (
	// Cursor defines the structure for a notify list cursor
	Cursor struct {
		Name            string
		CursorDate      time.Time
		LastProcessedID uuid.UUID
	}

	// Product is a marketplace product
	Product struct {
		Name          string
		Quantity      int
		Price         float64
		PriceCurrency string
	}

	// NewOrderNotification defines the structure of order notification
	NewOrderNotification struct {
		ID              uuid.UUID
		ReadableOrderID int64
		CreatedAt       time.Time
		UserNickname    string
		WebAppID        uuid.UUID
		WebAppName      string
		Products        []Product
	}

	// NewMarketplaceNotification defines the structure of new marketplace notification
	NewMarketplaceNotification struct {
		ID        uuid.UUID
		Name      string
		CreatedAt time.Time
	}

	// Repository provides access to the user storage
	Repository interface {
		GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error)
		GetReviewersNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error)
		GetNotifierCursor(ctx context.Context, name string) (Cursor, error)
		UpdateNotifierCursor(ctx context.Context, cur Cursor) error
		GetNotificationsForNewOrdersAfterCursor(ctx context.Context, cur Cursor) ([]NewOrderNotification, error)
		GetNotificationsForNewMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]NewMarketplaceNotification, error)
	}

	// Service provides user operations
	Service struct {
		repo                          Repository
		log                           *zap.Logger
		cache                         *ristretto.Cache
		ctx                           context.Context
		cancel                        context.CancelFunc
		orderProcessingTimer          time.Duration
		newMarketplaceProcessingTimer time.Duration
		botToken                      string
	}
)

const (
	orderNotifierName          = "order_notifications"
	newMarketplaceNotifierName = "new_marketplace_notifications"
)

// BuildMessage creates a notification message for a new order
func (o *NewOrderNotification) BuildMessage() (string, error) {
	var subtotal float64
	var productList strings.Builder
	var currency string
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		currency = p.PriceCurrency
		productList.WriteString(fmt.Sprintf(`\- %dx %s по цене %s %s
`, p.Quantity, escapeSpecialSymbols(p.Name), formatFloat(p.Price), formatCurrency(p.PriceCurrency)))
	}

	data, err := messageTemplate.ReadFile("newOrderMessage.md")
	if err != nil {
		return "", errors.Wrap(err, "messageTemplate.ReadFile")
	}

	return fmt.Sprintf(string(data),
		escapeSpecialSymbols(o.WebAppName),
		escapeSpecialSymbols(o.UserNickname),
		o.ReadableOrderID,
		formatRussianTime(o.CreatedAt),
		formatFloat(subtotal)+" "+formatCurrency(currency),
		strings.TrimRight(productList.String(), "; "),
	), nil
}

// BuildMessage creates a notification message for a new order
func (m *NewMarketplaceNotification) BuildMessage() (string, error) {
	data, err := messageTemplate.ReadFile("newMarketplaceMessage.md")
	if err != nil {
		return "", errors.Wrap(err, "messageTemplate.ReadFile")
	}

	return fmt.Sprintf(string(data),
		escapeSpecialSymbols(m.ID.String()),
		escapeSpecialSymbols(m.Name),
	), nil
}

// New creates a new user service
func New(repo Repository, log *zap.Logger, orderProcessingTimer time.Duration, newMarketplaceProcessingTimer time.Duration, botToken string) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if orderProcessingTimer == 0 {
		log.Fatal("order processing timer is not specified")
		return nil
	}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e2,     // number of keys to track frequency of (100).
		MaxCost:     100_000, // maximum cost of cache (100KB).
		BufferItems: 10,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal("cache *ristretto.Cache is nil, fatal")
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Service{
		repo:                          repo,
		log:                           log,
		ctx:                           ctx,
		cancel:                        cancel,
		cache:                         cache,
		orderProcessingTimer:          orderProcessingTimer,
		newMarketplaceProcessingTimer: orderProcessingTimer,
		botToken:                      botToken,
	}
}

// RunOrderNotifier starts a job that batch loads new orders
// and sends notifications to the owners of marketplaces
func (s *Service) RunOrderNotifier() error {
	ticker := time.NewTicker(s.orderProcessingTimer)

	for {
		select {
		case <-ticker.C:
			err := s.runOrderNotifierOnce()
			if err != nil {
				s.log.Error("runOrderNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

func (s *Service) runOrderNotifierOnce() error {
	defer s.cache.Clear()
	cursor, err := s.repo.GetNotifierCursor(s.ctx, orderNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	orderNotifications, err := s.repo.GetNotificationsForNewOrdersAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForNewOrdersAfterCursor")
	}

	if len(orderNotifications) == 0 {
		return nil
	}

	s.log.With(zap.String("count", strconv.Itoa(len(orderNotifications)))).Info("sending notifications for new orders")
	err = s.sendOrderNotifications(orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendOrderNotifications")
	}

	lastElem := orderNotifications[len(orderNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.CreatedAt,
		LastProcessedID: lastElem.ID,
		Name:            orderNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// RunNewMarketplaceNotifier starts a job that batch loads new marketplaces
// and sends notifications to the reviewers of marketplaces
func (s *Service) RunNewMarketplaceNotifier() error {
	ticker := time.NewTicker(s.newMarketplaceProcessingTimer)

	for {
		select {
		case <-ticker.C:
			err := s.runNewMarketplaceNotifierOnce()
			if err != nil {
				s.log.Error("runNewMarketplaceNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

func (s *Service) runNewMarketplaceNotifierOnce() error {
	return nil
}

// Shutdown stops the job
func (s *Service) Shutdown() error {
	s.cancel()
	<-s.ctx.Done()
	return nil
}

func (s *Service) sendOrderNotifications(orderNotifications []NewOrderNotification) error {
	var (
		bot *tgbotapi.BotAPI
		err error
	)

	for _, a := range orderNotifications {
		val, ok := s.cache.Get(a.WebAppID.String())
		if ok {
			bot = val.(*tgbotapi.BotAPI)
		} else {
			bot, err = tgbotapi.NewBotAPI(s.botToken)
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

func formatFloat(num float64) string {
	str := strconv.FormatFloat(num, 'f', -1, 64)
	parts := strings.Split(str, ".")
	intPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	n := len(intPart)
	if n <= 3 {
		return intPart + decimalPart
	}

	var result string
	for i := 0; i < n; i++ {
		result = string(intPart[n-1-i]) + result
		if (i+1)%3 == 0 && i != n-1 {
			result = "," + result
		}
	}
	return result + decimalPart
}

func formatCurrency(currency string) string {
	currency = strings.ToLower(currency)
	switch currency {
	case "usd":
		return "$"
	case "eur":
		return "€"
	case "rub":
		return "₽"
	default:
		return currency
	}
}

func formatRussianTime(t time.Time) string {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return ""
	}
	t = t.In(loc)
	return strings.ReplaceAll(t.Format("02.01.2006 15:04:05"), ".", "\\.")
}

var specialSymbols = []string{"*", "_", "#", "-"}

func escapeSpecialSymbols(s string) string {
	for _, sym := range specialSymbols {
		if strings.Contains(s, sym) {
			return strings.ReplaceAll(s, sym, "\\"+sym)
		}
	}

	return s
}
