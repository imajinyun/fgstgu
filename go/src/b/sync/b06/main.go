package main

import (
	"fmt"
	"sync"
)

var (
	counter      int
	counterGuard sync.Mutex
)

var (
	value      int
	valueGuard sync.RWMutex
)

func main() {
	setCounter(100)
	fmt.Println(getCounter())
	fmt.Println(getValue())
}

func getCounter() int {
	counterGuard.Lock()
	defer counterGuard.Unlock()
	return counter
}

func setCounter(x int) {
	counterGuard.Lock()
	counter = x
	counterGuard.Unlock()
}

func getValue() int {
	valueGuard.RLock()
	defer valueGuard.RUnlock()
	return value
}
