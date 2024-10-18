package notifications

import (
	"fmt"
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
		UserLanguage    string
		OwnerLanguage   string
		WebAppID        uuid.UUID
		WebAppName      string
		WebAppCurrency  string
		Products        []Product
		Status          string
		Comment         string
		PaymentType     string
		ExternalUserID  int64
	}

	// OrderNotification defines the structure of order notification
	OrderNotification struct {
		ID              uuid.UUID
		ReadableOrderID int64
		CreatedAt       time.Time
		BuyerNickname   string
		BuyerLanguage   string
		OwnerLanguage   string
		WebAppID        uuid.UUID
		WebAppName      string
		WebAppCurrency  string
		Products        []Product
		Status          string
		Comment         string
		PaymentType     string
		BuyerExternalID int64
	}

	// NewShopNotification defines the structure of new marketplace notification
	NewShopNotification struct {
		ID              uuid.UUID
		Name            string
		ShortName       string
		CreatedAt       time.Time
		OwnerUsername   string
		OwnerExternalID int64
		OwnerLanguage   string
	}

	// VerifiedShopNotification defines the structure of verified marketplace notification
	VerifiedShopNotification struct {
		ID                  uuid.UUID
		Name                string
		ShortName           string
		VerifiedAt          time.Time
		OwnerExternalUserID int64
		OwnerLanguage       string
	}

	// ChannelIntegrationSuccessNotification defines the structure of a successful channel integration notification
	ChannelIntegrationSuccessNotification struct {
		UserExternalID    int64
		UserLanguage      string
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}

	// ChannelIntegrationFailureNotification defines the structure of a failed channel integration notification
	ChannelIntegrationFailureNotification struct {
		UserExternalID    int64
		UserLanguage      string
		ChannelExternalID int64
		ChannelTitle      string
		ChannelName       string
	}
)

const (
	pathToAdminChannelIntegrated              = "admin/channel_integrated.md"
	pathToAdminChannelIntegrationFailure      = "admin/channel_integration_failure.md"
	pathToAdminGreetings                      = "admin/greetings_message.md"
	pathToAdminMarketplaceSentForVerification = "admin/marketplace_sent_for_verification.md"
	pathToAdminMarketplaceVerified            = "admin/marketplace_verified.md"

	pathToShoppigramMarketplaceNeedsVerification = "shoppigram/marketplace_needs_verification.md"

	pathToOrderConfirmedAdmin = "admin/order_confirmed.md"
	pathToOrderConfirmedBuyer = "customer/order_confirmed.md"
	pathToOrderDoneAdmin      = "admin/order_done.md"
	pathToOrderDoneBuyer      = "customer/order_done.md"

	langRu = "ru"
	langEn = "en"

	orderTypeP2P    = "p2p"
	orderTypeOnline = "online"
)

var botName = os.Getenv("BOT_NAME")

// MakeConfirmedNotificationForAdmin creates a notification message for a new order for an admin
func (o *OrderNotification) MakeConfirmedNotificationForAdmin(language string) (string, error) {
	var (
		subtotal    float64
		productList strings.Builder
		currency    string
	)
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		currency = o.WebAppCurrency
		productList.WriteString(fmt.Sprintf(`- %dx %s | %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(o.WebAppCurrency)))
	}

	template, err := templates.ReadFile(getPathToTemplate(language, pathToOrderConfirmedAdmin))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	comment := o.Comment
	if comment == "" {
		comment = getTranslation(language, "empty-comment")
	}

	var paymentStatus string
	if o.PaymentType == orderTypeP2P {
		paymentStatus = getTranslation(language, "payment-status-unpaid")
	} else if o.PaymentType == orderTypeOnline {
		paymentStatus = getTranslation(language, "payment-status-paid")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		o.WebAppName,
		"@"+o.BuyerNickname,
		o.ReadableOrderID,
		o.PaymentType,
		formatFloat(subtotal)+" "+formatCurrency(currency),
		paymentStatus,
		formatRussianTime(o.CreatedAt),
		comment,
		strings.TrimRight(productList.String(), "; "),
	)
	return finalMessage, nil
}

// MakeConfirmedNotificationForBuyer creates a notification message for a new order for a customer
func (o *OrderNotification) MakeConfirmedNotificationForBuyer(language string) (string, error) {
	var subtotal float64
	var productList strings.Builder
	currency := ""
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		if currency == "" {
			currency = o.WebAppCurrency
		}
		productList.WriteString(fmt.Sprintf(`- %dx %s: %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(o.WebAppCurrency)))
	}

	template, err := templates.ReadFile(getPathToTemplate(language, pathToOrderConfirmedBuyer))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	var paymentStatus string
	if o.PaymentType == orderTypeP2P {
		paymentStatus = getTranslation(language, "payment-status-unpaid")
	} else if o.PaymentType == orderTypeOnline {
		paymentStatus = getTranslation(language, "payment-status-paid")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		o.WebAppName,
		o.ReadableOrderID,
		formatFloat(subtotal)+" "+formatCurrency(currency),
		paymentStatus,
		strings.TrimRight(productList.String(), "; "),
	)

	return finalMessage, nil
}

