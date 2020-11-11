package main

import "fmt"

// EmptyInterface interface.
type EmptyInterface interface{}

// EmptyStruct struct.
type EmptyStruct struct{}

// EmptyFunc func.
func EmptyFunc() {}

func main() {
	DumpType(10)
	DumpType("Hello World")
	DumpType('A')
	DumpType(true)
	DumpType(map[int]int{})
	DumpType([3]int{})
	DumpType(EmptyStruct{})
	DumpType(EmptyFunc)
}

// DumpType func.
func DumpType(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Printf("%T, %v\n", v, v)
	case string:
		fmt.Printf("%T, %v\n", v, v)
	case bool:
		fmt.Printf("%T, %v\n", v, v)
	case int32:
		fmt.Printf("%T, %v\n", v, v)
	default:
		fmt.Printf("Other type: %T\n", v)
	}
}
