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

func TestCloudPaymentCheckPayment(t *testing.T) {
	duration, _ := time.ParseDuration("24h")
	orderUpdateTime, _ := time.Parse(time.DateTime, "2024-07-29 12:35:00")

	tests := []struct {
		in   inData
		out  int
		name string
	}{
		{
			in: inData{
				check: CloudPaymentsCheckRequest{
					InvoiceId: func() *string {
						id := "05b42581-0773-46ad-99ff-5c96ca4ed1f2"
						return &id
					}(),
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
			out: cloudPaymentsResponseCodeSuccess,
		},

		{
			in: inData{
				check: CloudPaymentsCheckRequest{
					InvoiceId: func() *string {
						id := "15b42581-0773-46ad-99ff-5c96ca4ed1f2"
						return &id
					}(),
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
			out: cloudPaymentsCheckResponseCodeWrongInvoiceID,
		},

		{
			in: inData{
				check: CloudPaymentsCheckRequest{
					InvoiceId: func() *string {
						id := "05b42581-0773-46ad-99ff-5c96ca4ed1f2"
						return &id
					}(),
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
			out: cloudPaymentsCheckResponseCodeWrongSum,
		},

		{
			in: inData{
				check: CloudPaymentsCheckRequest{
					InvoiceId: func() *string {
						id := "05b42581-0773-46ad-99ff-5c96ca4ed1f2"
						return &id
					}(),
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
			out: cloudPaymentsCheckResponseCodeWrongSum,
		},

		{
			in: inData{
				check: CloudPaymentsCheckRequest{
					InvoiceId: func() *string {
						id := "05b42581-0773-46ad-99ff-5c96ca4ed1f2"
						return &id
					}(),
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
			out: cloudPaymentsCheckResponseCodeTransactionExpired,
		},
		{
			in: inData{
				check: CloudPaymentsCheckRequest{},
				orderInfo: Order{
					ID:        uuid.MustParse("05b42581-0773-46ad-99ff-5c96ca4ed1f2"),
					Sum:       1000,
					Currency:  "rub",
					UpdatedAt: orderUpdateTime,
				},
				paymentMaxDuration: duration,
			},
			out: cloudPaymentsCheckResponseCodeWrongInvoiceID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := checkPayment(tt.in.check, tt.in.orderInfo, tt.in.paymentMaxDuration)
			if out != tt.out {
				t.Errorf("got %d, want %d", out, tt.out)
			}
		})
	}
}
