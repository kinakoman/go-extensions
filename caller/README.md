# Caller package

The `caller` package provides a small utility to capture information about a single stack frame (the caller) at runtime. It's useful for logging, diagnostics, or any situation where you want to know the file, line and function name of the code that invoked a helper.

## Key types and functions

- Frame: a struct containing fields about a stack frame:
    - `File` (string) – full path to the source file
    - `ShortFile` (string) – base filename only
    - `Line` (int) – source line number
    - `Func` (string) – full package path + function name
    - `ShortFunc` (string) – function name only
- `Here(skip int) Frame`: returns the `Frame` for the caller. `skip=0` returns the immediate caller of `Here`. Use larger values to walk further up the call stack.

## Usage

Import the package (module path):

    import "github.com/kinakoman/go-extensions/caller"

Example:

    package main

    import (
    	"fmt"
    	"github.com/kinakoman/go-extensions/caller"
    )

    func main() {
    	// 0 = immediate caller of Here (this main function)
    	f := caller.Here(0)
    	fmt.Println("Full:", f.Format())       // e.g. /path/to/main.go:25 main.go
    	fmt.Println("Short:", f.FormatShort()) // e.g. main.go:25 main
    }

Example output (format may vary depending on environment):

    /home/user/project/caller/main.go:25 main.go
    main.go:25 main
