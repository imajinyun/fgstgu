package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}
	func(array [3]int) {
		array[0] = 9
		fmt.Println(array)
	}(x)
	fmt.Println(x)

	fmt.Println()

	y := []int{1, 2, 3}
	func(slice []int) {
		slice[0] = 8
		fmt.Println(slice)
	}(y)
	fmt.Println(y)
}
