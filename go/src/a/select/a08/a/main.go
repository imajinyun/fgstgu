package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	go func() {
		time.Sleep(time.Second)
		fmt.Println(i)
	}()

	i = 2
	for {
	}
}
