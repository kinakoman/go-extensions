package tester

import "time"

// Wait provides methods to pause execution for a specified duration.
type Wait struct{}

// NewWait creates a new Wait instance.
func NewWait() *Wait {
	return &Wait{}
}

// WaitForMilliSeconds pauses execution for the specified number of milli seconds.
func (w *Wait) WaitForMilliSeconds(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}
