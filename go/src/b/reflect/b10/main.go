package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	typeOfA := reflect.TypeOf(a)
	ins := reflect.New(typeOfA)
	fmt.Println(ins.Type(), ins.Kind())

	var b string
	typeOfB := reflect.TypeOf(b)
	ins = reflect.New(typeOfB)
	fmt.Println(ins.Type(), ins.Kind())

	var c struct{}
	typeOfC := reflect.TypeOf(c)
	ins = reflect.New(typeOfC)
	fmt.Println(ins.Type(), ins.Kind())
}
