package cgohelper

//#include <stdio.h>
import "C"

// CChar type.
type CChar C.char

// GoString return a go string.
func (p *CChar) GoString() string {
	return C.GoString((*C.char)(p))
}

// PrintCString print a string.
func PrintCString(cs *C.char) {
	C.puts(cs)
}
