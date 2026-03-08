package exstr

import "testing"

func TestExStr(t *testing.T) {
	str := New("Hello, World!")

	if !str.Contains("World") {
		t.Errorf("Expected 'Hello, World!' to contain 'World'")
	}

	if !str.Match(`^Hello`) {
		t.Errorf("Expected 'Hello, World!' to match '^Hello'")
	}

	if str.Find(`\w+`) != "Hello" {
		t.Errorf("Expected Find to return 'Hello', got '%s'", str.Find(`\w+`))
	}

	if len(str.FindAll(`\w+`)) != 2 {
		t.Errorf("Expected FindAll to return 2 matches, got %d", len(str.FindAll(`\w+`)))
	}

	replaced := str.Replace("World", "Go")
	if replaced.String() != "Hello, Go!" {
		t.Errorf("Expected Replace to return 'Hello, Go!', got '%s'", replaced.String())
	}

	split := str.Split(", ")
	if len(split) != 2 || split[0] != "Hello" || split[1] != "World!" {
		t.Errorf("Expected Split to return ['Hello', 'World!'], got %v", split)
	}

	trimmed := New("  Hello, World!  ").Trim()
	if trimmed.String() != "Hello, World!" {
		t.Errorf("Expected Trim to return 'Hello, World!', got '%s'", trimmed.String())
	}
}

func TestExStrSlice(t *testing.T) {
	str := New("Hello,World")
	found := str.Split(",").Replace("H", "h").Contains("hello")
	if !found {
		t.Errorf("Expected ExstrSlice to contain 'hello', got %v", found)
	}
}
