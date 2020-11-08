package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	fmt.Println("Non cache channel: ", len(ch1), cap(ch1))
	go func() {
		data := <-ch1
		fmt.Println("Read data from ch1: ", data)
	}()
	ch1 <- 100
	time.Sleep(time.Second)
	fmt.Println("Assign ok, main func is over!")
	fmt.Println()

	ch2 := make(chan string)
	go toSender(ch2)
	for data := range ch2 {
		fmt.Println("Read data from ch2: ", data)
	}
	fmt.Println()

	ch3 := make(chan string, 6)
	go toSender(ch3)
	for data := range ch3 {
		fmt.Println("Read data: ", data)
	}
	fmt.Println("The main func is over!")
}

func toSender(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Get data: %d", i)
		fmt.Println("Put into the channel data: ", i)
	}
	defer close(ch)
}
