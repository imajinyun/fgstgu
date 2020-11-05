package main

import "fmt"

func main() {
	a()
	b(100)
	c()
}

func a() {
	fmt.Println("func a()")
}

func b(x int) {
	var t [10]int
	t[x] = 888
}

func c() {
	fmt.Println("func c()")
}
