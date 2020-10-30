package main

import "fmt"

func main() {
	var a = [3]int{}
	b := [3]int{}
	c := [...]int{}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()

	// invalid operation: c == nil (mismatched types [0]int and nil)
	// if c == nil {
	// 	fmt.Println("Array does not support comparison with nil.")
	// }

	var d []int = make([]int, 3)
	e := make([]int, 3)
	var f []int
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	if f == nil {
		fmt.Println("slice f is nil")
	}

	var numbers = make([]int, 3, 5)
	fmt.Printf("numbers type: %T\n", numbers)
	fmt.Printf("len = %d, cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)
	fmt.Println()

	numbers = []int{1, 2, 3}
	m := []int{1, 2, 3, 4, 5}
	n := [5]int{1, 2, 3, 4, 5}
	t := m[2:5]
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("n type: %T\n", n)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("t type: %T\n", t)
	fmt.Println()

	p := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", p[0:])
	fmt.Printf("%v\n", p[:10])
	fmt.Printf("%v\n", p[3:6])
}
