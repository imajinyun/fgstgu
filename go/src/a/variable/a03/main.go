package main

import "fmt"

type myFunc func(int) bool

// 函数变量
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice = ", slice)

	even := filter(slice, isEven)
	fmt.Println("Even elements: ", even)

	odd := filter(slice, isOdd)
	fmt.Println("Odd elements: ", odd)
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func filter(slice []int, fn myFunc) []int {
	var result []int
	for _, value := range slice {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}
