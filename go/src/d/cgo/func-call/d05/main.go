package main

//void Hello(char* s);
import "C"
import "fmt"

func main() {
	C.Hello(C.CString("🎉 Hello World, Hello C, Hello Go\n"))
}

//export Hello
func Hello(s *C.char) {
	fmt.Print(C.GoString(s))
}
