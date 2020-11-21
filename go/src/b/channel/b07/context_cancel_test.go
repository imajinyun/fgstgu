package b07

import (
	"context"
	"testing"
	"time"
)

func isContextCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isContextCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			t.Logf("Context %v is cancelled", i)
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
