package blockeditor

import "strings"

func decodeSettingKey(key string) string {
	key = trimSettingKey(key)
	return key
}

// emcodeSettingKey prepares the setting key for the moodal form
// to not conflict with any other field names
func encodeSettingKey(key string) string {
	key = trimSettingKey(key)
	return SETTINGS_PREFIX + key
}

func isSettingKey(key string) bool {
	return strings.HasPrefix(key, SETTINGS_PREFIX)
}

func trimSettingKey(key string) string {
	for strings.HasPrefix(key, SETTINGS_PREFIX) {
		key = strings.TrimPrefix(key, SETTINGS_PREFIX)
	}
	return key
}
