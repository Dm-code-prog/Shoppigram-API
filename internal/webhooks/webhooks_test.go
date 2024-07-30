package webhooks

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

type inData struct {
	check              CloudPaymentsCheckRequest
	orderInfo          Order
	paymentMaxDuration time.Duration
}

var (
	duration, _        = time.ParseDuration("24h")
	orderUpdateTime, _ = time.Parse(time.DateTime, "2024-07-29 12:35:00")

	CloudPaymentsCheckPaymentTests = map[string]struct {
		in  inData
		out int
	}{
		"valid payment": {
			inData{
				check: CloudPaymentsCheckRequest{
					InvoiceID:       "05b42581-0773-46ad-99ff-5c96ca4ed1f2",
					Amount:          1000.00,
					Currency:        "RUB",
					PaymentAmount:   "1000.00",
					PaymentCurrency: "RUB",
					DateTime:        "2024-07-29 12:35:00",
				},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			cloudPaymentsCheckResponseCodeSuccess,
		},

		"wrong InvoiceID": {
			inData{
				check: CloudPaymentsCheckRequest{
					InvoiceID:       "15b42581-0773-46ad-99ff-5c96ca4ed1f2",
					Amount:          1000.00,
					Currency:        "RUB",
					PaymentAmount:   "1000.00",
					PaymentCurrency: "RUB",
					DateTime:        "2024-07-29 12:35:00",
				},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			cloudPaymentsCheckResponseCodeWrongInvoiceID,
		},

		"wrong payment amount": {
			inData{
				check: CloudPaymentsCheckRequest{
					InvoiceID:       "05b42581-0773-46ad-99ff-5c96ca4ed1f2",
					Amount:          500.00,
					Currency:        "RUB",
					PaymentAmount:   "500.00",
					PaymentCurrency: "RUB",
					DateTime:        "2024-07-29 12:35:00",
				},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			cloudPaymentsCheckResponseCodeWrongSum,
		},

		"wrong currency": {
			inData{
				check: CloudPaymentsCheckRequest{
					InvoiceID:       "05b42581-0773-46ad-99ff-5c96ca4ed1f2",
					Amount:          1000.00,
					Currency:        "USD",
					PaymentAmount:   "1000.00",
					PaymentCurrency: "USD",
					DateTime:        "2024-07-29 12:35:00",
				},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			cloudPaymentsCheckResponseCodeWrongSum,
		},

		"payment expired": {
			inData{
				check: CloudPaymentsCheckRequest{
					InvoiceID:       "05b42581-0773-46ad-99ff-5c96ca4ed1f2",
					Amount:          1000.00,
					Currency:        "RUB",
					PaymentAmount:   "1000.00",
					PaymentCurrency: "RUB",
					DateTime:        "2024-08-01 12:35:00",
				},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			cloudPaymentsCheckResponseCodeTransactionExpired,
		},
		"empty request": {
			inData{
				check: CloudPaymentsCheckRequest{},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			cloudPaymentsCheckResponseCodeWrongInvoiceID,
		},
	}
)

func TestCloudPaymentCheckPayment(t *testing.T) {
	for name, test := range CloudPaymentsCheckPaymentTests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			out := checkPayment(test.in.check, test.in.orderInfo, test.in.paymentMaxDuration)
			if out != test.out {
				t.Errorf("got %d, want %d", out, test.out)
			}
		})
	}
}
