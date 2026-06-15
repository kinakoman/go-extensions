package exchan

import (
	"context"
	"testing"
)

func TestExchan(t *testing.T) {
	e := New[int](1)
	ctx := context.Background()

	// Test sending and receiving
	if err := e.Send(ctx, 42); err != nil {
		t.Fatalf("Send failed: %v", err)
	}

	v, err := e.Recv(ctx)
	if err != nil {
		t.Fatalf("Recv failed: %v", err)
	}
	if v != 42 {
		t.Fatalf("Expected 42, got %d", v)
	}

	// Test closing
	e.Close()
	if err := e.Send(ctx, 43); err != ErrChannelClosed {
		t.Fatalf("Expected ErrChannelClosed, got %v", err)
	}
	if _, err := e.Recv(ctx); err != ErrChannelClosed {
		t.Fatalf("Expected ErrChannelClosed, got %v", err)
	}
}

func TestExchanContext(t *testing.T) {
	e := New[int](1)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Test sending with canceled context
	if err := e.Send(ctx, 42); err != context.Canceled {
		t.Fatalf("Expected context.Canceled, got %v", err)
	}

	// Test receiving with canceled context
	if _, err := e.Recv(ctx); err != context.Canceled {
		t.Fatalf("Expected context.Canceled, got %v", err)
	}
}
