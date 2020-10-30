package main

import "fmt"

// goto 语句
func main() {
	var i, j int
	i = 1
	fmt.Println("Print prime number:")

Looper:
	for i < 100 {
		i++
		for j = 2; j < i; j++ {
			if i%j == 0 {
				goto Looper
			}
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
