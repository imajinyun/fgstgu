package main

import "fmt"

// MyStruct struct.
type MyStruct struct{}

// NilOrNot func.
func NilOrNot(v interface{}) bool {
	return v == nil
}

// Go 语言的接口类型不是任意类
func main() {
	var s *MyStruct
	fmt.Println(s == nil)

	// 调用 NilOrNot 函数时发生了隐式的类型转换。
	// 在类型转换时，*MyStruct 类型会转换成 interface{} 类型，
	// 转换后的变量不仅包含转换前的变量，还包含变量的类型信息 MyStruct，
	// 所以转换后的变量与 nil 不相等。
	fmt.Println(NilOrNot(s))
}
