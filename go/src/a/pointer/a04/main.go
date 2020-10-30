package main

import "fmt"

// 使用指针作为函数的参数
func main() {
	a := 50
	fmt.Printf("a = %d\n", a)
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("&a = %X\n", &a)

	var b *int = &a
	change(b)
	fmt.Printf("a = %d, b = %d\n", a, *b)
}

func change(value *int) {
	*value = 100
}
