package main

import (
	"fmt"
)

var a int = 100
var b int = 200

// 全局变量和局部变量
func main() {
	a, b, c := 10, 20, 0
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	c = sum(a, b)
	fmt.Printf("c = %d\n", c)
}

func sum(a, b int) (c int) {
	a++
	b += 4
	c = a + b
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	return c
}
