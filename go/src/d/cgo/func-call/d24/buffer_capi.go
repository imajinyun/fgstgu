package main

/*
#cgo CXXFLAGS: -std=c++11
#include "./buffer_capi.h"
*/
import "C"

type cgoBufferT C.Buffer_T

func cgoNewBuffer(size int) *cgoBufferT {
	p := C.NewBuffer(C.int(size))
	return (*cgoBufferT)(p)
}

func cgoDeleteBuffer(p *cgoBufferT) {
	C.DeleteBuffer((*C.Buffer_T)(p))
}

func cgoBufferData(p *cgoBufferT) *C.char {
	return C.BufferData((*C.Buffer_T)(p))
}

func cgoBufferSize(p *cgoBufferT) C.int {
	return C.BufferSize((*C.Buffer_T)(p))
}
