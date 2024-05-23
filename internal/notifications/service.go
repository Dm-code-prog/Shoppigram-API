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

//go:embed templates/*.md
var templates embed.FS

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
		ID            uuid.UUID
		Name          string
		ShortName     string
		CreatedAt     time.Time
		OwnerUsername string
	}

	// VerifiedMarketplaceNotification defines the structure of verified marketplace notification
	VerifiedMarketplaceNotification struct {
		ID                  uuid.UUID
		Name                string
		ShortName           string
		VerifiedAt          time.Time
		OwnerExternalUserID int64
	}

	// AddUserToNewOrderNotificationsRequest creates a new order notification
	// list entry for some marketplace
	AddUserToNewOrderNotificationsRequest struct {
		WebAppID    uuid.UUID
		AdminChatID int64
	}

	// Repository provides access to the user storage
	Repository interface {
		GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error)
		GetReviewersNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error)
		GetNotifierCursor(ctx context.Context, name string) (Cursor, error)
		UpdateNotifierCursor(ctx context.Context, cur Cursor) error
		GetNotificationsForNewOrdersAfterCursor(ctx context.Context, cur Cursor) ([]NewOrderNotification, error)
		GetNotificationsForNewMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]NewMarketplaceNotification, error)
		GetNotificationsForVerifiedMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]VerifiedMarketplaceNotification, error)
		AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsRequest) error
	}

	// Service provides user operations
	Service struct {
		repo                               Repository
		log                                *zap.Logger
		cache                              *ristretto.Cache
		ctx                                context.Context
		cancel                             context.CancelFunc
		newOrderProcessingTimer            time.Duration
		newMarketplaceProcessingTimer      time.Duration
		verifiedMarketplaceProcessingTimer time.Duration
		botToken                           string
	}
)

