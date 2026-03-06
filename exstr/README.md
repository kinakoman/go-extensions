# Exstr package

The `exstr` package provides an extended string type (`Exstr`) that wraps Go's built-in `string` with convenient methods for common string operations such as pattern matching, replacement, splitting, and trimming.

## Key types and functions

- `Exstr`: Extended string type based on `string`.
- `New(s string) Exstr`: Creates a new `Exstr` from a regular string.
- `String() string`: Returns the underlying string value.
- `Contains(substr string) bool`: Checks if the string contains the given substring.
- `Match(pattern string) bool`: Checks if the string matches the given regular expression pattern.
- `Find(pattern string) string`: Returns the first substring matching the given regular expression pattern.
- `FindAll(pattern string) []string`: Returns all substrings matching the given regular expression pattern.
- `Replace(old, new string) Exstr`: Replaces all occurrences of `old` with `new`.
- `Split(sep string) []string`: Splits the string by the given separator.
- `Trim() Exstr`: Removes all leading and trailing whitespace.

## Usage

Import the package:

    import "github.com/kinakoman/go-extensions/exstr"

Example:

    package main

    import (
    	"fmt"
    	"github.com/kinakoman/go-extensions/exstr"
    )

    func main() {
    	s := exstr.New("  Hello, World!  ")

    	// Trim whitespace
    	trimmed := s.Trim()
    	fmt.Println(trimmed) // "Hello, World!"

    	// Check containment
    	fmt.Println(trimmed.Contains("World")) // true

    	// Regex match
    	fmt.Println(trimmed.Match(`^Hello`)) // true

    	// Find first match
    	fmt.Println(trimmed.Find(`\w+`)) // "Hello"

    	// Find all matches
    	fmt.Println(trimmed.FindAll(`\w+`)) // ["Hello" "World"]

    	// Replace
    	fmt.Println(trimmed.Replace("World", "Go")) // "Hello, Go!"

    	// Split
    	fmt.Println(trimmed.Split(", ")) // ["Hello" "World!"]
    }
