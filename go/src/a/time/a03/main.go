package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.After(5 * time.Second)
	fmt.Println(time.Now())
	data := <-ch
	fmt.Printf("%T\n", data)
	fmt.Println(data)
}
