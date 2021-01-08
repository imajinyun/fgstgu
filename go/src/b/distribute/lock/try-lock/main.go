package main

import "sync"

// Lock struct.
type Lock struct {
	ch chan struct{}
}

// NewLock returns a new Lock.
func NewLock() Lock {
	var lock Lock
	lock.ch = make(chan struct{}, 1)
	lock.ch <- struct{}{}
	return lock
}

// Lock method for Lock.
func (l Lock) Lock() bool {
	res := false
	select {
	case <-l.ch:
	default:
	}
	return res
}

// Unlock method for Lock.
func (l Lock) Unlock() {
	l.ch <- struct{}{}
}

var counter int

func main() {
	var lock = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock() {
				println("lock failed")
				return
			}
			counter++
			println("current counter:", counter)
			lock.Unlock()
		}()
	}
	wg.Wait()
}
