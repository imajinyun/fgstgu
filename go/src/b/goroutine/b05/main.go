package main

import (
	"fmt"
	"time"
)

func main() {
	go running()
	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("Times:", times)
		time.Sleep(time.Second)
	}
}
