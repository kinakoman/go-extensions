package event

import (
	"context"
	"sync"
)

type EventBus[T any] struct {
	mu          sync.RWMutex
	subscribers map[chan T]struct{}
	addSub      chan chan T
	removeSub   chan chan T
	publishCh   chan T
	stopCh      chan struct{}

	startedFlag bool
	stoppedFlag bool
}

func NewEventBus[T any]() *EventBus[T] {
	eb := &EventBus[T]{
		subscribers: make(map[chan T]struct{}),
		addSub:      make(chan chan T),
		removeSub:   make(chan chan T),
		publishCh:   make(chan T),
		stopCh:      make(chan struct{}),
	}
	return eb
}

// main loop to handle subscriptions and publishing
func (eb *EventBus[T]) run() {
	for {
		select {
		case ch := <-eb.addSub:
			eb.subscribers[ch] = struct{}{}
		case ch := <-eb.removeSub:
			if _, ok := eb.subscribers[ch]; ok {
				delete(eb.subscribers, ch)
				close(ch) // ← only when existed
			}
		case evt := <-eb.publishCh:
			eb.mu.RLock()
			for ch := range eb.subscribers {
				select {
				case ch <- evt:
					// subscriber に送信成功
				default:
					// slow consumer → drop
					eb.mu.RUnlock()
					eb.mu.Lock()
					if _, ok := eb.subscribers[ch]; ok {
						delete(eb.subscribers, ch)
						close(ch) // ← only when existed
					}
					eb.mu.Unlock()
					eb.mu.RLock()
				}
			}
			eb.mu.RUnlock()

		case <-eb.stopCh:
			// close all subscriber channels
			for ch := range eb.subscribers {
				close(ch)
			}
			return
		}
	}
}

func (eb *EventBus[T]) Subscribe(ctx context.Context, buffer int) chan T {
	eb.mu.RLock()
	stopped := eb.stoppedFlag
	started := eb.startedFlag
	eb.mu.RUnlock()

	if !started {
		panic("EventBus must be started before subscribing. Call Start() method first.")
	}

	ch := make(chan T, buffer)

	if stopped {
		close(ch)
		return ch
	}

	select {
	case eb.addSub <- ch:
	// if registration succeeds, move to next step
	case <-eb.stopCh:
		close(ch)
		return ch
	}

	go func() {
		select {
		case <-ctx.Done():
			eb.Unsubscribe(ch)
		case <-eb.stopCh:
			// EventBus stopped, no action needed
		}
	}()

	return ch
}

func (eb *EventBus[T]) Unsubscribe(ch chan T) {
	eb.mu.RLock()
	stopped := eb.stoppedFlag
	started := eb.startedFlag
	eb.mu.RUnlock()

	if !started {
		panic("EventBus must be started before unsubscribing. Call Start() method first.")
	}

	if stopped {
		return
	}

	select {
	case eb.removeSub <- ch:
	case <-eb.stopCh:
	}
}

func (eb *EventBus[T]) Publish(evt T) {
	eb.mu.RLock()
	stopped := eb.stoppedFlag
	started := eb.startedFlag
	eb.mu.RUnlock()

	if !started {
		panic("EventBus must be started before publishing. Call Start() method first.")
	}

	if stopped {
		return
	}

	select {
	case eb.publishCh <- evt:
	case <-eb.stopCh:
	}
}

func (eb *EventBus[T]) Start() {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if eb.startedFlag {
		panic("EventBus has already been started.")
	}

	eb.startedFlag = true
	if eb.stoppedFlag {
		panic("Cannot start EventBus after it has been stopped.")
	}
	ready := make(chan struct{})
	go func() {
		close(ready)
		eb.run()
	}()
	<-ready
}

func (eb *EventBus[T]) Stop() {
	eb.mu.Lock()

	if eb.stoppedFlag {
		eb.mu.Unlock()
		return
	}

	eb.stoppedFlag = true
	close(eb.stopCh)
	eb.mu.Unlock()
}
