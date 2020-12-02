package main

import "fmt"

func main() {
	var a = [3]int{0}
	b := [3]int{}
	c := [...]int{}
	d := [...]int{}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()

	if a == b {
		fmt.Println("a == b is", a == b)
	}
	if c == d {
		fmt.Println("c == d is", c == d)
	}

	// invalid operation: c == nil (mismatched types [0]int and nil)
	// if c == nil {
	// 	fmt.Println("Array does not support comparison with nil.")
	// }
}
