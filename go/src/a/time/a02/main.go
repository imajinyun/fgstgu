package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("%T\n", timer)
	fmt.Println(time.Now())
	data := <-timer.C
	fmt.Printf("%T\n", timer.C)
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}
