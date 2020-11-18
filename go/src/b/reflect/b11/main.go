package main

import (
	"fmt"
	"reflect"
)

/**
 * -> cd /path/to/go
 * -> export GOPATH=/path/to/go/
 * -> go build -o main a/package-path/a02
 * -> ./main
 */
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
