package webhooks

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
	"time"
)

type (

	// CloudPaymentsCheckRequest represents needed fields from check request from CloudPayments
	// https://developers.cloudpayments.ru/#check
	CloudPaymentsCheckRequest struct {
		TransactionId     int64   `json:"TransactionId"`               // Номер транзакции в системе
		Amount            float64 `json:"Amount"`                      // Сумма оплаты из параметров платежа
		Currency          string  `json:"Currency"`                    // Валюта: RUB/USD/EUR/GBP из параметров платежа (см. справочник)
		PaymentAmount     string  `json:"PaymentAmount"`               // Сумма списания
		PaymentCurrency   string  `json:"PaymentCurrency"`             // Валюта списания
		DateTime          string  `json:"DateTime"`                    // Дата/время создания платежа во временной зоне UTC
		CardId            *string `json:"CardId,omitempty"`            // Уникальный идентификатор карты в системе CloudPayments
		CardFirstSix      string  `json:"CardFirstSix"`                // Первые 6 цифр номера карты
		CardLastFour      string  `json:"CardLastFour"`                // Последние 4 цифры номера карты
		CardType          string  `json:"CardType"`                    // Платежная система карты: Visa, Mastercard, Maestro или "МИР"
		CardExpDate       string  `json:"CardExpDate"`                 // Срок действия карты в формате MM/YY
		TestMode          bool    `json:"TestMode"`                    // Признак тестового режима (Bit - 1 или 0)
		Status            string  `json:"Status"`                      // Статус платежа в случае успешного завершения
		OperationType     string  `json:"OperationType"`               // Тип операции: Payment/Refund/CardPayout
		InvoiceId         *string `json:"InvoiceId,omitempty"`         // Номер заказа из параметров платежа
		AccountId         *string `json:"AccountId,omitempty"`         // Идентификатор пользователя из параметров платежа
		SubscriptionId    *string `json:"SubscriptionId,omitempty"`    // Идентификатор подписки (для рекуррентных платежей)
		TokenRecipient    *string `json:"TokenRecipient,omitempty"`    // Токен получателя платежа
		Name              *string `json:"Name,omitempty"`              // Имя держателя карты
		Email             *string `json:"Email,omitempty"`             // E-mail адрес плательщика
		IpAddress         *string `json:"IpAddress,omitempty"`         // IP-адрес плательщика
		IpCountry         *string `json:"IpCountry,omitempty"`         // Двухбуквенный код страны нахождения плательщика по ISO3166-1
		IpCity            *string `json:"IpCity,omitempty"`            // Город нахождения плательщика
		IpRegion          *string `json:"IpRegion,omitempty"`          // Регион нахождения плательщика
		IpDistrict        *string `json:"IpDistrict,omitempty"`        // Округ нахождения плательщика
		IpLatitude        *string `json:"IpLatitude,omitempty"`        // Широта нахождения плательщика
		IpLongitude       *string `json:"IpLongitude,omitempty"`       // Долгота нахождения плательщика
		Issuer            *string `json:"Issuer,omitempty"`            // Название банка-эмитента карты
		IssuerBankCountry *string `json:"IssuerBankCountry,omitempty"` // Двухбуквенный код страны эмитента карты по ISO3166-1
		Description       *string `json:"Description,omitempty"`       // Назначение оплаты из параметров платежа
		CardProduct       *string `json:"CardProduct,omitempty"`       // Тип карточного продукта
		PaymentMethod     *string `json:"PaymentMethod,omitempty"`     // Метод оплаты ApplePay, GooglePay, Яндекс Пэй или T-Pay
		Data              *string `json:"Data,omitempty"`              // Произвольный набор параметров, переданных в транзакцию
	}

	// CloudPaymentsPayRequest represents the fields from the "Pay" webhook from CloudPayments
	// https://developers.cloudpayments.ru/#pay
	CloudPaymentsPayRequest struct {
		TransactionId                         int64    `json:"TransactionId"`                                   // Номер транзакции в системе
		Amount                                float64  `json:"Amount"`                                          // Сумма оплаты из параметров платежа
		Currency                              string   `json:"Currency"`                                        // Валюта: RUB/USD/EUR/GBP из параметров платежа
		PaymentAmount                         string   `json:"PaymentAmount"`                                   // Сумма списания
		PaymentCurrency                       string   `json:"PaymentCurrency"`                                 // Валюта списания
		DateTime                              string   `json:"DateTime"`                                        // Дата/время создания платежа во временной зоне UTC
		CardId                                *string  `json:"CardId,omitempty"`                                // Уникальный идентификатор карты в системе CloudPayments
		CardFirstSix                          string   `json:"CardFirstSix"`                                    // Первые 6 цифр номера карты
		CardLastFour                          string   `json:"CardLastFour"`                                    // Последние 4 цифры номера карты
		CardType                              string   `json:"CardType"`                                        // Платежная система карты: Visa, Mastercard, Maestro или "МИР"
		CardExpDate                           string   `json:"CardExpDate"`                                     // Срок действия карты в формате MM/YY
		TestMode                              bool     `json:"TestMode"`                                        // Признак тестового режима (Bit - 1 или 0)
		Status                                string   `json:"Status"`                                          // Статус платежа после авторизации
		OperationType                         string   `json:"OperationType"`                                   // Тип операции: Payment/CardPayout
		GatewayName                           string   `json:"GatewayName"`                                     // Идентификатор банка-эквайера
		InvoiceId                             *string  `json:"InvoiceId,omitempty"`                             // Номер заказа из параметров платежа
		AccountId                             *string  `json:"AccountId,omitempty"`                             // Идентификатор пользователя из параметров платежа
		SubscriptionId                        *string  `json:"SubscriptionId,omitempty"`                        // Идентификатор подписки (для рекуррентных платежей)
		Name                                  *string  `json:"Name,omitempty"`                                  // Имя держателя карты
		Email                                 *string  `json:"Email,omitempty"`                                 // E-mail адрес плательщика
		IpAddress                             *string  `json:"IpAddress,omitempty"`                             // IP-адрес плательщика
		IpCountry                             *string  `json:"IpCountry,omitempty"`                             // Двухбуквенный код страны нахождения плательщика по ISO3166-1
		IpCity                                *string  `json:"IpCity,omitempty"`                                // Город нахождения плательщика
		IpRegion                              *string  `json:"IpRegion,omitempty"`                              // Регион нахождения плательщика
		IpDistrict                            *string  `json:"IpDistrict,omitempty"`                            // Округ нахождения плательщика
		IpLatitude                            *string  `json:"IpLatitude,omitempty"`                            // Широта нахождения плательщика
		IpLongitude                           *string  `json:"IpLongitude,omitempty"`                           // Долгота нахождения плательщика
		Issuer                                *string  `json:"Issuer,omitempty"`                                // Название банка-эмитента карты
		IssuerBankCountry                     *string  `json:"IssuerBankCountry,omitempty"`                     // Двухбуквенный код страны эмитента карты по ISO3166-1
		Description                           *string  `json:"Description,omitempty"`                           // Назначение оплаты из параметров платежа
		AuthCode                              *string  `json:"AuthCode,omitempty"`                              // Код авторизации
		Data                                  *string  `json:"Data,omitempty"`                                  // Произвольный набор параметров, переданных в транзакцию
		Token                                 *string  `json:"Token,omitempty"`                                 // Токен карты для повторных платежей без ввода реквизитов
		TotalFee                              *float64 `json:"TotalFee,omitempty"`                              // Значение общей комиссии
		CardProduct                           *string  `json:"CardProduct,omitempty"`                           // Тип карточного продукта
		PaymentMethod                         *string  `json:"PaymentMethod,omitempty"`                         // Метод оплаты ApplePay, GooglePay, Яндекс Пэй или T-Pay
		FallBackScenarioDeclinedTransactionId *int64   `json:"FallBackScenarioDeclinedTransactionId,omitempty"` // Номер первой неуспешной транзакции
		Rrn                                   *string  `json:"Rrn,omitempty"`                                   // Уникальный номер банковской транзакции, который назначается обслуживающим банком
	}

	// CloudPaymentsResponse represents check response for CloudPayments
	CloudPaymentsResponse struct {
		Code int8 `json:"code,omitempty"`
	}

	// Order represents order record in database
	Order struct {
		ID        uuid.UUID
		UpdatedAt time.Time
		Sum       float64
		Currency  string
		State     string
	}

	// SavePaymentExtraInfoParams represents the parameters for saving extra info about the payment
	SavePaymentExtraInfoParams struct {
		OrderID            uuid.UUID
		Provider           string
		OrderStateSnapshot string
		EventType          string
		ExtraInfo          []byte
	}

	// Repository provides access to the webhooks storage
	Repository interface {
		GetOrder(ctx context.Context, id uuid.UUID) (Order, error)
		SetOrderStateConfirmed(ctx context.Context, id uuid.UUID) error
		SavePaymentExtraInfo(ctx context.Context, params SavePaymentExtraInfoParams) error
	}

	// CloudPaymentsService is the service for handling CloudPayments webhooks
	CloudPaymentsService struct {
		repo                          Repository
		maxDurationForHandlingPayment time.Duration
		log                           *zap.Logger
	}
)

