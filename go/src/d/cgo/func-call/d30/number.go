package main

import "C"

// -> go build -buildmode=c-shared -o number.so
func main() {}

//export average
func average(a, b C.int) C.int {
	return (a + b) / 2
}
