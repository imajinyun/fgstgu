package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 64)
	go producer(2, ch)
	go producer(3, ch)
	go consumer(ch)
	time.Sleep(300 * time.Microsecond)
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
