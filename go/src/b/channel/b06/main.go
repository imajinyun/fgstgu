package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go a(ch)
	data := <-ch
	fmt.Println("Received:", data)
	ch <- "Go"
	ch <- "Language"
	go b(ch)
	go c(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("The main func is over!")
}

func a(ch chan string) {
	ch <- "A"
	d1 := <-ch
	d2 := <-ch
	fmt.Println("Response:", d1, d2)
}

func b(ch chan<- string) {
	str := "B"
	ch <- str
	fmt.Println("Write:", str)
	// invalid operation: cannot receive from send-only channel ch (variable of type chan<- string)
	// data := <-ch
	// <-ch
}

func c(ch <-chan string) {
	data := <-ch
	fmt.Println("Read: ", data)

	// invalid operation: cannot send to receive-only type <-chan string
	// ch <- "C"
}
