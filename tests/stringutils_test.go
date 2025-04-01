package tests

import (
	"go_utilities_toolkit/utils/stringutils"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	valid := "test@example.com"
	invalid := "invalid-email"
	if !stringutils.IsValidEmail(valid) {
		t.Errorf("Expected true, got false for valid email")
	}
	if stringutils.IsValidEmail(invalid) {
		t.Errorf("Expected false, got true for invalid email")
	}
}

func TestIsValidPhone(t *testing.T) {
	valid := "9876543210"
	invalid := "1234"
	if !stringutils.IsValidPhone(valid) {
		t.Errorf("Expected valid phone number")
	}
	if stringutils.IsValidPhone(invalid) {
		t.Errorf("Expected invalid phone number")
	}
}

func TestSlugify(t *testing.T) {
	result := stringutils.Slugify("Hello World!")
	expected := "hello-world"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestTruncate(t *testing.T) {
	text := "Hello, this is a long sentence."
	result := stringutils.Truncate(text, 10)
	if result != "Hello, thi..." {
		t.Errorf("Unexpected truncate result: %s", result)
	}
}
