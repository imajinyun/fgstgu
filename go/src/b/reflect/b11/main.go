package main

import (
	"fmt"
	"reflect"
)

func main() {
	caller := reflect.ValueOf(add)
	params := []reflect.Value{
		reflect.ValueOf(100),
		reflect.ValueOf(88),
	}
	result := caller.Call(params)
	fmt.Println(result[0].Int())
}

func add(a, b int) int {
	return a + b
}
