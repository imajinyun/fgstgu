package main

import (
	"fmt"
	"reflect"
)

// 反射值对象的零值和有效性判断
func main() {
	// 1. IsNil 常被用于判断指针是否为空；
	// 2. IsValid 常被用于判定返回值是否有效；
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	fmt.Println("var a *int:", reflect.ValueOf(a).IsValid())

	// panic: reflect: call of reflect.Value.IsNil on zero Value
	// fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsNil())
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	s := struct{}{}
	// panic: reflect: call of reflect.Value.IsNil on zero Value
	// fmt.Println("Struct not exist member:", reflect.ValueOf(s).FieldByName("").IsNil())
	fmt.Println("Struct not exist member:", reflect.ValueOf(s).FieldByName("").IsValid())

	m := map[int]int{}
	// panic: reflect: call of reflect.Value.IsNil on zero Value
	// fmt.Println("Map not exist key:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(1)).IsNil())
	fmt.Println("Map not exist key:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(1)).IsValid())
}
