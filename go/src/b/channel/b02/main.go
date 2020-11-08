package main

import "fmt"

func main() {
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Printf("%T\n", ch1)
	ch2 := make(chan bool)
	go func() {
		data, ok := <-ch1
		if ok {
			fmt.Println("Children goroutine get data: ", data)
		}
		ch2 <- true
	}()
	ch1 <- 100
	<-ch2
	fmt.Println("🧨 Main func is over!")
}
