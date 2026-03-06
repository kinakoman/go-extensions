package exstr

import (
	"regexp"
	"strings"
)

// ExStr is an extended string type that provides additional methods for string manipulation.
type ExStr string

// New creates a new ExStr from a regular string.
func New(s string) ExStr {
	return ExStr(s)
}

// String returns the underlying string value of ExStr.
func (s ExStr) String() string {
	return string(s)
}

// Contains checks if ExStr contains substr.
func (s ExStr) Contains(substr string) bool {
	return strings.Contains(string(s), substr)
}

// Match checks if ExStr matches the given regular expression pattern.
func (s ExStr) Match(pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(string(s))
}

// Find returns the first substring of ExStr that matches the given regular expression pattern.
func (s ExStr) Find(pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.FindString(string(s))
}

// FindAll returns all substrings of ExStr that match the given regular expression pattern.
func (s ExStr) FindAll(pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(string(s), -1)
}

// Replace replaces all occurrences of old with new in ExStr.
func (s ExStr) Replace(old, new string) ExStr {
	return ExStr(strings.ReplaceAll(string(s), old, new))
}

// Split splits ExStr into a slice of substrings separated by the given separator.
func (s ExStr) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

// Trim removes all leading and trailing whitespace from ExStr.
func (s ExStr) Trim() ExStr {
	return ExStr(strings.TrimSpace(string(s)))
}
