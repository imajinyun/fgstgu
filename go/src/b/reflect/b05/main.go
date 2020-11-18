package main

import (
	"fmt"
	"reflect"
)

// 从反射值对象中获取值
func main() {
	var a int = 1024
	valueOfA := reflect.ValueOf(a)

	var b int = valueOfA.Interface().(int)
	var c int = int(valueOfA.Int())
	fmt.Println(b, c)
}
