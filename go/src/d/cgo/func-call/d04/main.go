package main

// #include "hello.h"
import "C"

func main() {
	C.Hello(C.CString("🎉 Hello World, Hello C, Hello Go\n"))
}
