package main

import "fmt"

func main() {
	var a int = 100
	var b *int
	b = &a

	fmt.Printf("a type: %T, value: %v\n", a, a)
	fmt.Printf("&a type: %T, value: %v\n", &a, &a)
	fmt.Printf("*&a type: %T, value: %v\n", *&a, *&a)
	fmt.Printf("iptr type: %T, value: %v\n", b, b)
	fmt.Printf("*iptr type: %T, value: %v\n", *b, *b)
}
