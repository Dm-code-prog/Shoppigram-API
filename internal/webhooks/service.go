package webhooks

import (
	"context"
	"encoding/json"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type (
	// CreateOrUpdateTelegramChannelRequest contains the data about a Telegram channel, Shoppigram bot is added to
	CreateOrUpdateTelegramChannelRequest struct {
		ExternalID      int64
		Title           string
		Name            string
		OwnerExternalID int64
		IsPublic        bool
	}

	// ChannelStorage is an interface for storing and retrieving data about Telegram channels
	// from our own storage
	ChannelStorage interface {
		CreateOrUpdateTelegramChannel(ctx context.Context, req CreateOrUpdateTelegramChannelRequest) error
	}

	// NotifyChannelIntegrationSuccessRequest contains the data required to notify a user about a successful
	// channel integration with Shoppigram
	NotifyChannelIntegrationSuccessRequest struct {
		UserExternalID    int64
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	CloudPaymentsCheckRequest struct {
		InvoiceID       string
		Amount          float64
		Currency        string
		PaymentAmount   string
		PaymentCurrency string
		DateTime        string
	}

	CloudPaymentsCheckResponce struct {
		Code int `json:"code"`
	}

	// Notifier is the service for notifications
	// The interface requires a method for notifying a user about a successful
	// channel integration with Shoppigram
	Notifier interface {
		NotifyChannelIntegrationSuccess(ctx context.Context, request NotifyChannelIntegrationSuccessRequest) error
	}

	// Service is the service for handling Telegram webhooks
	Service struct {
		channelStorage  ChannelStorage
		notifier        Notifier
		log             *zap.Logger
		shoppigramBotID int64
	}
)

// New returns a new instance of the Service
func New(channelStorage ChannelStorage, notifier Notifier, log *zap.Logger, shoppigramBotID int64) *Service {
	return &Service{
		channelStorage:  channelStorage,
		notifier:        notifier,
		log:             log,
		shoppigramBotID: shoppigramBotID,
	}
}

// HandleTelegramWebhook is the entry point for a webhook request from Telegram.
//
// It is supposed to determine the type of Update and call the correct handler for the update.
// If there is no handler for the update, it should return nil and log the update as JSON.
//
// We differentiate the updates based on the optional fields of the Update struct. At most one optional field is present
// at any given update. However, we can have more than one handler for a given Telegram update type.
// In this case, each handler provides a function that determines if it can handle the update.
func (s *Service) HandleTelegramWebhook(ctx context.Context, update tgbotapi.Update) error {
	switch {
	case s.isUpdateTypeShoppigramBotAddedToChannelAsAdmin(update):
		return s.handleUpdateTypeShoppigramBotAddedToChannelAsAdmin(ctx, update)
	default:
		b, err := json.MarshalIndent(update, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshal data")
		}
		s.log.Info(
			"received Telegram webhook, but we don't have a handler for this type of update",
			zap.Any("webhook_data", json.RawMessage(b)),
		)
	}

	return nil
}

func (s *Service) handleUpdateTypeShoppigramBotAddedToChannelAsAdmin(ctx context.Context, update tgbotapi.Update) error {
	event := update.MyChatMember

	err := s.channelStorage.CreateOrUpdateTelegramChannel(ctx, CreateOrUpdateTelegramChannelRequest{
		ExternalID:      event.Chat.ID,
		Title:           event.Chat.Title,
		Name:            event.Chat.UserName,
		OwnerExternalID: event.From.ID,
		// Private channels don't have a name (username), only title
		// Public channels have a unique name (username) and a title
		IsPublic: event.Chat.UserName != "",
	})
	if err != nil {
		return errors.Wrap(err, "s.channelStorage.CreateOrUpdateTelegramChannel")
	}

	err = s.notifier.NotifyChannelIntegrationSuccess(ctx, NotifyChannelIntegrationSuccessRequest{
		UserExternalID:    event.From.ID,
		ChannelExternalID: event.Chat.ID,
		ChannelTitle:      event.Chat.Title,
		ChannelName:       event.Chat.UserName,
	})
	if err != nil {
		return errors.Wrap(err, "s.notifier.NotifyChannelIntegrationSuccess")
	}

	return nil
}

func (s *Service) isUpdateTypeShoppigramBotAddedToChannelAsAdmin(update tgbotapi.Update) bool {
	if update.MyChatMember == nil {
		return false
	}
	event := update.MyChatMember

	if event.Chat.Type != "channel" {
		return false
	}

	if event.NewChatMember.Status != "administrator" {
		return false
	}

	if event.NewChatMember.User.ID != s.shoppigramBotID {
		return false
	}

	if !event.NewChatMember.CanPostMessages {
		return false
	}

	return true
}

type (
	Repository interface {
		GetOrder(ctx context.Context, invoiceId string) (Order, error)
	}

	Order struct {
		ID        uuid.UUID
		UpdatedAt pgtype.Timestamp
		Sum       int64
	}

	CloudPaymentsService struct {
		repo Repository
	}
)

func NewCloudPaymentsService(repo Repository) *CloudPaymentsService {
	return &CloudPaymentsService{
		repo: repo,
	}
}

func (s *CloudPaymentsService) HandleCloudPaymentsWebHook(ctx context.Context, request http.Request) error {
	// Pass raw data here
	// Check what type of request is this
	// Pass data to handler
	// Parse data into structure in handler

	checkRequest, err := parseCloudPaymentsCheckRequest(request)
	if err == nil {
		InvoiceID := checkRequest.InvoiceID
		order, err := s.repo.GetOrder(ctx, InvoiceID) // Change here!!!
		if err != nil {
			return errors.Wrap(err, "s.repo.GetOrder()")
		}
		return handleCloudPaymentsCheckWebHook(ctx, checkRequest, order)
	}

	return errors.New("No handler for this request")
}

func handleCloudPaymentsCheckWebHook(ctx context.Context, check CloudPaymentsCheckRequest, orderInfo Order) error {
	if check.Amount != float64(orderInfo.Sum) {
		return errors.New("Sum is incorrect") // All checks here
	}
	return nil
}

func parseCloudPaymentsCheckRequest(request http.Request) (CloudPaymentsCheckRequest, error) {
	p := make([]byte, 1024)
	_, err := request.Body.Read(p)
	if err != nil {
		return CloudPaymentsCheckRequest{}, errors.Wrap(err, "")
	}
	var checkRequest map[string]interface{}

	err = json.Unmarshal(p, &checkRequest)
	if err != nil {
		// error
	}

	InvoiceId, ok := checkRequest["InvoiceId"].(string)
	if !ok {
		return CloudPaymentsCheckRequest{}, errors.New("Not check request")
	}
	Amount, ok := checkRequest["Amount"].(float64)
	if !ok {
		return CloudPaymentsCheckRequest{}, errors.New("Not check request")
	}
	Currency, ok := checkRequest["Currency"].(string)
	if !ok {
		return CloudPaymentsCheckRequest{}, errors.New("Not check request")
	}
	PaymentAmount, ok := checkRequest["PaymentAmount"].(string)
	if !ok {
		return CloudPaymentsCheckRequest{}, errors.New("Not check request")
	}
	PaymentCurrency, ok := checkRequest["PaymentCurrency"].(string)
	if !ok {
		return CloudPaymentsCheckRequest{}, errors.New("Not check request")
	}
	DateTime, ok := checkRequest["DateTime"].(string)
	if !ok {
		return CloudPaymentsCheckRequest{}, errors.New("Not check request")
	}

	return CloudPaymentsCheckRequest{
		InvoiceID:       InvoiceId,
		Amount:          Amount,
		Currency:        Currency,
		PaymentAmount:   PaymentAmount,
		PaymentCurrency: PaymentCurrency,
		DateTime:        DateTime,
	}, nil
}
