package notifications

import (
	"github.com/shoppigram-com/marketplace-api/internal/notifications/templates/en"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/templates/ru"
	"time"
)

func getTranslation(lang, key string) string {
	switch lang {
	case langRu:
		return ru.Translations[key]
	case langEn:
		return en.Translations[key]
	default:
		return ru.Translations[key]
	}
}

func isLanguageValid(lang string) bool {
	for _, v := range validLangCodes {
		if lang == v {
			return true
		}
	}
	return false
}

func checkAndGetLangCode(lang string) string {
	if isLanguageValid(lang) {
		return lang
	}
	return fallbackLanguage
}

func formatRussianTime(t time.Time) string {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return ""
	}
	t = t.In(loc)
	return t.Format("02.01.2006 15:04:05")
}