const (
	providerCloudPayments = "cloud_payments"

	eventTypeCheck = "check"
	eventTypePay   = "pay"

	orderStateConfirmed = "confirmed"

	// cloudPaymentsResponseCodeSuccess 0 means success for all webhook types
	cloudPaymentsResponseCodeSuccess                   = 0
	cloudPaymentsCheckResponseCodeWrongInvoiceID       = 10
	cloudPaymentsCheckResponseCodeWrongSum             = 12
	cloudPaymentsCheckResponseCodeCantHandleThePayment = 13
	cloudPaymentsCheckResponseCodeTransactionExpired   = 20
)

// NewCloudPayments returns a new instance of CloudPaymentsService
func NewCloudPayments(repo Repository, log *zap.Logger, maxDurationForHandlingPayment time.Duration) *CloudPaymentsService {
	return &CloudPaymentsService{
		repo:                          repo,
		log:                           log,
		maxDurationForHandlingPayment: maxDurationForHandlingPayment,
	}
}

// HandleCloudPaymentsCheckWebHook is the entry point for a webhook request from CloudPayments
//
// It suppose to determine, what type of request was made, and generate a response
func (s *CloudPaymentsService) HandleCloudPaymentsCheckWebHook(ctx context.Context, checkRequest CloudPaymentsCheckRequest) (resp CloudPaymentsResponse, err error) {
	var invoiceID string
	if checkRequest.InvoiceId != nil {
		invoiceID = *checkRequest.InvoiceId
	} else {
		return CloudPaymentsResponse{Code: cloudPaymentsCheckResponseCodeWrongInvoiceID}, nil
	}

	id, err := uuid.Parse(invoiceID)
	if err != nil {
		return CloudPaymentsResponse{Code: cloudPaymentsCheckResponseCodeWrongInvoiceID}, nil
	}

	order, err := s.repo.GetOrder(ctx, id)
	if err != nil {
		if errors.Is(err, ErrorOrderDoesntExist) {
			return CloudPaymentsResponse{Code: cloudPaymentsCheckResponseCodeWrongInvoiceID}, nil
		}
		return CloudPaymentsResponse{Code: cloudPaymentsCheckResponseCodeCantHandleThePayment}, errors.Wrap(err, "s.repo.GetOrder")
	}

	extraInfoJSON, err := json.Marshal(checkRequest)
	if err != nil {
		return CloudPaymentsResponse{
			Code: cloudPaymentsCheckResponseCodeCantHandleThePayment,
		}, nil
	}

	// save the extra info
	err = s.repo.SavePaymentExtraInfo(ctx, SavePaymentExtraInfoParams{
		OrderID:            order.ID,
		Provider:           providerCloudPayments,
		OrderStateSnapshot: order.State,
		EventType:          eventTypeCheck,
		ExtraInfo:          extraInfoJSON,
	})
	if err != nil {
		return CloudPaymentsResponse{
			Code: cloudPaymentsCheckResponseCodeCantHandleThePayment,
		}, errors.Wrap(err, "s.repo.SavePaymentExtraInfo")
	}

	return CloudPaymentsResponse{
		Code: int8(checkPayment(checkRequest, order, s.maxDurationForHandlingPayment)),
	}, nil
}

