package b07

import (
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancelChannel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			t.Logf("Channel %v is cancelled", i)
		}(i, cancelChan)
	}
	cancel2(cancelChan)
	time.Sleep(time.Second)
}
