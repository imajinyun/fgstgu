package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 1024
	valueOfA := reflect.ValueOf(a)
	fmt.Printf("%T, %v\n", valueOfA, valueOfA)

	// panic: reflect: reflect.Value.SetInt using unaddressable value
	// valueOfA.SetInt(2048)

	// 获取变量 a 的反射值对象，即 a 的地址
	valueOfA = reflect.ValueOf(&a)

	// 取出变量 a 地址的元素，即 a 的值
	valueOfA = valueOfA.Elem()

	// 修改变量 a 的值
	valueOfA.SetInt(2048)
	fmt.Printf("%T, %v\n", valueOfA, valueOfA)
}
