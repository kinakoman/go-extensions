package exchan

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrChannelClosed = errors.New("exchan: channel is closed")
)

// Exchan is an extended channel wrapper that never panics.
// It does NOT close the underlying data channel. Close() only closes done.
type Exchan[T any] struct {
	c    chan T
	done chan struct{}
	once sync.Once
}

func New[T any](buffer int) *Exchan[T] {
	return &Exchan[T]{
		c:    make(chan T, buffer),
		done: make(chan struct{}),
	}
}

// Close stops the Exchan. It is safe to call multiple times.
func (e *Exchan[T]) Close() {
	e.once.Do(func() {
		close(e.done)
	})
}

// Done returns a channel that is closed when Exchan is closed.
func (e *Exchan[T]) Done() <-chan struct{} { return e.done }

// Send sends v, blocking until it succeeds, ctx is done, or Exchan is closed.
func (e *Exchan[T]) Send(ctx context.Context, v T) error {
	select {
	case <-e.done:
		return ErrChannelClosed
	case <-ctx.Done():
		return ctx.Err()
	case e.c <- v:
		return nil
	}
}

// Recv receives a value, blocking until it succeeds, ctx is done, or Exchan is closed.
func (e *Exchan[T]) Recv(ctx context.Context) (T, error) {
	select {
	case <-e.done:
		var zero T
		return zero, ErrChannelClosed
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	case v := <-e.c:
		return v, nil
	}
}

func (e *Exchan[T]) Len() int { return len(e.c) }
func (e *Exchan[T]) Cap() int { return cap(e.c) }
