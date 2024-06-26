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

// BuildMessageAdmin creates a notification message for a new order for an admin
func (o *NewOrderNotification) BuildMessageAdmin() (string, error) {
	var subtotal float64
	var productList strings.Builder
	var currency string
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		currency = p.PriceCurrency
		productList.WriteString(fmt.Sprintf(`\- %dx %s по цене %s %s
`, p.Quantity, escapeSpecialSymbols(p.Name), escapeSpecialSymbols(formatFloat(p.Price)), formatCurrency(p.PriceCurrency)))
	}

	newOrderMessageTemplate, err := templates.ReadFile("templates/new_order_message.admin.md")
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

// BuildMessageCustomer creates a notification message for a new order for a customer
func (o *NewOrderNotification) BuildMessageCustomer() (string, error) {
	var subtotal float64
	var productList strings.Builder
	var currency string
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		currency = p.PriceCurrency
		productList.WriteString(fmt.Sprintf(`\- %dx %s по цене %s %s
`, p.Quantity, escapeSpecialSymbols(p.Name), escapeSpecialSymbols(formatFloat(p.Price)), formatCurrency(p.PriceCurrency)))
	}

	newOrderMessageTemplate, err := templates.ReadFile("templates/new_order_message.customer.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	return fmt.Sprintf(
		escapeSpecialSymbols(string(newOrderMessageTemplate)),
		escapeSpecialSymbols(o.WebAppName),
		escapeSpecialSymbols(formatFloat(subtotal))+" "+formatCurrency(currency),
		strings.TrimRight(productList.String(), "; "),
	), nil
}


// BuildMessageShoppigram creates a notification message for a new marketplace
func (m *NewMarketplaceNotification) BuildMessageShoppigram() (string, error) {
	newMarketplaceMessageTemplate, err := templates.ReadFile("templates/marketplace_needs_verification.shoppigram.md")
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

// BuildMessageAdmin creates a notification message for a marketplace on verification
func (m *NewMarketplaceNotification) BuildMessageAdmin() (string, error) {
	marketplaceVerificationMessageTemplate, err := templates.ReadFile("templates/marketplace_sent_for_verification.admin.md")
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}
	tmaLink, err := TMALinkingScheme{
		PageName: "/admin/marketplaces/" + m.ID.String(),
		PageData: map[string]any{},
	}.ToBase64String()
	if err != nil {
		return "", errors.Wrap(err, "TMALinkingScheme.ToBase64String")
	}

	return fmt.Sprintf(
		escapeSpecialSymbols(string(marketplaceVerificationMessageTemplate)),
		escapeSpecialSymbols(m.Name),
		escapeSpecialSymbols(tmaLink),
	), nil
}

// BuildMessage creates a notification message for a verified marketplace
func (m *VerifiedMarketplaceNotification) BuildMessage() (string, error) {
	verifiedMarketplaceMessageTemplate, err := templates.ReadFile("templates/marketplace_verified.admin.md")
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
		escapeSpecialSymbols(string(channelIntegrationSuccessMessageTemplate)),
		escapeSpecialSymbols(m.ChannelTitle),
		escapeSpecialSymbols(tmaLink),
	), nil
}
