package main

/*
#include <stdlib.h>
void *makeslice(size_t memsize) {
	return malloc(memsize);
}
*/
import "C"
import "unsafe"

func makeByteSlice(n int) []byte {
	p := C.makeslice(C.size_t(n))
	return ((*[1<<32 + 1]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	s := makeByteSlice(1<<32 + 1)
	s[len(s)-1] = 255
	println(s[len(s)-1])
	freeByteSlice(s)
}
