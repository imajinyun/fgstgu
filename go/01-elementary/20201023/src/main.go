package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Printf("The length of a: %d\n", len(a))
	fmt.Printf("The capacity of a: %d\n", cap(a))
	fmt.Printf("The value of a: %d\n", a)

	fmt.Printf("\n")

	b := make([]int, 5, 9)
	fmt.Printf("The length of b: %d\n", len(b))
	fmt.Printf("The capacity of b: %d\n", cap(b))
	fmt.Printf("The value of b: %d\n", b)
}
