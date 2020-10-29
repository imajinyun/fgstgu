package main

import "fmt"

func main() {
	var lines int = 10

	fmt.Println("Print multiplication table:")
	for i := 1; i < lines; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}
