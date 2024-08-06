package notifications

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

type (
	// NewOrderNotification defines the structure of order notification
	NewOrderNotification struct {
		ID              uuid.UUID
		ReadableOrderID int64
		CreatedAt       time.Time
		UserNickname    string
		WebAppID        uuid.UUID
		WebAppName      string
		Products        []Product
		status          string
		ExternalUserID  int64
	}

	// NewMarketplaceNotification defines the structure of new marketplace notification
	NewMarketplaceNotification struct {
		ID              uuid.UUID
		Name            string
		ShortName       string
		CreatedAt       time.Time
		OwnerUsername   string
		OwnerExternalID int64
		ImageBaseUrl    string
	}

	// VerifiedMarketplaceNotification defines the structure of verified marketplace notification
	VerifiedMarketplaceNotification struct {
		ID                  uuid.UUID
		Name                string
		ShortName           string
		VerifiedAt          time.Time
		OwnerExternalUserID int64
	}

	// ChannelIntegrationSuccessNotification defines the structure of a successful channel integration notification
	ChannelIntegrationSuccessNotification struct {
		UserExternalID    int64
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}
)

var botName = os.Getenv("BOT_NAME")

// BuildMessageAdmin creates a notification message for a new order for an admin
func (o *NewOrderNotification) BuildMessageAdmin() (string, error) {
	var (
		subtotal    float64
		productList strings.Builder
		currency    string
	)
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		currency = p.PriceCurrency
		productList.WriteString(fmt.Sprintf(`- %dx %s по цене %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(p.PriceCurrency)))
	}

	newOrderMessageTemplate, err := templates.ReadFile("templates/admin/new_order_message.admin.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}
	finalMessage := fmt.Sprintf(
		string(newOrderMessageTemplate),
		o.WebAppName,
		o.UserNickname,
		o.ReadableOrderID,
		formatFloat(subtotal)+" "+formatCurrency(currency),
		"status",
		formatRussianTime(o.CreatedAt),
		"no comment",
		strings.TrimRight(productList.String(), "; "),
	)
	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

// BuildMessageCustomer creates a notification message for a new order for a customer
func (o *NewOrderNotification) BuildMessageCustomer() (string, error) {
	var subtotal float64
	var productList strings.Builder
	var currency string
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		currency = p.PriceCurrency
		productList.WriteString(fmt.Sprintf(`- %dx %s по цене %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(p.PriceCurrency)))
	}

	newOrderMessageTemplate, err := templates.ReadFile("templates/customer/new_order_message.customer.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(newOrderMessageTemplate),
		o.WebAppName,
		formatFloat(subtotal)+" "+formatCurrency(currency),
		strings.TrimRight(productList.String(), "; "),
	)

	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

// BuildMessageShoppigram creates a notification message for a new marketplace
func (m *NewMarketplaceNotification) BuildMessageShoppigram() (string, error) {
	newMarketplaceMessageTemplate, err := templates.ReadFile("templates/shoppigram/marketplace_needs_verification.shoppigram.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(newMarketplaceMessageTemplate),
		m.OwnerUsername,
		m.Name,
		m.ShortName,
		m.ImageBaseUrl+"/"+m.ShortName+"/logo",
		marketplaceURL+m.ID.String(),
	)

	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

// BuildMessageAdmin creates a notification message for a marketplace on verification
func (m *NewMarketplaceNotification) BuildMessageAdmin() (string, error) {
	marketplaceVerificationMessageTemplate, err := templates.ReadFile("templates/admin/marketplace_sent_for_verification.admin.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(marketplaceVerificationMessageTemplate),
		m.Name,
	)

	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

// BuildMessage creates a notification message for a verified marketplace
func (m *VerifiedMarketplaceNotification) BuildMessage() (string, error) {
	verifiedMarketplaceMessageTemplate, err := templates.ReadFile("templates/marketplace_verified.admin.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	return fmt.Sprintf(
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, string(verifiedMarketplaceMessageTemplate)),
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, m.Name),
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, "https://t.me/"+botName+"/"+m.ShortName),
	), nil
}

// BuildMessage creates a notification message for a successful channel integration
func (m *ChannelIntegrationSuccessNotification) BuildMessage() (string, error) {
	channelIntegrationSuccessMessageTemplate, err := templates.ReadFile("templates/channel_integrated.admin.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	tmaLink, err := TMALinkingScheme{
		PageName: "/admin",
		PageData: map[string]any{},
	}.ToBase64String()
	if err != nil {
		return "", errors.Wrap(err, "TMALinkingScheme.ToBase64String")
	}

	return fmt.Sprintf(
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, string(channelIntegrationSuccessMessageTemplate)),
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, m.ChannelTitle),
		tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, "https://t.me/"+botName+"/"+"app?startapp="+tmaLink),
	), nil
}
