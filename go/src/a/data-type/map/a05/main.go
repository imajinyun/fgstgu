package main

import "fmt"

func main() {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op + op }
	m[3] = func(op int) int { return op * op }
	fmt.Printf("m: %T, %v, %p\n", m, m, &m)
	fmt.Println(m[1](1))
	fmt.Println(m[2](2))
	fmt.Println(m[3](3))

	n := map[int]bool{}
	n[1] = true
	fmt.Printf("Length: %d\n", len(n))
	if n[1] {
		fmt.Printf("Element `%v` is exists.\n", n[1])
	} else {
		fmt.Printf("Element `%v` is not exists.\n", n[1])
	}

	delete(n, 1)
	if n[1] {
		fmt.Printf("Element `%v` is exists.\n", n[1])
	} else {
		fmt.Printf("Element `%v` is not exists.\n", n[1])
	}
}
