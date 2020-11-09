package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
func main() {
	var mutex sync.RWMutex
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Goroutine %d try to read lock\n", i)
			mutex.RLock()
			fmt.Printf("Goroutine %d is read locked\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("Goroutine %d read unlock\n", i)
			mutex.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Main try to read lock")
	mutex.Lock()
	fmt.Println("Main is read locked")
	mutex.Unlock()
	fmt.Println("Main read unlock")
}
