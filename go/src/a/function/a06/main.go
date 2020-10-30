package main

import "fmt"

// 函数传数组
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("1. array variable a memory address: %p, value: %v\n", &a, a)
	change(a)
	fmt.Printf("3. array variable a memory address: %p, value: %v\n", &a, a)
	update(&a)
	fmt.Printf("5. array variable a memory address: %p, value: %v\n", &a, a)
}

func change(a [5]int) {
	fmt.Printf("2. array variable a memory address: %p, value: %v\n", &a, a)
	a[0] = 9
}

func update(a *[5]int) {
	fmt.Printf("4. array variable a memory address: %p, value: %v\n", &a, *a)
	(*a)[1] = 8
}
