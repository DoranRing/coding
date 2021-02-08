package datastruct

import (
	"testing"
)

func TestNewSeparateChainHash(t *testing.T) {
	hash := NewSeparateChainHash(128)
	hash.Put("one", "one value")
	hash.Put("two", "two value")
	if !hash.Contains("one") {
		t.Errorf("should: %t, actual: %t", true, false)
	}
	if hash.Contains("three") {
		t.Errorf("should: %t, actual: %t", false, true)
	}

	twoVal, ok := hash.Get("two")
	if !ok {
		t.Errorf("should: %t, actual: %t", true, false)
	}
	if twoVal != "two value" {
		t.Errorf("should: %s, actual: %s", "two value", twoVal)
	}

	hash.Delete("two")
	_, ok = hash.Get("two")
	if ok {
		t.Errorf("should: %t, actual: %t", false, true)
	}
}
