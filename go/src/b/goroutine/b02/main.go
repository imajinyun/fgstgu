package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	var input string
	fmt.Scanln(&input)
}

func test() {
	var times int
	for {
		times++
		fmt.Println("Times: ", times)
		time.Sleep(time.Second)
	}
}
