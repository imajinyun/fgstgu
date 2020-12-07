package main

import "reflect"

// -> go run -gcflags="-S -N" main.go
func main() {
	i := 99
	_ = reflect.TypeOf(i)
}
