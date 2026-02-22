# Tester package

The `tester` package provides small, convenient helpers for writing Go tests: a lightweight assertion helper (`Assert`) and a simple timing helper (`Wait`). The helpers are designed to keep tests readable and reduce boilerplate.

## Key types and functions

- `Assert`: assertion helper that wraps `testing.T` and can either report errors or fail fatally.
    - `NewAssert(t *testing.T, fatal bool) *Assert` — create a new `Assert`. When `fatal` is true, failing assertions call `t.Fatal`; otherwise they call `t.Error`.
    - Methods:
    - `IsErrNil(err error, message string)` — asserts `err == nil`.
    - `IsTrue(condition bool, message string)` — asserts a boolean condition is true.
    - `AreEqual(expected, actual interface{}, message string)` — asserts equality using `!=`.
    - `IsNotNil(object interface{}, message string)` — asserts value is non-nil.
    - `IsNotEmpty(s string, message string)` — asserts a string is not empty.

- `Wait`: a tiny helper for pausing test execution.
    - `NewWait() *Wait` — create a new `Wait` instance.
    - `WaitForMilliSeconds(millis int)` — sleep for the specified number of milliseconds.

## Usage

Import the package (module path):

    import "github.com/kinakoman/go-extensions/tester"

Example — using `Assert` in a test:

    package mypkg_test

    import (
    	"testing"
    	"github.com/kinakoman/go-extensions/tester"
    )

    func TestSomething(t *testing.T) {
    	a := tester.NewAssert(t, true) // fatal on failure

    	var err error = nil
    	a.IsErrNil(err, "error should be nil")

    	a.IsTrue(2+2 == 4, "math broke")
    	a.AreEqual("foo", "foo", "strings match")
    }

Example — using `Wait` in a test:

    package mypkg_test

    import (
    	"testing"
    	"github.com/kinakoman/go-extensions/tester"
    	"time"
    )

    func TestWait(t *testing.T) {
    	w := tester.NewWait()
    	start := time.Now()
    	w.WaitForMilliSeconds(100)
    	if time.Since(start) < 100*time.Millisecond {
    		t.Error("did not wait long enough")
    	}
    }

## Installing

If you're using Go modules, add the import to your code and run `go get`:

    go get github.com/kinakoman/go-extensions/tester@latest

## Notes

- `Assert` uses `testing.T` and marks helpers with `t.Helper()` so failure lines point to the caller in your test code.
- `AreEqual` uses the Go `!=` operator and thus compares values according to Go's equality rules; for deep-equality of complex structures consider using `reflect.DeepEqual` or other comparison helper if needed.
- `Wait` is a minimal wrapper around `time.Sleep` and intended only for simple timing checks in tests.

## License

This project is **not under an open-source license**,  
but you are welcome to **use, modify, and share** the code for any purpose.  
Please keep the author credit somewhere in your project.  
No warranty or liability is provided.
