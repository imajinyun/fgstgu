package main

import "fmt"

func main() {
	a := [4]float64{12.82, 23.99, 10.01, 17.42}
	b := [...]string{"Hello", "World", "Hello", "Go"}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Print(b[i], "\t")
	}
	fmt.Println()
}
