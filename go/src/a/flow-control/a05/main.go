package main

import "fmt"

func main() {
Looper1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue Looper1
			}
		}
	}

	fmt.Println()

Looper2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break Looper2
			}
		}
	}
}
