package main

/*
#include <string.h>
const char arr[10];
const char *str = "HelloWorld";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a0 []byte
	var aHdr = (*reflect.SliceHeader)(unsafe.Pointer(&a0))
	aHdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
	aHdr.Len = 10
	aHdr.Cap = 10
	fmt.Printf("%T %v\n", a0, a0)

	a1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
	fmt.Printf("%T %v\n", a1, a1)

	var s0 string
	var s0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(C.str))
	slen := int(C.strlen(C.str))
	s0Hdr.Len = slen
	fmt.Printf("%T %v\n", s0, s0)

	fmt.Println()

	fmt.Println(string(*C.str))
	s1 := (*[31]byte)(unsafe.Pointer(&C.str))[:slen:slen]
	fmt.Printf("%T %v %v\n", s1, s1, string(s1))
	s2 := (*[1 << 30]byte)(unsafe.Pointer(&C.str))[:slen:slen]
	fmt.Printf("%T %v %v\n", s1, s2, string(s2))
}
