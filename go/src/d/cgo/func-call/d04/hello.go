package main

//#include "hello.h"
import "C"
import "fmt"

//export Hello
func Hello(s *C.char) {
	fmt.Print(C.GoString(s))
}
