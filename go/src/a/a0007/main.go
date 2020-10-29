package main

import (
	"fmt"
)

func main() {
	var n int = 10

	fmt.Println("The first grammatical form of the for loop:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	j := 0
	for ; j < n; j++ {
		fmt.Printf("%d ", j)
	}
	fmt.Println()

	k := 0
	for ; ; k++ {
		if k > n-1 {
			break
		}
		fmt.Printf("%d ", k)
	}
	fmt.Println()

	fmt.Println("\nThe second grammatical form of the for loop:")
	p := 0
	for p < n {
		fmt.Printf("%d ", p)
		p++
	}
	fmt.Println()

	fmt.Println("\nThe third grammatical form of the for loop:")
	q := 0
	for {
		if q > n-1 {
			break
		}
		fmt.Printf("%d ", q)
		q++
	}
	fmt.Println()

	fmt.Println("\nThe fourth grammatical form of the for loop:")
	str := "01a34b67c9d"
	for k, v := range str {
		fmt.Printf("Index: %3d, Value: %3d, Char: %3c\n", k, v, v)
	}
	fmt.Println()
}
