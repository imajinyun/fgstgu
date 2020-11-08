package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
		ch <- 200
		close(ch)
		ch <- 999
	}()

	for i := 0; i < 10; i++ {
		data, ok := <-ch
		fmt.Printf("Read data %d: %v, %v\n", i, data, ok)
	}
}
