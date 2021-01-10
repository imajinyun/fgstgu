package main

import (
	"fmt"
)

func main() {
	for i := range random(100) {
		fmt.Println(i)
	}
}

func random(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}
