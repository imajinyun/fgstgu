package main

import (
	"fmt"
	"sync"
	"time"
)

// 条件变量
func main() {
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex}
	condition := false
	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("Children goroutine is locked")
		fmt.Println("Children goroutine change value and send notification")
		condition = true
		cond.Signal()
		fmt.Println("Children goroutine continue...")
		time.Sleep(5 * time.Second)
		fmt.Println("Children goroutine is unlocked")
		cond.L.Unlock()
	}()

	cond.L.Lock()
	fmt.Println("Main is locked")
	if !condition {
		fmt.Println("Main is about enter to the wait")
		cond.Wait()
		fmt.Println("Main wakes up")
	}
	fmt.Println("Main is continue...")
	fmt.Println("Main is unlocked")
	cond.L.Unlock()
}
