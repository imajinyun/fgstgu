package b08

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Singleton struct.
type Singleton struct{}

var (
	instance    *Singleton
	initialized uint32
	mu          sync.Mutex
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		instance := Instance()
		t.Logf("%p", instance)
	}
}

// Instance returns a Singleton instance.
func Instance() *Singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &Singleton{}
	}
	return instance
}
