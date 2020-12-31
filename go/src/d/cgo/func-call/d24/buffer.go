package main

import "unsafe"

// Buffer struct.
type Buffer struct {
	cgoPtr *cgoBufferT
}

// NewBuffer returns a new Buffer.
func NewBuffer(size int) *Buffer {
	return &Buffer{
		cgoPtr: cgoNewBuffer(size),
	}
}

// Delete to delete hold's buffer.
func (p *Buffer) Delete() {
	cgoDeleteBuffer(p.cgoPtr)
}

// Data to get buffer data.
func (p *Buffer) Data() []byte {
	data := cgoBufferData(p.cgoPtr)
	size := cgoBufferSize(p.cgoPtr)
	return ((*[1 << 31]byte)(unsafe.Pointer(data)))[0:int(size):int(size)]
}
