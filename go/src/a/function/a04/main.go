package main

import "fmt"

// 函数传 int 型参数
func main() {
	a := 100
	fmt.Printf("1. variable a memory address: %p, value: %v\n", &a, a)
	change(a)
	fmt.Printf("3. variable a memory address: %p, value: %v\n", &a, a)
	update(&a)
	fmt.Printf("5. variable a memory address: %p, value: %v\n", &a, a)
}

func change(a int) {
	fmt.Printf("1. variable a memory address: %p, value: %d\n", &a, a)
	a = 200
}

func update(a *int) {
	fmt.Printf("4. variable a memory address: %p, value: %d\n", a, *a)
	*a = 300
}
