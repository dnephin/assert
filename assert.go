package assert

import (
	"reflect"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// TestingT is an interface which defines the methods of testing.T that are
// required by this package
type TestingT interface {
	Fatalf(string, ...interface{})
}

// Equal compare the actual value to the expected value and fails the test if
// they are not equal.
func Equal(t TestingT, actual, expected interface{}) {
	if expected != actual {
		t.Fatalf("Expected '%v' (%T) got '%v' (%T)", expected, expected, actual, actual)
	}
}

// DeepEqual uses reflect.DeepEqual to compare the actual value to the expected
// value and fails the test if they are not equal.
func DeepEqual(t TestingT, actual, expected interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(
			"\nExpected:\n%s\n\nActual:\n%s\n",
			spew.Sdump(expected),
			spew.Sdump(actual))
	}
}

// Contains asserts that the string contains a substring, otherwise it fails the
// test.
func Contains(t TestingT, actual, contains string) {
	if !strings.Contains(actual, contains) {
		t.Fatalf("Expected '%s' to contain '%s'", actual, contains)
	}
}

// NilError asserts that the error is nil, otherwise it fails the test.
func NilError(t TestingT, err error) {
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err.Error())
	}
}

// Error asserts that error is not nil, and contains the expected text,
// otherwise it fails the test.
func Error(t TestingT, err error, contains string) {
	if err == nil {
		t.Fatalf("Expected an error, but error was nil")
	}

	if !strings.Contains(err.Error(), contains) {
		t.Fatalf("Expected error to contain '%s', got '%s'", contains, err.Error())
	}
}

// ErrorEquals asserts that error is not nil, and the text equals expected,
// otherwise it fails the test.
func ErrorEquals(t TestingT, err error, expected string) {
	if err == nil {
		t.Fatalf("Expected an error, but error was nil")
	}

	if err.Error() != expected {
		t.Fatalf("Expected error to be '%s', got '%s'", expected, err.Error())
	}
}
