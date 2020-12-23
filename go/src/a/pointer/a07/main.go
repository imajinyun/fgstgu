package main

import "fmt"

func main() {
	fmt.Println(*a(100))
	fmt.Println(b())
}

func a(x int) *int {
	return &x
}

func b() int {
	x := new(int)
	return *x
}
