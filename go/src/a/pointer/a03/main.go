package main

import "fmt"

// 空指针
func main() {
	var ptr *string

	if ptr != nil {
		fmt.Println("Not nil")
	}
	fmt.Println("It is nil")
}
