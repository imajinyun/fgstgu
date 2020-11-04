package main

import "fmt"

// 数组定义
func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a =", a)

	b := [5]int{5, 6, 7, 8, 9}
	fmt.Println("b =", b)

	c := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println("c =", c)

	d := [...]float64{1, 3, 5, 7, 9}
	fmt.Println("d =", d)
}
