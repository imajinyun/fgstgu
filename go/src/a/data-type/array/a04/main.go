package main

import "fmt"

func main() {
	var a = [3]int{}
	b := [3]int{}
	c := [...]int{}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()

	// invalid operation: c == nil (mismatched types [0]int and nil)
	// if c == nil {
	// 	fmt.Println("Array does not support comparison with nil.")
	// }
}
