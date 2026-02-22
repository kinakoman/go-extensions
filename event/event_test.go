package event

import (
	"context"
	"errors"
	"testing"
	"time"
)

func recvOrTimeout[T any](ch <-chan T, d time.Duration) (T, bool) {
	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(d):
		var zero T
		return zero, false
	}
}

func expectPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic but none occurred")
		}
	}()
	f()
}

// ============================================================
// panic if Subscribe before Start
// ============================================================
func TestSubscribeBeforeStartPanics(t *testing.T) {
	eb := NewEventBus[int]()
	expectPanic(t, func() {
		eb.Subscribe(context.Background(), 1)
	})
}

// ============================================================
// panic if UnSubscribe before Start
// ============================================================
func TestUnsubscribeBeforeStartPanics(t *testing.T) {
	eb := NewEventBus[int]()
	ch := make(chan int)
	expectPanic(t, func() {
		eb.Unsubscribe(ch)
	})
}

// ============================================================
// panic if Publish before Start
// ============================================================
func TestPublishBeforeStartPanics(t *testing.T) {
	eb := NewEventBus[int]()
	expectPanic(t, func() {
		eb.Publish(10)
	})
}

// ============================================================
// Basic publish and receive
// ============================================================
func TestEventBus_PublishReceive(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()
	defer eb.Stop()

	ctx := context.Background()
	sub := eb.Subscribe(ctx, 1)

	eb.Publish(123)

	v, ok := recvOrTimeout(sub, time.Second)
	if !ok || v != 123 {
		t.Fatalf("expected 123, got %v (ok=%v)", v, ok)
	}
}

// ============================================================
// slow subscriber should be dropped
// ============================================================
func TestEventBus_SlowSubscriberDropped(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()
	defer eb.Stop()

	// wait until the event bus is running
	time.Sleep(10 * time.Millisecond)

	ctx := context.Background()
	sub := eb.Subscribe(ctx, 0) // unbuffered channel

	eb.Publish(10)

	_, ok := recvOrTimeout(sub, 100*time.Millisecond)
	if ok {
		t.Fatalf("expected slow consumer to be closed")
	}
}

// ============================================================
// stop subscription on ctx cancel
// ============================================================
func TestEventBus_SubscribeCtx_CancelStopsSubscription(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()
	defer eb.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	sub := eb.Subscribe(ctx, 1)

	cancel()

	_, ok := recvOrTimeout(sub, time.Second)
	if ok {
		t.Fatalf("expected subscriber to be closed after ctx cancel")
	}
}

// ============================================================
// ctx cancel before Stop
// ============================================================
func TestEventBus_SubscribeCtx_CancelBeforeStop(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()

	ctx, cancel := context.WithCancel(context.Background())
	sub := eb.Subscribe(ctx, 1)

	cancel()

	_, ok := recvOrTimeout(sub, time.Second)
	if ok {
		t.Fatalf("expected subscriber to close on ctx cancel")
	}

	eb.Stop() // ctx.Cancel → Stop の順でも安全
}

// ============================================================
// ctx cancel prevents delivery
// ============================================================
func TestEventBus_SubscribeCtx_CancelPreventsPublish(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()
	defer eb.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	sub := eb.Subscribe(ctx, 1)

	cancel()

	_, ok := recvOrTimeout(sub, time.Second)
	if ok {
		t.Fatalf("expected subscriber to be closed after ctx cancel")
	}

	eb.Publish(50)

	v, ok := recvOrTimeout(sub, 100*time.Millisecond)
	if ok || v != 0 {
		t.Fatalf("expected no delivery after ctx cancel, got %v (ok=%v)", v, ok)
	}
}

// ============================================================
// No deadlock on Publish after Stop
// ============================================================
func TestEventBus_PublishAfterStop_NoDeadlock(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()
	eb.Stop()

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)
		eb.Publish(999) // デッドロックせず即 return
	}()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatalf("Publish after Stop deadlocked")
	}
}

// ============================================================
// Subscribe(ctx) after Stop has closed chan, no deadlock
// ============================================================
func TestEventBus_SubscribeCtx_AfterStop_NoDeadlock(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()
	eb.Stop()

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		ctx := context.Background()
		ch := eb.Subscribe(ctx, 1)

		_, ok := recvOrTimeout(ch, time.Second)
		if ok {
			errCh <- errors.New("expected closed chan after Stop")
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatal(err.Error())
		}
	case <-time.After(time.Second):
		t.Fatalf("Subscribe(ctx) after Stop deadlocked")
	}
}

// ============================================================
// Unsubscribe after Stop has no deadlock
// ============================================================
func TestEventBus_UnsubscribeAfterStop_NoDeadlock(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()

	ctx := context.Background()
	sub := eb.Subscribe(ctx, 1)

	eb.Stop()

	done := make(chan struct{}, 1)

	go func() {
		defer close(done)
		eb.Unsubscribe(sub)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatalf("Unsubscribe after Stop deadlocked")
	}
}

// ============================================================
// Stop is idempotent
// ============================================================
func TestEventBus_StopIdempotent(t *testing.T) {
	eb := NewEventBus[int]()
	eb.Start()

	ctx := context.Background()
	sub := eb.Subscribe(ctx, 1)

	eb.Stop()
	eb.Stop()
	eb.Stop()

	_, ok := recvOrTimeout(sub, time.Second)
	if ok {
		t.Fatalf("subscriber should be closed after Stop")
	}
}
