package notifications

import (
	"strconv"
	"strings"
)

func formatFloat(num float64) string {
	str := strconv.FormatFloat(num, 'f', -1, 64)
	parts := strings.Split(str, ".")
	intPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	n := len(intPart)
	if n <= 3 {
		return intPart + decimalPart
	}

	var result string
	for i := 0; i < n; i++ {
		result = string(intPart[n-1-i]) + result
		if (i+1)%3 == 0 && i != n-1 {
			result = "," + result
		}
	}
	return result + decimalPart
}

func formatCurrency(currency string) string {
	currency = strings.ToLower(currency)
	switch currency {
	case "usd":
		return "$"
	case "eur":
		return "€"
	case "rub":
		return "₽"
	default:
		return currency
	}
}
