package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	typeOf := reflect.TypeOf(a)
	fmt.Printf("Name: %v, Kind: %v\n", typeOf.Name(), typeOf.Kind())
}
