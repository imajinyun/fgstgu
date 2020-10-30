package main

import "fmt"

func main() {
	var a []int = make([]int, 3)
	b := make([]int, 3)
	var c []int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	if c == nil {
		fmt.Println("slice c is nil")
	}
	fmt.Println()

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
