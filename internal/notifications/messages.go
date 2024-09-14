package notifications

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/templates/en"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/templates/ru"
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

	// NewMarketplaceNotification defines the structure of new marketplace notification
	NewMarketplaceNotification struct {
		ID              uuid.UUID
		Name            string
		ShortName       string
		CreatedAt       time.Time
		OwnerUsername   string
		OwnerExternalID int64
		OwnerLanguage   string
		ImageBaseUrl    string
	}

	// VerifiedMarketplaceNotification defines the structure of verified marketplace notification
	VerifiedMarketplaceNotification struct {
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
)

const (
	pathToAdminChannelIntegrated              = "admin/channel_integrated.md"
	pathToAdminGreetings                      = "admin/greetings_message.md"
	pathToAdminMarketplaceSentForVerification = "admin/marketplace_sent_for_verification.md"
	pathToAdminMarketplaceVerified            = "admin/marketplace_verified.md"

	pathToShoppigramMarketplaceNeedsVerification = "shoppigram/marketplace_needs_verification.md"

	pathToOrderConfirmedAdmin = "admin/order_confirmed.md"
	pathToOrderConfirmedBuyer = "customer/order_confirmed.md"
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
		productList.WriteString(fmt.Sprintf(`- %dx %s по цене %s %s
`, p.Quantity, p.Name, formatFloat(p.Price), formatCurrency(o.WebAppCurrency)))
	}

	newOrderMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToOrderConfirmedAdmin))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	comment := o.Comment
	if comment == "" {
		if language == langRu {
			comment = ru.Translations["empty-comment"]
		} else if language == langEn {
			comment = en.Translations["empty-comment"]
		}
	}

	var paymentStatus string
	if o.PaymentType == orderTypeP2P {
		if language == langRu {
			paymentStatus = ru.Translations["payment-status-unpaid"]
		} else if language == langEn {
			paymentStatus = en.Translations["payment-status-unpaid"]
		}
	} else if o.PaymentType == orderTypeOnline {
		if language == langRu {
			paymentStatus = ru.Translations["payment-status-paid"]
		} else if language == langEn {
			paymentStatus = en.Translations["payment-status-paid"]
		}
	}

	finalMessage := fmt.Sprintf(
		string(newOrderMessageTemplate),
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
	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
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

	newOrderMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToOrderConfirmedBuyer))
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
		string(newOrderMessageTemplate),
		o.WebAppName,
		o.ReadableOrderID,
		formatFloat(subtotal)+" "+formatCurrency(currency),
		paymentStatus,
		strings.TrimRight(productList.String(), "; "),
	)

	return finalMessage, nil
}

func (o *OrderNotification) MakeDoneNotificationForBuyer(language string) (string, error) {
	newOrderMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToOrderDoneBuyer))
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
		string(newOrderMessageTemplate),
		o.WebAppName,
		o.ReadableOrderID,
		formatFloat(subtotal)+" "+formatCurrency(o.WebAppCurrency),
		paymentStatus,
		strings.TrimRight(productList.String(), "; "),
	)

	return finalMessage, nil
}

// BuildMessageShoppigram creates a notification message for a new marketplace
func (m *NewMarketplaceNotification) BuildMessageShoppigram(language string) (string, error) {
	newMarketplaceMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToShoppigramMarketplaceNeedsVerification))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(newMarketplaceMessageTemplate),
		m.OwnerUsername,
		m.Name,
		m.ShortName,
		m.ImageBaseUrl+"/"+m.ShortName+"/logo",
		marketplaceBaseURL+m.ID.String(),
	)

	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

// BuildMessageAdmin creates a notification message for a marketplace on verification
func (m *NewMarketplaceNotification) BuildMessageAdmin(language string) (string, error) {
	marketplaceVerificationMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToAdminMarketplaceSentForVerification))
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
func (m *VerifiedMarketplaceNotification) BuildMessage(language string) (string, error) {
	verifiedMarketplaceMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToAdminMarketplaceVerified))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(verifiedMarketplaceMessageTemplate),
		m.Name,
		"https://t.me/"+botName+"/"+m.ShortName,
	)
	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

// BuildMessage creates a notification message for a successful channel integration
func (m *ChannelIntegrationSuccessNotification) BuildMessage(language string) (string, error) {
	channelIntegrationSuccessMessageTemplate, err := templates.ReadFile(getPathToFile(language, pathToAdminChannelIntegrated))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}

	finalMessage := fmt.Sprintf(
		string(channelIntegrationSuccessMessageTemplate),
		m.ChannelTitle,
	)

	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, finalMessage), nil
}

func BuildGreetigsMessage(language string) (string, error) {
	greetingsMessage, err := templates.ReadFile(getPathToFile(language, pathToAdminGreetings))
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile")
	}
	return tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, string(greetingsMessage)), nil
}

func getPathToFile(lang string, path string) string {
	return "templates/" + lang + "/" + path
}
