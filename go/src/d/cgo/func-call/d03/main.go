package main

//#include <hello.h>
import "C"

// go build -o main
// ./main
func main() {
	C.Hello(C.CString("🎉 Hello World, Hello C, Hello Go\n"))
}
