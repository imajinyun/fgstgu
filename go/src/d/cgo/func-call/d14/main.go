package main

/*
#include <string.h>
const char *myStringArray[] = {
  "NAME_OF_FIRST_THING",
  "NAME_OF_SECOND_THING",
  "NAME_OF_THIRD_THING"
};
const char *myString = "Hello World";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	arrLength, strLength := 3, int(C.strlen(C.myString))
	arr := (*[1 <<31]*C.char)(unsafe.Pointer(&C.myStringArray))[:arrLength:arrLength]
	byt := (*[31]byte)(unsafe.Pointer(&C.myString))[:strLength:strLength]

	for _, s := range arr {
		fmt.Println(C.GoString(s))
	}
	fmt.Println(byt)
}

func byteToStr(b []byte) (s string) {
	data := make([]byte, len(b))
	for i, c := range b {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(b)
	return s
}
