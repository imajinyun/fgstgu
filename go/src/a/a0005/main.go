package main

import (
	"fmt"
)

func main() {
	var a uint = 128
	var b uint = 64
	var c uint = 2
	var d uint = 32

	fmt.Printf("a & b = %d\n", a&b)
	fmt.Printf("a | b = %d\n", a|b)
	fmt.Printf("a ^ b = %d\n", a^b)
	fmt.Printf("c << 2 = %d\n", c<<2)
	fmt.Printf("d >> 2 = %d\n", d>>2)
}
