package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		// time.Sleep(5 * time.Millisecond)
		a <- 100
	}()
	go func() {
		// time.Sleep(1 * time.Second)
		b <- 200
	}()

	for i := 0; i < 10; i++ {
		select {
		case data := <-a:

			fmt.Println("Read data from a:", data)
		case data := <-b:
			fmt.Println("Read data from b:", data)
		default:
			fmt.Println("Executing default...")
		}
	}
}
