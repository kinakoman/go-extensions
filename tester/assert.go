package tester

import (
	"fmt"
	"testing"
)

// Assert provides assertion methods for testing.
type Assert struct {
	t     *testing.T // The testing.T instance
	fatal bool       // Whether to fail fatally on assertion failure
}

// NewAssert creates a new Assert instance.
func NewAssert(t *testing.T, f bool) *Assert {
	return &Assert{t: t, fatal: f}
}

// errorWithFatal reports an error with file and line number information.
func (a *Assert) errorWithFatal(message string) {
	a.t.Helper()
	if a.fatal {
		a.t.Fatal(message)
		return
	}

	a.t.Error(message)
}

// IsErrNil asserts that the given error is nil.
func (a *Assert) IsErrNil(err error, message string) {
	a.t.Helper()
	if err != nil {
		a.errorWithFatal(message + ": expected nil error, but got " + err.Error())
	}
}

// IsTrue asserts that the given condition is true.
func (a *Assert) IsTrue(condition bool, message string) {
	a.t.Helper()
	if !condition {
		a.errorWithFatal(message + ": expected true, but got false")
	}
}

// AreEqual asserts that the expected and actual values are equal.
func (a *Assert) AreEqual(expected, actual interface{}, message string) {
	a.t.Helper()
	if expected != actual {
		a.errorWithFatal(fmt.Sprintf("%s: expected %v, but got %v", message, expected, actual))
	}
}

// IsNotNil asserts that the given object is not nil.
func (a *Assert) IsNotNil(object interface{}, message string) {
	a.t.Helper()
	if object == nil {
		a.errorWithFatal(message + ": expected non-nil value, but got nil")
	}
}

// IsNotEmpty asserts that the given string is not empty.
func (a *Assert) IsNotEmpty(s string, message string) {
	a.t.Helper()
	if s == "" {
		a.errorWithFatal(message + ": expected non-empty string, but got empty string")
	}
}
