package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("🎉 Context...")
		}
	}
}
