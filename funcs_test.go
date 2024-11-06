package blockeditor

import "testing"

func Test_TrimPrefixedKey(t *testing.T) {
	prefix := "abc_"
	key := prefix + prefix + "key"

	key = trimPrefixedKey(key, prefix)

	if key != "key" {
		t.Error("Expected key to be trimmed, but found:", key)
	}
}

func Test_PrefixKey(t *testing.T) {
	prefix := "abc_"
	key := prefix + prefix + "key"

	key = prefixKey(key, prefix)

	if key != prefix+"key" {
		t.Error("Expected key to be prefixed, but found:", key)
	}
}

func Test_UnprefixKey(t *testing.T) {
	prefix := "abc_"
	key := prefix + prefix + "key"

	key = unprefixKey(key, prefix)

	if key != "key" {
		t.Error("Expected key to be decoded, but found:", key)
	}
}