// MakeDoneNotificationForAdmin creates a notification message for a done order for an admin
func (o *OrderNotification) MakeDoneNotificationForAdmin(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToOrderDoneAdmin))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	paymentStatus := getTranslation(language, "payment-status-paid")

	var subtotal float64
	var productList strings.Builder
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		productList.WriteString(fmt.Sprintf(`- %dx %s: %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(o.WebAppCurrency)))
	}

	finalMessage := fmt.Sprintf(
		string(template),
		o.WebAppName,
		"@"+o.BuyerNickname,
		o.ReadableOrderID,
		formatFloat(subtotal)+" "+formatCurrency(o.WebAppCurrency),
		paymentStatus,
		strings.TrimRight(productList.String(), "; "),
	)

	return finalMessage, nil
}

func (o *OrderNotification) MakeDoneNotificationForBuyer(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToOrderDoneBuyer))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	paymentStatus := getTranslation(language, "payment-status-paid")

	var subtotal float64
	var productList strings.Builder
	for _, p := range o.Products {
		subtotal += p.Price * float64(p.Quantity)
		productList.WriteString(fmt.Sprintf(`- %dx %s: %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(o.WebAppCurrency)))
	}

	finalMessage := fmt.Sprintf(
		string(template),
		o.WebAppName,
		o.ReadableOrderID,
		formatFloat(subtotal)+" "+formatCurrency(o.WebAppCurrency),
		paymentStatus,
		strings.TrimRight(productList.String(), "; "),
	)

	return finalMessage, nil
}

// BuildMessageShoppigram creates a notification message for a new marketplace
func (m *NewShopNotification) BuildMessageShoppigram(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToShoppigramMarketplaceNeedsVerification))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		m.OwnerUsername,
		m.Name,
		m.ShortName,
		marketplaceBaseURL+m.ID.String(),
	)

	return finalMessage, nil
}

// BuildMessageAdmin creates a notification message for a marketplace on verification
func (m *NewShopNotification) BuildMessageAdmin(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToAdminMarketplaceSentForVerification))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		m.Name,
	)

	return finalMessage, nil
}

// BuildMessage creates a notification message for a verified marketplace
func (m *VerifiedShopNotification) BuildMessage(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToAdminMarketplaceVerified))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		m.Name,
		makeShopURL(botName, m.ShortName),
	)
	return finalMessage, nil
}

// BuildMessage creates a notification message for a successful channel integration
func (m *ChannelIntegrationSuccessNotification) BuildMessage(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToAdminChannelIntegrated))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		m.ChannelTitle,
	)

	return finalMessage, nil
}

func BuildGreetigsMessage(language string) (string, error) {
	greetingsMessage, err := templates.ReadFile(getPathToTemplate(language, pathToAdminGreetings))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}
	return string(greetingsMessage), nil
}

// BuildMessage creates a notification message for a failed channel integration
func (m *ChannelIntegrationFailureNotification) BuildMessage(language string) (string, error) {
	template, err := templates.ReadFile(getPathToTemplate(language, pathToAdminChannelIntegrationFailure))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(template),
		botName,
		m.ChannelTitle,
	)

	return finalMessage, nil
}

func getPathToTemplate(lang string, path string) string {
	return "templates/" + lang + "/" + path
}
