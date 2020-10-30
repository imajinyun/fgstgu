package main

import "fmt"

func main() {
	fmt.Println("Print prime number:")
	var num int = 100
	var i, j int
	for i = 2; i <= num; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
