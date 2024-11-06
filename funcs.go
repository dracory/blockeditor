package blockeditor

import "strings"

func unprefixKey(key string, prefix string) string {
	key = trimPrefixedKey(key, prefix)
	return key
}

// emcodeSettingKey prepares the setting key for the moodal form
// to not conflict with any other field names
func prefixKey(key string, prefix string) string {
	key = trimPrefixedKey(key, prefix)
	return prefix + key
}

func isPrefixedKey(key string, prefix string) bool {
	return strings.HasPrefix(key, prefix)
}

// trimSettingKey removes the settings prefix
// repeatedly until the key is no longer prefixed
func trimPrefixedKey(key string, prefix string) string {
	for strings.HasPrefix(key, prefix) {
		key = strings.TrimPrefix(key, prefix)
	}
	return key
}
