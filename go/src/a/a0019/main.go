package main

import "fmt"

// 递归函数
func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial2(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}
