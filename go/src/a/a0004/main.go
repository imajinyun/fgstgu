package main

import (
	"fmt"
)

var a byte
var b rune

var c bool

var a1 int
var a2 uint
var b1 int8
var b2 uint8
var c1 int16
var c2 uint16
var d1 int32
var d2 uint32
var e1 int64
var e2 uint64

var f1 complex64
var f2 complex128

func main() {
	fmt.Println("Output byte|tune:")
	fmt.Printf("a type: %T\ta value: %v\n", a, a)
	fmt.Printf("b type: %T\tb value: %v\n", b, b)
	fmt.Println()

	fmt.Println("Output bool:")
	fmt.Printf("c type: %T\tc value: %v\n", c, c)
	fmt.Println()

	fmt.Println("Output int|int8|int16|int32|int64:")
	fmt.Printf("a1 type: %T\ta1 value: %v\n", a1, a1)
	fmt.Printf("a2 type: %T\ta2 value: %v\n", a2, a2)
	fmt.Printf("b1 type: %T\tb1 value: %v\n", b1, b1)
	fmt.Printf("b2 type: %T\tb2 value: %v\n", b2, b2)
	fmt.Printf("c1 type: %T\tc1 value: %v\n", c1, c1)
	fmt.Printf("c2 type: %T\tc2 value: %v\n", c2, c2)
	fmt.Printf("d1 type: %T\td1 value: %v\n", d1, d1)
	fmt.Printf("d2 type: %T\td2 value: %v\n", d2, d2)
	fmt.Printf("e1 type: %T\te1 value: %v\n", e1, e1)
	fmt.Printf("e2 type: %T\te2 value: %v\n", e2, e2)
	fmt.Println()

	fmt.Println("Output complex64|complex128:")
	fmt.Printf("f1 type: %T\tf1 value: %v\n", f1, f1)
	fmt.Printf("f2 type: %T\tf2 value: %v\n", f2, f2)
}
