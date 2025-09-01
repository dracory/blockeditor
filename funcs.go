package blockeditor

import "strings"

func unprefixKey(key, prefix string) string {
	key = trimPrefixedKey(key, prefix)
	return key
}

// emcodeSettingKey prepares the setting key for the moodal form
// to not conflict with any other field names
func prefixKey(key, prefix string) string {
	key = trimPrefixedKey(key, prefix)
	return prefix + key
}

func isPrefixedKey(key, prefix string) bool {
	return strings.HasPrefix(key, prefix)
}

// trimSettingKey removes the settings prefix
// repeatedly until the key is no longer prefixed
func trimPrefixedKey(key, prefix string) string {
	for strings.HasPrefix(key, prefix) {
		key = strings.TrimPrefix(key, prefix)
	}
	return key
}

// isJSON is naive implementation for superficial, rough and fast checking for JSON
func isJSON(str string) bool {
	if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		return true
	}

	if strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]") {
		return true
	}

	return false
}
