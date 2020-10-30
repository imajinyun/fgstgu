package main

import (
	"fmt"
	"math"
)

// 匿名函数
func main() {

	// 在定义时调用匿名函数
	func(str string) {
		fmt.Println("Hello, ", str)
	}("World")

	// 将匿名函数赋值给变量
	fn := func(str string) {
		fmt.Println("Hello, ", str)
	}
	fn("Go Go Go!!!")

	// 匿名函数用作回调函数
	arr := []float64{1, 9, 16, 25, 36}
	visit(arr, func(v float64) {
		v = math.Sqrt(v)
		fmt.Printf("%.2f\t", v)
	})
	fmt.Println()
	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.0f\t", v)
	})
	fmt.Println()
}

func visit(list []float64, fn func(float64)) {
	for _, value := range list {
		fn(value)
	}
}
