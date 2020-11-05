package main

import "fmt"

func main() {
	a()
	b()
	c()
}

func a() {
	fmt.Println("func a()")
}

func b() {
	panic("func b(): panic")
}

func c() {
	fmt.Printf("func c()")
}
