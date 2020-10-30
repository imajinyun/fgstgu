package main

import "fmt"

// 闭包函数调用
func main() {
	fn1 := counter()

	fmt.Printf("fn type: %T\n", fn1)
	fmt.Println(fn1())
	fmt.Println(fn1())
	fmt.Println(fn1())

	fn2 := counter()
	fmt.Println(fn2())
	fmt.Println(fn2())
}

func counter() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	fmt.Println("counter inner func: ", res)
	return res
}
