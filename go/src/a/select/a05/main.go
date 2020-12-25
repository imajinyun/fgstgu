package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(ch)
	}
	time.Sleep(time.Second)
	close(ch)
}

func worker(ch chan bool) {
	for {
		select {
		case <-ch:
		default:
			fmt.Println("🎉 Go")
		}
	}
}
