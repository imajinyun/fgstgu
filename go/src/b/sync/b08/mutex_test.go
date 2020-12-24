package b08

import (
	"sync"
	"testing"
)

var mutexTotal struct {
	sync.Mutex
	value int
}

func TestMutexWorker(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go mutexWorker(&wg)
	go mutexWorker(&wg)
	wg.Wait()
	t.Log(mutexTotal.value)
}

func mutexWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutexTotal.Mutex.Lock()
		mutexTotal.value++
		mutexTotal.Mutex.Unlock()
	}
}
