package main

/*
#include <stdio.h>
static void Hello(const char* s) {
	puts(s);
}
*/
import "C"

func main() {
	C.Hello(C.CString("🎉 Hello World, Hello C, Hello Go\n"))
}
