package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int, 10)
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("🎉 Hello World, Hello Go!")
		}()
	}
	for i := 0; i < cap(done); i++ {
		time.Sleep(time.Second)
	}
}
