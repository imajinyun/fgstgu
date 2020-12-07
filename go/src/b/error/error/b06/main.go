package main

import (
	"fmt"
	"reflect"
)

type MyError struct{}

func (*MyError) Error() string {
	return ""
}

func main() {
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()

	// *MyError 指针类型实现了 error 接口
	myErrorPtr := reflect.TypeOf(&MyError{})

	// MyError 类型并没有实现 error 接口
	myError := reflect.TypeOf(MyError{})

	fmt.Println(myErrorPtr.Implements(typeOfError))
	fmt.Println(myError.Implements(typeOfError))
}
