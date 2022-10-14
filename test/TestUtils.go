package test

import "testing"

func AssertError(t *testing.T, e error) {
	if e == nil {
		t.Fatal("Expected error but wasn't error")
	}
}

func AssertNoError(t *testing.T, e error) {
	if e != nil {
		t.Fatal("Expected no error but was error")
	}
}

func AssertTrue(t *testing.T, b bool) {
	if !b {
		t.Fatal("Expected true but was false")
	}
}

func AssertFalse(t *testing.T, b bool) {
	if b {
		t.Fatal("Expected false but was true")
	}
}

func AssertSameInt(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Fatalf("Actual value '%d' does not equal expected value '%d", actual, expected)
	}
}

func AssertSameString(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Fatalf("Actual value '%s' does not equal expected value '%s", actual, expected)
	}
}
