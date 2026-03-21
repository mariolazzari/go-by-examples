package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// signal.NotifyContext returns a context that’s canceled when one of the listed signals arrives.
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// The program will wait here until one of the configured signals is received.
	fmt.Println("awaiting signal")
	<-ctx.Done()

	// context.Cause reports why the context was canceled.
	// For a signal-triggered cancellation, this includes the signal value.
	fmt.Println()
	fmt.Println(context.Cause(ctx))
	fmt.Println("exiting")
}
