package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		data := <-a
		fmt.Println("a channel data: ", data)
	}()
	go func() {
		time.Sleep(3 * time.Second)
		data := <-b
		fmt.Println("b channel data: ", data)
	}()

	select {
	case a <- 100:
		close(a)
		fmt.Println("Write data to a channel")
	case b <- 200:
		close(b)
		fmt.Println("Write data to b channel")
	case <-time.After(2 * time.Millisecond):
		fmt.Println("Executing program in delay channel")
	}
	time.Sleep(4 * time.Second)
	fmt.Println("The main func is over!")
}
