package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(50 * time.Microsecond)
	fmt.Println("🧨 Main func is over!")
}

func hello() {
	fmt.Println("🎉 Hello goroutine!")
}
