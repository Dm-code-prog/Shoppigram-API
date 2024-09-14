package notifications

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
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

func formatRussianTime(t time.Time) string {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return ""
	}
	t = t.In(loc)
	return t.Format("02.01.2006 15:04:05")
}

func isLanguageValid(lang string) bool {
	for _, v := range validLangCodes {
		if lang == v {
			return true
		}
	}
	return false
}

func getButtonText(lang string, key string) (string, error) {
	var bt map[string]string
	data, err := templates.ReadFile("templates/" + lang + pathToButtonsText)
	if err != nil {
		return "", errors.Wrap(err, "templates.ReadFile("+lang+".json)")
	}
	err = json.Unmarshal(data, &bt)
	if err != nil {
		return "", errors.Wrap(err, "json.Unmarshal(data, bt)")
	}
	return bt[key], nil
}
