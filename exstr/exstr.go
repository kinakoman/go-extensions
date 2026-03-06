package exstr

import (
	"regexp"
	"strings"
)

// Exstr is an extended string type that provides additional methods for string manipulation.
type Exstr string

// New creates a new Exstr from a regular string.
func New(s string) Exstr {
	return Exstr(s)
}

// String returns the underlying string value of Exstr.
func (s Exstr) String() string {
	return string(s)
}

// Contains checks if Exstr contains substr.
func (s Exstr) Contains(substr string) bool {
	return strings.Contains(string(s), substr)
}

// ContainsAny checks if Exstr contains any of the characters in chars.
func (s Exstr) ContainsAny(chars string) bool {
	return strings.ContainsAny(string(s), chars)
}

// Match checks if Exstr matches the given regular expression pattern.
func (s Exstr) Match(pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(string(s))
}

// Find returns the first substring of Exstr that matches the given regular expression pattern.
func (s Exstr) Find(pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.FindString(string(s))
}

// FindAll returns all substrings of Exstr that match the given regular expression pattern.
func (s Exstr) FindAll(pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(string(s), -1)
}

// Replace replaces all occurrences of old with new in Exstr.
func (s Exstr) Replace(old, new string) Exstr {
	return Exstr(strings.ReplaceAll(string(s), old, new))
}

// Split splits Exstr into a slice of substrings separated by the given separator.
func (s Exstr) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

// Trim removes all leading and trailing whitespace from Exstr.
func (s Exstr) Trim() Exstr {
	return Exstr(strings.TrimSpace(string(s)))
}
