package exstr

import (
	"regexp"
	"strings"
)

// Exstr is an extended string type that provides additional methods for string manipulation.
type Exstr string

// ExstrSlice is a slice of Exstr, providing methods to work with multiple Exstr values.
type ExstrSlice []Exstr

// New creates a new Exstr from a regular string.
func New(s string) Exstr {
	return Exstr(s)
}

// String returns the underlying string value of Exstr.
func (s Exstr) String() string {
	return string(s)
}

// Strings returns a slice of regular strings from an ExstrSlice.
func (s ExstrSlice) Strings() []string {
	strs := make([]string, len(s))
	for i, exstr := range s {
		strs[i] = exstr.String()
	}
	return strs
}

// Contains checks if Exstr contains substr.
func (s Exstr) Contains(substr string) bool {
	if substr == "" {
		return false
	}

	return strings.Contains(string(s), substr)
}

// Contains checks if any Exstr in the slice contains substr.
func (s ExstrSlice) Contains(substr string) bool {
	if substr == "" {
		return false
	}

	for _, exstr := range s {
		if exstr.Contains(substr) {
			return true
		}
	}
	return false
}

// ContainsAny checks if Exstr contains any of the characters in chars.
func (s Exstr) ContainsAny(chars string) bool {
	if chars == "" {
		return false
	}

	return strings.ContainsAny(string(s), chars)
}

// Match checks if Exstr matches the given regular expression pattern.
func (s Exstr) Match(pattern string) bool {
	if pattern == "" {
		return false
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(string(s))
}

// Match checks if any Exstr in the slice matches the given regular expression pattern.
func (s ExstrSlice) Match(pattern string) bool {
	for _, exstr := range s {
		if exstr.Match(pattern) {
			return true
		}
	}
	return false
}

// Find returns the first substring of Exstr that matches the given regular expression pattern.
func (s Exstr) Find(pattern string) Exstr {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}
	return Exstr(re.FindString(string(s)))
}

// Find returns the first substring of any Exstr in the slice that matches the given regular expression pattern.
func (s ExstrSlice) Find(pattern string) ExstrSlice {
	results := make([]Exstr, len(s))
	for i, exstr := range s {
		if match := exstr.Find(pattern); match != "" {
			results[i] = match
		}
	}
	return results
}

// FindAll returns all substrings of Exstr that match the given regular expression pattern.
func (s Exstr) FindAll(pattern string) ExstrSlice {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return ExstrSlice{}
	}

	matches := re.FindAllString(string(s), -1)

	slice := make(ExstrSlice, len(matches))
	for i, match := range matches {
		slice[i] = Exstr(match)
	}
	return slice
}

// Replace replaces all occurrences of old with new in Exstr.
func (s Exstr) Replace(old, new string) Exstr {
	return Exstr(strings.ReplaceAll(string(s), old, new))
}

// Replace replaces all occurrences of old with new in each Exstr in the slice.
func (s ExstrSlice) Replace(old, new string) ExstrSlice {
	slice := make(ExstrSlice, 0, len(s))
	for _, exstr := range s {
		slice = append(slice, exstr.Replace(old, new))
	}
	return slice
}

// Split splits Exstr into a slice of substrings separated by the given separator.
func (s Exstr) Split(sep string) ExstrSlice {
	parts := strings.Split(string(s), sep)
	slice := make(ExstrSlice, len(parts))
	for i, part := range parts {
		slice[i] = Exstr(part)
	}
	return slice
}

// Trim removes all leading and trailing whitespace from Exstr.
func (s Exstr) Trim() Exstr {
	return Exstr(strings.TrimSpace(string(s)))
}

// Trim removes all leading and trailing whitespace from each Exstr in the slice.
func (s ExstrSlice) Trim() ExstrSlice {
	slice := make(ExstrSlice, len(s))
	for i, exstr := range s {
		slice[i] = exstr.Trim()
	}
	return slice
}
