package cast

import (
	"errors"
	"math"
	"strconv"
)

var (
	ErrOverflow    = errors.New("overflow occurred")
	ErrStringParse = errors.New("string parsing error")
)

// IntToInt32 converts an int to int32, checking for overflow.
func IntToInt32(i int) (int32, error) {
	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0, ErrOverflow
	}
	return int32(i), nil
}

// Int64ToInt32 converts an int64 to int32, checking for overflow.
func Int64ToInt32(i int64) (int32, error) {
	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0, ErrOverflow
	}
	return int32(i), nil
}

// StringToInt32 converts a string to int32, checking for parsing errors and overflow.
func StringToInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, ErrStringParse
	}
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0, ErrOverflow
	}
	return int32(v), nil
}

// StringToInt64 converts a string to int64, checking for parsing errors and overflow.
func StringToInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, ErrStringParse
	}
	return v, nil
}

// Int32ToString converts an int32 to a string.
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// Int64ToString converts an int64 to a string.
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
