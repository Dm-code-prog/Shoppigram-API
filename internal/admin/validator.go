package admin

func isShortNameValid(shortName string) bool {
	return shortNameRegex.MatchString(shortName)
}

func isNameValid(name string) bool {
	return len(name) >= 3
}

func isProductNameValid(name string) bool {
	return len(name) >= 3 && len(name) <= 75
}

func isValidImageExtension(ext string) bool {
	switch ext {
	case "png", "jpg", "jpeg", "webp":
		return true
	}

	return false
}
