package main

import "fmt"

func main() {
	{
		defer fmt.Println("defer...")
		fmt.Println("block...")
	}
	fmt.Println("main...")
}
