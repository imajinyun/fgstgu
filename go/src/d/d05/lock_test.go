// -> go test -bench=. -cpuprofile=cpu.prof
// -> go tool pprof cpu.prof

package d05

import (
	"fmt"
	"sync"
	"testing"
)

var cache map[string]string

const numOfReader int = 50
const readerTimes = 100000

func init() {
	cache = make(map[string]string)
	cache["a"] = "aaa"
	cache["b"] = "bbb"
}

func lockFreeAccess() {
	var wg sync.WaitGroup
	wg.Add(numOfReader)
	for i := 0; i < numOfReader; i++ {
		go func() {
			for j := 0; j < readerTimes; j++ {
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func lockAccess() {
	var wg sync.WaitGroup
	wg.Add(numOfReader)
	m := new(sync.RWMutex)
	for i := 0; i < numOfReader; i++ {
		go func() {
			for j := 0; j < readerTimes; j++ {
				m.RLock()
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
				m.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkLockFree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockFreeAccess()
	}
}

func BenchmarkLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockAccess()
	}
}
