package main

import (
	"fmt"
	"sync"
)

var (
	counter      int
	counterGuard sync.Mutex
)

func main() {
	setCounter(100)
	fmt.Println(getCounter())
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
