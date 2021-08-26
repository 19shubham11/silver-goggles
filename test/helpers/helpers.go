package helpers

import (
	"errors"
	"strings"
	"testing"
)

func AssertError(t *testing.T, expected, got error) {
	t.Helper()

	if !errors.Is(got, expected) {
		t.Fatalf("Expected error %v got %v", expected, got)
	}
}

// AssertSubstrings takes in a string and an array of expected sub-strings
func AssertSubstrings(t *testing.T, output string, expectedStrings []string) {
	t.Helper()

	for _, str := range expectedStrings {
		if !strings.Contains(output, str) {
			t.Errorf("Expected output to contain %s", str)
		}
	}
}
