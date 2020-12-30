package main

import (
	"fgstgu/go/src/d/cgo/func-call/d19/qsort"
	"fmt"
)

func main() {
	values := []int{731, 62, 381, 77, 3, 25, 0, 99, 1984}
	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)

	values = []int{100, 7, 18, 64, 128, 2, 0, 77, 91, 13}
	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
}
