package main

import "fmt"

// 指针的指针
func main() {
	var a int
	var b *int
	var c **int
	a = 100
	b = &a
	fmt.Println("b addr: ", b)
	c = &b
	fmt.Println("c addr: ", c)
	fmt.Printf("a = %d, *b = %d, **c = %d\n", a, *b, **c)
}
