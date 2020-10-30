package main

import "fmt"

func a(a, b int) (res int, err error) {
	return a + b, err
}

func b(a int, b int) (res int, err error) {
	return a * b, err
}

func c(a int, b int) (int, int, error) {
	return a, b, nil
}

// syntax error: mixed named and unnamed function parameters
// func d(a, b int) (res int, error) {
// 	return a + b, nil
// }

// 函数定义的几种方式
func main() {
	fmt.Println(a(3, 3))
	fmt.Println(b(4, 4))
	fmt.Println(c(5, 6))
}
