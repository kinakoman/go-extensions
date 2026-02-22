# Exchan package

The `exchan` package provides an extended channel wrapper (Exchan) that ensures safety operations with context cancellation and prevents panic when sending to closed channels. It includes a done channel for simple synchronization.

## Key types and functions

- `Exchan[T]`: A generic wrapper for `chan T` with additional safety features.
- `New[T](buffer int) *Exchan[T]`: Creates a new `Exchan` with the specified buffer size.
- `Send(ctx context.Context, v T) error`: Sends a value to the channel. Returns `ErrChannelClosed` if the channel is closed, or `ctx.Err()` if the context is cancelled.
- `Recv(ctx context.Context) (T, error)`: Receives a value from the channel. Returns `ErrChannelClosed` if the channel is closed, or `ctx.Err()` if the context is cancelled.
- `Close()`: Closes the channel cleanly, preventing further sends but allowing remaining items to be received (in typical usage, `Recv` returns zero and error after close, implementation detail: `Recv` checks `done` first).
- `Done() <-chan struct{}`: Returns a read-only channel that is closed when `Close()` is called.
- `Len() int`, `Cap() int`: Returns the length and capacity of the internal channel.

## Usage

Import the package:

    import "github.com/kinakoman/go-extensions/exchan"

Example:

    package main

    import (
    	"context"
    	"fmt"
    	"github.com/kinakoman/go-extensions/exchan"
    )

    func main() {
    	ctx := context.Background()
    	e := exchan.New[string](10)
    	defer e.Close()

    	// Safe send
    	err := e.Send(ctx, "hello")
    	if err != nil {
    		fmt.Println("Send error:", err)
    		return
    	}

    	// Safe receive
    	msg, err := e.Recv(ctx)
    	if err != nil {
    		fmt.Println("Recv error:", err)
    		return
    	}
    	fmt.Println("Received:", msg)
    }
