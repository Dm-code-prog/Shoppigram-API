package notifications

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

	// ChannelIntegrationSuccessNotification defines the structure of a successful channel integration notification
	ChannelIntegrationSuccessNotification struct {
		UserExternalID    int64
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}
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
		escapeSpecialSymbols(string(newOrderMessageTemplate)),
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

// BuildMessage creates a notification message for a successful channel integration
func (m *ChannelIntegrationSuccessNotification) BuildMessage() (string, error) {
	channelIntegrationSuccessMessageTemplate, err := templates.ReadFile("templates/channel_integration_success.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	tmaLink, err := TMALinkingScheme{
		PageName: "/admin/new-marketplace",
		PageData: map[string]any{
			"integrated_tg_channel_name":        m.ChannelName,
			"integrated_tg_channel_external_id": m.ChannelExternalID,
		},
	}.ToBase64String()
	if err != nil {
		return "", errors.Wrap(err, "TMALinkingScheme.ToBase64String")
	}

	return fmt.Sprintf(
		escapeSpecialSymbols(string(channelIntegrationSuccessMessageTemplate)),
		escapeSpecialSymbols(m.ChannelTitle),
		escapeSpecialSymbols(tmaLink),
	), nil
}
