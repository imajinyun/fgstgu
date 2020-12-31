package main

//#include <stdio.h>
import "C"
import (
	"unsafe"
)

// -> go build -o main
// -> ./main
func main() {
	buf := NewBuffer(1024)
	defer buf.Delete()
	copy(buf.Data(), []byte("🎉 Hello World!!!\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