// HandleCloudPaymentsPayWebHook is the entry point for a webhook request from CloudPayments
func (s *CloudPaymentsService) HandleCloudPaymentsPayWebHook(ctx context.Context, payRequest CloudPaymentsPayRequest) (resp CloudPaymentsResponse, err error) {
	var invoiceID string
	if payRequest.InvoiceId != nil {
		invoiceID = *payRequest.InvoiceId
	} else {
		return CloudPaymentsResponse{}, errors.New("invoiceId is nil")
	}

	id, err := uuid.Parse(invoiceID)
	if err != nil {
		return CloudPaymentsResponse{}, errors.Wrap(err, "uuid.Parse")
	}

	err = s.repo.SetOrderStateConfirmed(ctx, id)
	if err != nil {
		return CloudPaymentsResponse{}, errors.Wrap(err, "s.repo.SetOrderStateConfirmed")
	}

	extraInfoJSON, err := json.Marshal(payRequest)
	if err != nil {
		return CloudPaymentsResponse{}, errors.Wrap(err, "json.Marshal")
	}

	// save the extra info
	err = s.repo.SavePaymentExtraInfo(ctx, SavePaymentExtraInfoParams{
		OrderID:            id,
		Provider:           providerCloudPayments,
		OrderStateSnapshot: orderStateConfirmed,
		EventType:          eventTypePay,
		ExtraInfo:          extraInfoJSON,
	})
	if err != nil {
		return CloudPaymentsResponse{}, errors.Wrap(err, "s.repo.SavePaymentExtraInfo")
	}

	return CloudPaymentsResponse{Code: cloudPaymentsResponseCodeSuccess}, nil

}

