package account

import (
	"testing"
)

func TestGetSidTest(t *testing.T) {
	cookieStr := "bGlfbWluZz0zMTU1MjEz"
	expected := "li_ming=3155213"

	actualSid, err := GetSidTest(cookieStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if actualSid != expected {
		t.Errorf("Expected sid: %s, but got: %s", expected, actualSid)
	}

	cookieStr = "bGlfaHVhPTkyMzE0NDY="
	expected = "li_hua=9231446"

	actualSid, err = GetSidTest(cookieStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if actualSid != expected {
		t.Errorf("Expected sid: %s, but got: %s", expected, actualSid)
	}
}
