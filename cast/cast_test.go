package cast

import (
	"math"
	"testing"
)

func TestIntToInt32(t *testing.T) {
	tests := []struct {
		input    int
		expected int32
		wantErr  error
	}{
		{0, 0, nil},
		{math.MaxInt32, math.MaxInt32, nil},
		{math.MinInt32, math.MinInt32, nil},
		{math.MaxInt32 + 1, 0, ErrOverflow},
		{math.MinInt32 - 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		got, err := IntToInt32(tt.input)
		if got != tt.expected {
			t.Errorf("IntToInt32(%d) got %d, expected %d", tt.input, got, tt.expected)
		}
		if err != tt.wantErr {
			t.Errorf("IntToInt32(%d) error %v, expected %v", tt.input, err, tt.wantErr)
		}
	}
}

func TestInt64ToInt32(t *testing.T) {
	tests := []struct {
		input    int64
		expected int32
		wantErr  error
	}{
		{0, 0, nil},
		{math.MaxInt32, math.MaxInt32, nil},
		{math.MinInt32, math.MinInt32, nil},
		{math.MaxInt32 + 1, 0, ErrOverflow},
		{math.MinInt32 - 1, 0, ErrOverflow},
	}

	for _, tt := range tests {
		got, err := Int64ToInt32(tt.input)
		if got != tt.expected {
			t.Errorf("Int64ToInt32(%d) got %d, expected %d", tt.input, got, tt.expected)
		}
		if err != tt.wantErr {
			t.Errorf("Int64ToInt32(%d) error %v, expected %v", tt.input, err, tt.wantErr)
		}
	}
}

func TestStringToInt32(t *testing.T) {
	tests := []struct {
		input    string
		expected int32
		wantErr  error
	}{
		{"0", 0, nil},
		{"2147483647", math.MaxInt32, nil},
		{"-2147483648", math.MinInt32, nil},
		{"2147483648", 0, ErrStringParse}, // strconv.ParseInt with bitSize 32 returns error on overflow
		{"invalid", 0, ErrStringParse},
	}

	for _, tt := range tests {
		got, err := StringToInt32(tt.input)
		if got != tt.expected {
			t.Errorf("StringToInt32(%s) got %d, expected %d", tt.input, got, tt.expected)
		}
		if err != tt.wantErr {
			t.Errorf("StringToInt32(%s) error %v, expected %v", tt.input, err, tt.wantErr)
		}
	}
}

func TestStringToInt64(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		wantErr  error
	}{
		{"0", 0, nil},
		{"9223372036854775807", math.MaxInt64, nil},
		{"-9223372036854775808", math.MinInt64, nil},
		{"invalid", 0, ErrStringParse},
	}

	for _, tt := range tests {
		got, err := StringToInt64(tt.input)
		if got != tt.expected {
			t.Errorf("StringToInt64(%s) got %d, expected %d", tt.input, got, tt.expected)
		}
		if err != tt.wantErr {
			t.Errorf("StringToInt64(%s) error %v, expected %v", tt.input, err, tt.wantErr)
		}
	}
}