const (
	newOrderNotifierName            = "new_order_notifications"
	newMarketplaceNotifierName      = "new_marketplace_notifications"
	verifiedMarketplaceNotifierName = "verified_marketplace_notifications"
	marketplaceURL                  = "https://web-app.shoppigram.com/app/"
	webAppURL                       = "https://t.me/shoppigrambot/"
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
`, p.Quantity, escapeSpecialSymbols(p.Name), escapeSpecialSymbols(formatFloat(p.Price)), formatCurrency(p.PriceCurrency)))
	}

	newOrderMessageTemplate, err := templates.ReadFile("templates/new_order_message.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	return fmt.Sprintf(
		escapeSpecialSymbols(string((newOrderMessageTemplate))),
		escapeSpecialSymbols(o.WebAppName),
		escapeSpecialSymbols(o.UserNickname),
		o.ReadableOrderID,
		formatRussianTime(o.CreatedAt),
		escapeSpecialSymbols(formatFloat(subtotal))+" "+formatCurrency(currency),
		strings.TrimRight(productList.String(), "; "),
	), nil
}

// BuildMessage creates a notification message for a new marketplace
func (m *NewMarketplaceNotification) BuildMessage() (string, error) {
	newMarketplaceMessageTemplate, err := templates.ReadFile("templates/new_marketplace_message.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	return fmt.Sprintf(
		escapeSpecialSymbols(string(newMarketplaceMessageTemplate)),
		escapeSpecialSymbols(m.OwnerUsername),
		escapeSpecialSymbols(m.Name),
		escapeSpecialSymbols(m.ShortName),
		escapeSpecialSymbols(marketplaceURL+m.ID.String()),
	), nil
}

// BuildMessage creates a notification message for a verified marketplace
func (m *VerifiedMarketplaceNotification) BuildMessage() (string, error) {
	verifiedMarketplaceMessageTemplate, err := templates.ReadFile("templates/verified_marketplace_message.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	return fmt.Sprintf(
		escapeSpecialSymbols(string(verifiedMarketplaceMessageTemplate)),
		escapeSpecialSymbols(m.Name),
		escapeSpecialSymbols(webAppURL+m.ShortName),
	), nil
}

// New creates a new user service
func New(repo Repository, log *zap.Logger, newOrderProcessingTimer time.Duration, newMarketplaceProcessingTimer time.Duration, verifiedMarketplaceProcessingTimer time.Duration, botToken string) *Service {
	if log == nil {
		log, _ = zap.NewProduction()
		log.Warn("log *zap.Logger is nil, using zap.NewProduction")
	}
	if newOrderProcessingTimer == 0 {
		log.Fatal("new order processing timer is not specified")
		return nil
	}
	if newMarketplaceProcessingTimer == 0 {
		log.Fatal("new marketplace processing timer is not specified")
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
		repo:                               repo,
		log:                                log,
		ctx:                                ctx,
		cancel:                             cancel,
		cache:                              cache,
		newOrderProcessingTimer:            newOrderProcessingTimer,
		newMarketplaceProcessingTimer:      newMarketplaceProcessingTimer,
		verifiedMarketplaceProcessingTimer: verifiedMarketplaceProcessingTimer,
		botToken:                           botToken,
	}
}

// RunNewOrderNotifier starts a job that batch loads new orders
// and sends notifications to the owners of marketplaces
func (s *Service) RunNewOrderNotifier() error {
	ticker := time.NewTicker(s.newOrderProcessingTimer)

	for {
		select {
		case <-ticker.C:
			err := s.runNewOrderNotifierOnce()
			if err != nil {
				s.log.Error("runNewOrderNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runNewOrderNotifierOnce executes one iteration of loading a batch of new
// orders and sending notifications to the owners of marketplaces
func (s *Service) runNewOrderNotifierOnce() error {
	defer s.cache.Clear()
	cursor, err := s.repo.GetNotifierCursor(s.ctx, newOrderNotifierName)
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

	s.log.With(
		zap.String("count", strconv.Itoa(len(orderNotifications))),
	).Info("sending notifications for new orders")
	err = s.sendNewOrderNotifications(orderNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendNewOrderNotifications")
	}

	lastElem := orderNotifications[len(orderNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.CreatedAt,
		LastProcessedID: lastElem.ID,
		Name:            newOrderNotifierName,
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

// runNewMarketplaceNotifierOnce executes one iteration of loading a batch of new
// marketplaces and sending notifications to the reviewers of marketplaces
func (s *Service) runNewMarketplaceNotifierOnce() error {
	defer s.cache.Clear()
	cursor, err := s.repo.GetNotifierCursor(s.ctx, newMarketplaceNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	marketplaceNotifications, err := s.repo.GetNotificationsForNewMarketplacesAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForNewMarketplacesAfterCursor")
	}

	if len(marketplaceNotifications) == 0 {
		return nil
	}

	s.log.With(
		zap.String("count", strconv.Itoa(len(marketplaceNotifications))),
	).Info("sending notifications for new marketplaces")
	err = s.sendNewMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendNewMarketplaceNotifications")
	}

	lastElem := marketplaceNotifications[len(marketplaceNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.CreatedAt,
		LastProcessedID: lastElem.ID,
		Name:            newMarketplaceNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// RunVerifiedMarketplaceNotifier starts a job that batch loads verified marketplaces
// and sends notifications to the owners of those marketplaces
func (s *Service) RunVerifiedMarketplaceNotifier() error {
	ticker := time.NewTicker(s.verifiedMarketplaceProcessingTimer)

	for {
		select {
		case <-ticker.C:
			err := s.runVerifiedMarketplaceNotifierOnce()
			if err != nil {
				s.log.Error("runVerifiedMarketplaceNotifierOnce failed", logging.SilentError(err))
				continue
			}
		case <-s.ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}

// runVerifiedMarketplaceNotifierOnce executes one iteration of loading a batch of
// verified marketplaces and sending notifications to the owners of those marketplaces
func (s *Service) runVerifiedMarketplaceNotifierOnce() error {
	defer s.cache.Clear()
	cursor, err := s.repo.GetNotifierCursor(s.ctx, verifiedMarketplaceNotifierName)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotifierCursor")
	}

	marketplaceNotifications, err := s.repo.GetNotificationsForVerifiedMarketplacesAfterCursor(s.ctx, cursor)
	if err != nil {
		return errors.Wrap(err, "s.repo.GetNotificationsForVerifiedMarketplacesAfterCursor")
	}

	if len(marketplaceNotifications) == 0 {
		return nil
	}

	s.log.With(
		zap.String("count", strconv.Itoa(len(marketplaceNotifications))),
	).Info("sending notifications for verified marketplaces")
	err = s.sendVerifiedMarketplaceNotifications(marketplaceNotifications)
	if err != nil {
		return errors.Wrap(err, "s.sendVerifiedMarketplaceNotifications")
	}

	lastElem := marketplaceNotifications[len(marketplaceNotifications)-1]

	err = s.repo.UpdateNotifierCursor(s.ctx, Cursor{
		CursorDate:      lastElem.VerifiedAt,
		LastProcessedID: lastElem.ID,
		Name:            verifiedMarketplaceNotifierName,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.UpdateNotifierCursor")
	}

	return nil
}

// Shutdown stops all of the notifications
func (s *Service) Shutdown() error {
	s.cancel()
	<-s.ctx.Done()
	return nil
}

// sendNewOrderNotifications sends batch of notifications for new orders
func (s *Service) sendNewOrderNotifications(orderNotifications []NewOrderNotification) error {
	var (
		bot *tgbotapi.BotAPI
		err error
	)

	for _, notification := range orderNotifications {
		val, ok := s.cache.Get(notification.WebAppID.String())
		if ok {
			bot = val.(*tgbotapi.BotAPI)
		} else {
			bot, err = tgbotapi.NewBotAPI(s.botToken)
			if err != nil {
				return errors.Wrap(err, "tgbotapi.NewBotAPI")
			}

			s.cache.SetWithTTL(notification.WebAppID.String(), bot, 0, 10*time.Minute)
		}

		nl, err := s.repo.GetAdminsNotificationList(s.ctx, notification.WebAppID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetAdminsNotificationList")
		}

		// need to get chat id's of users, who we are going to send messages
		for _, v := range nl {
			msgTxt, err := notification.BuildMessage()
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

// sendNewMarketplaceNotifications sends batch of notifications for new marketplaces
func (s *Service) sendNewMarketplaceNotifications(marketplaceNotifications []NewMarketplaceNotification) error {
	var (
		bot *tgbotapi.BotAPI
		err error
	)

	for _, notification := range marketplaceNotifications {
		val, ok := s.cache.Get(notification.ID.String())
		if ok {
			bot = val.(*tgbotapi.BotAPI)
		} else {
			bot, err = tgbotapi.NewBotAPI(s.botToken)
			if err != nil {
				return errors.Wrap(err, "tgbotapi.NewBotAPI")
			}

			s.cache.SetWithTTL(notification.ID.String(), bot, 0, 10*time.Minute)
		}

		nl, err := s.repo.GetReviewersNotificationList(s.ctx, notification.ID)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetReviewersNotificationList")
		}

		// need to get chat id's of users, who we are going to send messages
		for _, v := range nl {
			msgTxt, err := notification.BuildMessage()
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

// sendVerifiedMarketplaceNotifications sends batch of notifications for verified marketplaces
func (s *Service) sendVerifiedMarketplaceNotifications(marketplaceNotifications []VerifiedMarketplaceNotification) error {
	var (
		bot *tgbotapi.BotAPI
		err error
	)

	for _, notification := range marketplaceNotifications {
		val, ok := s.cache.Get(notification.ID.String())
		if ok {
			bot = val.(*tgbotapi.BotAPI)
		} else {
			bot, err = tgbotapi.NewBotAPI(s.botToken)
			if err != nil {
				return errors.Wrap(err, "tgbotapi.NewBotAPI")
			}

			s.cache.SetWithTTL(notification.ID.String(), bot, 0, 10*time.Minute)
		}

		msgTxt, err := notification.BuildMessage()
		if err != nil {
			return errors.Wrap(err, "a.BuildMessage")
		}

		msg := tgbotapi.NewMessage(notification.OwnerExternalUserID, msgTxt)
		msg.ParseMode = tgbotapi.ModeMarkdownV2
		_, err = bot.Send(msg)
		if err != nil {
			if strings.Contains(err.Error(), "Bad Request: chat not found") {
				s.log.With(
					zap.String("method", "bot.Send"),
					zap.String("user_id", strconv.FormatInt(notification.OwnerExternalUserID, 10)),
				).Warn(err.Error())
				continue
			}
			return errors.Wrap(err, "bot.Send")
		}

	}

	return nil
}

// AddUserToNewOrderNotifications creates a new order notification
// list entry for some marketplace
func (s *Service) AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsRequest) error {
	err := s.repo.AddUserToNewOrderNotifications(ctx, req)
	if err != nil {
		return errors.Wrap(err, "s.repo.AddUserToNewOrderNotifications")
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

var specialSymbols = []string{"_", "#", "-", ".", "!", "<", ">"}

func escapeSpecialSymbols(s string) string {
	for _, sym := range specialSymbols {
		if strings.Contains(s, sym) {
			s = strings.ReplaceAll(s, sym, "\\"+sym)
		}
	}

	return s
}
