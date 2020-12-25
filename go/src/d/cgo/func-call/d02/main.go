package main

//void Hello(const char* s);

import "C"

func main() {
	C.Hello(C.CString("🎉 Hello World, Hello C, Hello Go\n"))
}
