# Cast package

The `cast` package provides utilities for safe type casting and conversion, handling overflows and parsing errors gracefully. It is designed to ensure that data conversions, especially between integer types and from strings, are performed safely without unexpected behavior or panics.

## Key functions

- `IntToInt32(i int) (int32, error)`: converts an `int` to `int32`, checking for overflow.
- `Int64ToInt32(i int64) (int32, error)`: converts an `int64` to `int32`, checking for overflow.
- `StringToInt32(s string) (int32, error)`: converts a `string` to `int32`, handling parsing errors and overflow.
- `StringToInt64(s string) (int64, error)`: converts a `string` to `int64`, handling parsing errors.

## Usage

Import the package:

    import "github.com/kinakoman/go-extensions/cast"

Example:

    package main

    import (
    	"fmt"
    	"github.com/kinakoman/go-extensions/cast"
    )

    func main() {
    	val, err := cast.StringToInt32("12345")
    	if err != nil {
    		fmt.Println("Error:", err)
    		return
    	}
    	fmt.Printf("Converted value: %d\n", val)

    	// Overflow example
    	_, err = cast.IntToInt32(2147483648) // MaxInt32 + 1
    	if err != nil {
    		fmt.Println("Overflow error:", err)
    	}
    }
