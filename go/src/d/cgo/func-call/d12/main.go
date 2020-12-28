package main

/*
#include <stdint.h>
union A {
	int i;
	float f;
};

union B {
	int8_t i8;
	int64_t i64;
};

enum C {
	ONE,
	TWO,
};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var a C.union_A

	// [4]uint8
	fmt.Printf("%T\n", a)

	var b C.union_B

	// [8]uint8
	fmt.Printf("%T\n", b)


	// 操作 C 语言的联合类型变量：
	// 注意：对于复杂的联合类型，通过在 C 语言中定义辅助函数的方式处理。
	// 1. 在C语言中定义辅助函数；
	// 2. 通过 Go 语言的 "encoding/binary" 手工解码成员（需要注意大端小端问题）；
	// 3. 使用 unsafe 包强制转换为对应类型（这是性能最好的方式）；
	var aa C.union_A
	fmt.Println("aa.i:", *(*C.int)(unsafe.Pointer(&aa)))
	fmt.Println("aa.f:", *(*C.float)(unsafe.Pointer(&aa)))

	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}
