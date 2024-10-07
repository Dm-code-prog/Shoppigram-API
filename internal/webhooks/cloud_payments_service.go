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
		TransactionId     int64   // Номер транзакции в системе
		Amount            float64 // Сумма оплаты из параметров платежа
		Currency          string  // Валюта: RUB/USD/EUR/GBP из параметров платежа (см. справочник)
		PaymentAmount     string  // Сумма списания
		PaymentCurrency   string  // Валюта списания
		DateTime          string  // Дата/время создания платежа во временной зоне UTC
		CardId            *string // Уникальный идентификатор карты в системе CloudPayments
		CardFirstSix      string  // Первые 6 цифр номера карты
		CardLastFour      string  // Последние 4 цифры номера карты
		CardType          string  // Платежная система карты: Visa, Mastercard, Maestro или "МИР"
		CardExpDate       string  // Срок действия карты в формате MM/YY
		TestMode          bool    // Признак тестового режима (Bit - 1 или 0)
		Status            string  // Статус платежа в случае успешного завершения
		OperationType     string  // Тип операции: Payment/Refund/CardPayout
		InvoiceId         *string // Номер заказа из параметров платежа
		AccountId         *string // Идентификатор пользователя из параметров платежа
		SubscriptionId    *string // Идентификатор подписки (для рекуррентных платежей)
		TokenRecipient    *string // Токен получателя платежа
		Name              *string // Имя держателя карты
		Email             *string // E-mail адрес плательщика
		IpAddress         *string // IP-адрес плательщика
		IpCountry         *string // Двухбуквенный код страны нахождения плательщика по ISO3166-1
		IpCity            *string // Город нахождения плательщика
		IpRegion          *string // Регион нахождения плательщика
		IpDistrict        *string // Округ нахождения плательщика
		IpLatitude        *string // Широта нахождения плательщика
		IpLongitude       *string // Долгота нахождения плательщика
		Issuer            *string // Название банка-эмитента карты
		IssuerBankCountry *string // Двухбуквенный код страны эмитента карты по ISO3166-1
		Description       *string // Назначение оплаты из параметров платежа
		CardProduct       *string // Тип карточного продукта
		PaymentMethod     *string // Метод оплаты ApplePay, GooglePay, Яндекс Пэй или T-Pay
		Data              *string // Произвольный набор параметров, переданных в транзакцию
	}

	// CloudPaymentsPayRequest represents the fields from the "Pay" webhook from CloudPayments
	// https://developers.cloudpayments.ru/#pay
	CloudPaymentsPayRequest struct {
		TransactionId                         int64    // Номер транзакции в системе
		Amount                                float64  // Сумма оплаты из параметров платежа
		Currency                              string   // Валюта: RUB/USD/EUR/GBP из параметров платежа
		PaymentAmount                         string   // Сумма списания
		PaymentCurrency                       string   // Валюта списания
		DateTime                              string   // Дата/время создания платежа во временной зоне UTC
		CardId                                *string  // Уникальный идентификатор карты в системе CloudPayments
		CardFirstSix                          string   // Первые 6 цифр номера карты
		CardLastFour                          string   // Последние 4 цифры номера карты
		CardType                              string   // Платежная система карты: Visa, Mastercard, Maestro или "МИР"
		CardExpDate                           string   // Срок действия карты в формате MM/YY
		TestMode                              bool     // Признак тестового режима (Bit - 1 или 0)
		Status                                string   // Статус платежа после авторизации
		OperationType                         string   // Тип операции: Payment/CardPayout
		GatewayName                           string   // Идентификатор банка-эквайера
		InvoiceId                             *string  // Номер заказа из параметров платежа
		AccountId                             *string  // Идентификатор пользователя из параметров платежа
		SubscriptionId                        *string  // Идентификатор подписки (для рекуррентных платежей)
		Name                                  *string  // Имя держателя карты
		Email                                 *string  // E-mail адрес плательщика
		IpAddress                             *string  // IP-адрес плательщика
		IpCountry                             *string  // Двухбуквенный код страны нахождения плательщика по ISO3166-1
		IpCity                                *string  // Город нахождения плательщика
		IpRegion                              *string  // Регион нахождения плательщика
		IpDistrict                            *string  // Округ нахождения плательщика
		IpLatitude                            *string  // Широта нахождения плательщика
		IpLongitude                           *string  // Долгота нахождения плательщика
		Issuer                                *string  // Название банка-эмитента карты
		IssuerBankCountry                     *string  // Двухбуквенный код страны эмитента карты по ISO3166-1
		Description                           *string  // Назначение оплаты из параметров платежа
		AuthCode                              *string  // Код авторизации
		Data                                  *string  // Произвольный набор параметров, переданных в транзакцию
		Token                                 *string  // Токен карты для повторных платежей без ввода реквизитов
		TotalFee                              *float64 // Значение общей комиссии
		CardProduct                           *string  // Тип карточного продукта
		PaymentMethod                         *string  // Метод оплаты ApplePay, GooglePay, Яндекс Пэй или T-Pay
		FallBackScenarioDeclinedTransactionId *int64   // Номер первой неуспешной транзакции
		Rrn                                   *string  // Уникальный номер банковской транзакции, который назначается обслуживающим банком
	}

	// CloudPaymentsResponse represents check response for CloudPayments
	CloudPaymentsResponse struct {
		Code int8 `json:"code"`
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
		InvoiceID          uuid.UUID
		Provider           string
		OrderStateSnapshot string
		EventType          string
		Error              string
		ExtraInfo          []byte
		Response           []byte
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

	//orderStateConfirmed = "confirmed"

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
func (s *CloudPaymentsService) HandleCloudPaymentsCheckWebHook(ctx context.Context, checkRequest CloudPaymentsCheckRequest) (CloudPaymentsResponse, error) {
	res, mainErr := s.handleCloudPaymentsCheckWebHook(ctx, checkRequest)
	var errorText string
	if mainErr != nil {
		errorText = mainErr.Error()
	}

	extraInfoJSON, err := json.Marshal(checkRequest)
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	responseJSON, err := json.Marshal(res)
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	invoiceID, err := uuid.Parse(*checkRequest.InvoiceId)
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	// save the extra info
	err = s.repo.SavePaymentExtraInfo(ctx, SavePaymentExtraInfoParams{
		InvoiceID: invoiceID,
		Provider:  providerCloudPayments,
		EventType: eventTypeCheck,
		ExtraInfo: extraInfoJSON,
		Response:  responseJSON,
		Error:     errorText,
	})
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	return res, mainErr
}

func (s *CloudPaymentsService) handleCloudPaymentsCheckWebHook(ctx context.Context, checkRequest CloudPaymentsCheckRequest) (resp CloudPaymentsResponse, err error) {
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

	code := int8(checkPayment(checkRequest, order, s.maxDurationForHandlingPayment))
	return CloudPaymentsResponse{
		Code: code,
	}, nil
}

// HandleCloudPaymentsPayWebHook is the entry point for a webhook request from CloudPayments
func (s *CloudPaymentsService) HandleCloudPaymentsPayWebHook(ctx context.Context, payRequest CloudPaymentsPayRequest) (resp CloudPaymentsResponse, err error) {
	res, mainErr := s.handleCloudPaymentsPayWebHook(ctx, payRequest)
	var errorText string
	if mainErr != nil {
		errorText = mainErr.Error()
	}

	extraInfoJSON, err := json.Marshal(payRequest)
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	responseJSON, err := json.Marshal(res)
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	invoiceID, err := uuid.Parse(*payRequest.InvoiceId)
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	// save the extra info
	err = s.repo.SavePaymentExtraInfo(ctx, SavePaymentExtraInfoParams{
		InvoiceID: invoiceID,
		Provider:  providerCloudPayments,
		EventType: eventTypePay,
		ExtraInfo: extraInfoJSON,
		Response:  responseJSON,
		Error:     errorText,
	})
	if err != nil {
		return CloudPaymentsResponse{}, err
	}

	return res, err
}

func (s *CloudPaymentsService) handleCloudPaymentsPayWebHook(ctx context.Context, payRequest CloudPaymentsPayRequest) (resp CloudPaymentsResponse, err error) {
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
	return strings.EqualFold(cur1, cur2)
}

func isPaymentExpired(orderCreated time.Time, paymentWasMade time.Time, maxDuration time.Duration) bool {
	return paymentWasMade.Sub(orderCreated) > maxDuration
}
