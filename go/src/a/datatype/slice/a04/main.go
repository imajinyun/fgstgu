package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a []string
	b := make([]string, 0, 16)
	dump("a", a)
	dump("b", b)
	fmt.Println()

	for i := 0; i <= 16; i++ {
		a = append(a, strconv.Itoa(i))
		dump("a", a)
	}
	fmt.Println()

	for i := 0; i <= 16; i++ {
		b = append(b, strconv.Itoa(i))
		dump("b", b)
	}
}

func dump(name string, x []string) {
	fmt.Print(name, ": ")
	fmt.Printf("addr=%p\tlen=%v\tcap=%d\tvalue=%v\n", x, len(x), cap(x), x)
}
