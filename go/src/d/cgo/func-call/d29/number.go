package main

import "C"

// -> go build -buildmode=c-archive -o number.a
func main() {}

//export average
func average(a, b C.int) C.int {
	return (a + b) / 2
}
