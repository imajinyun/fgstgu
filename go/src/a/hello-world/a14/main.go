package main

import "fmt"

// Go 实现自重写。
func main() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
package main

import "fmt"

func mian() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
