package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)

	// Goroutine to check for signals to gracefully finish all funcs.
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
		select {
		case sig := <-signalChannel:
			fmt.Printf("Received signal: %s\n", sig)
			done()
		case <-gctx.Done():
			fmt.Println("Closing signal goroutine")
			return gctx.Err()
		}
		return nil
	})

	// Just a ticker every 2s.
	g.Go(func() error {
		ticker := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker 2s ticked")
			case <-gctx.Done():
				fmt.Println("Closing ticker 2s goroutine")
				return gctx.Err()
			}
		}
	})

	// Just a ticker every 1s.
	g.Go(func() error {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker 1s ticked")
			case <-gctx.Done():
				fmt.Println("Closing ticker 1s goroutine")
				return gctx.Err()
			}
		}
	})

	// force a stop after 60s.
	time.AfterFunc(60*time.Second, func() {
		fmt.Println("Force finished after 60s")
		done()
	})

	// Wait for all errgroup goroutines.
	err := g.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("Context was canceled")
		} else {
			fmt.Printf("Received error: %v", err)
		}
	} else {
		fmt.Println("Finished clean")
	}
}
