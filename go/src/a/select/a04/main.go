package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go worker(ch)
	time.Sleep(time.Second)
	ch <- true
}

func worker(ch chan bool) {
	for {
		select {
		case <- ch:
		default:
			fmt.Println("🎉 OK")
		}
	}
}
