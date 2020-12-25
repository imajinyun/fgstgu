package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan int, 64)
	go producer(3, ch)
	go producer(5, ch)
	go consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func producer(factor int, push chan<- int) {
	for i := 0; ; i++ {
		push <- i * factor
	}
}

func consumer(pull <-chan int) {
	for v := range pull {
		fmt.Println(v)
	}
}
