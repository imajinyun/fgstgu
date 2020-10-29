package main

import (
	"fmt"
)

func main() {
	lines := 9

	fmt.Println("Print right triangle:")
	for i := 0; i < lines; i++ {
		for j := 0; j <= 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Print up isosceles triangle:")
	for i := 0; i < lines; i++ {
		for j := lines; j >= 0; j-- {
			if i >= j {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Print down isosceles triangle:")
	for i := 0; i < lines; i++ {
		for j := 0; j < lines; j++ {
			if i <= j {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}
