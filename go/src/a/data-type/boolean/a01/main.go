package main

import "fmt"

func main() {
	var a bool
	fmt.Printf("a: %T, %v, %p\n", a, a, &a)

	a = true
	fmt.Printf("a: %T, %v, %p\n", a, a, &a)

	b := make([]bool, 3, 5)
	fmt.Printf("b: %T, %v, %p\n", b, b, &b)

	var c []bool = []bool{}
	fmt.Printf("c: %T, %v, %p\n", c, c, &c)

	d := []bool{false, false, false}
	fmt.Printf("d: %T, %v, %p\n", d, d, &d)
}
