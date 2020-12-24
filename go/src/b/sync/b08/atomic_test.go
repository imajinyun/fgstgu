package b08

import (
	"sync"
	"sync/atomic"
	"testing"
)

var atomicTotal uint64

func TestAtomicWorker(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go atomicWorker(&wg)
	go atomicWorker(&wg)
	wg.Wait()
	t.Log(atomicTotal)
}

func atomicWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&atomicTotal, i)
	}
}