func checkPayment(check CloudPaymentsCheckRequest, orderInfo Order, paymentMaxDuration time.Duration) int {
	if check.InvoiceId == nil {
		return cloudPaymentsCheckResponseCodeWrongInvoiceID
	}

	if *check.InvoiceId != orderInfo.ID.String() {
		return cloudPaymentsCheckResponseCodeWrongInvoiceID
	}

	if check.Amount != orderInfo.Sum || !isCurrenciesEqual(check.Currency, orderInfo.Currency) {
		return cloudPaymentsCheckResponseCodeWrongSum
	}
	orderUpdateTime := orderInfo.UpdatedAt
	paymentTime, err := time.Parse(time.DateTime, check.DateTime)
	if err != nil {
		return cloudPaymentsCheckResponseCodeCantHandleThePayment
	}
	if isPaymentExpired(orderUpdateTime, paymentTime, paymentMaxDuration) {
		return cloudPaymentsCheckResponseCodeTransactionExpired
	}
	return cloudPaymentsResponseCodeSuccess
}

func isCurrenciesEqual(cur1 string, cur2 string) bool {
	return strings.ToLower(cur1) == strings.ToLower(cur2)
}

func isPaymentExpired(orderCreated time.Time, paymentWasMade time.Time, maxDuration time.Duration) bool {
	return paymentWasMade.Sub(orderCreated) > maxDuration
}
