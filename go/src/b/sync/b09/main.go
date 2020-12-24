package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Singleton struct.
type Singleton struct{}

// Once struct.
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do method for Once.
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

var (
	instance *Singleton
	once     sync.Once
)

// Instance method for Singleton.
func Instance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func loadConfig() map[string]int {
	return map[string]int{"a": 10, "b": 11, "c": 12}
}

func main() {
	var config atomic.Value
	config.Store(loadConfig())

	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
			fmt.Println(config.Load())
		}
	}()
	time.Sleep(5 * time.Second)
}
