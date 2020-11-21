package b07

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func oneResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestOneResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(oneResponse())
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}
