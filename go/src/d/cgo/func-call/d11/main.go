package main

/*
struct A {
	// 位字段无法访问
	int size: 10;

	// 零长的数组也无法访问
	float arr[];
};
*/
import "C"
import (
	"fmt"
)

func main() {
	var a C.struct_A
	fmt.Println(a)

	// 报错：位字段无法访问
	fmt.Println(a.size)

	// 报错：零长的数组无法访问
	fmt.Println(a.arr)
}
