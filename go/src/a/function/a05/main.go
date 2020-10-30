package main

import "fmt"

// 函数传 slice 型参数
func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("1. slice variable a memory address: %p, value: %v\n", &a, a)
	change(a)
	fmt.Printf("3. slice variable a memory address: %p, value: %v\n", &a, a)
	update(&a)
	fmt.Printf("5. slice variable a memory address: %p, value: %v\n", &a, a)
}

func change(a []int) {
	fmt.Printf("2. slice variable a memory address: %p, value: %v\n", &a, a)
	a[0] = 9
}

func update(a *[]int) {
	fmt.Printf("4. slice variable a memory address: %p, value: %v\n", a, *a)
	(*a)[1] = 8
}
