# Event package

The `event` package provides a lightweight in-process event bus helper. It lets you define event types, register handlers, and publish events in a decoupled way without introducing an external message broker.

## Key types and functions

- `EventBus`: core event bus type.
    - `NewEventBus() *EventBus` — constructs a new bus instance.
    - `Subscribe(eventName string, handler Handler)` — registers a handler for a named event.
    - `Unsubscribe(eventName string, handler Handler)` — removes a previously-registered handler.
    - `Publish(eventName string, payload interface{})` — publishes an event to all handlers registered for that name.

- `Handler`: function type for event handlers.
    - `type Handler func(Event)` — each handler receives an `Event` value.

- `Event`: struct representing a published event.
    - Likely fields (see `event.go` for the exact definition):
        - `Name string` — event name.
        - `Payload interface{}` — arbitrary data carried with the event.
        - Additional metadata fields if needed.

## Usage

Import the package (module path):

    import "github.com/kinakoman/go-extensions/event"

Example — basic publish/subscribe:

    package main

    import (
    	"fmt"
    	"github.com/kinakoman/go-extensions/event"
    )

    func main() {
    	bus := event.NewEventBus()

    	// subscribe handler
    	bus.Subscribe("user.created", func(e event.Event) {
    		fmt.Printf("got event %s: %#v\n", e.Name, e.Payload)
    	})

    	// publish event
    	bus.Publish("user.created", map[string]interface{}{
    		"id":   123,
    		"name": "Alice",
    	})
    }
