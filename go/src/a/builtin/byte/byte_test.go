package byte_test

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestStrToByte(t *testing.T) {
	s := "你好Go"
	b := strToByte(s)
	t.Logf("\n%#v\n%v\n", b, b)
}

func TestByteToStr(t *testing.T) {
	b := []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0x47, 0x6f}
	s := byteToStr(b)
	t.Logf("\n%#v\n%v\n", s, s)
}

func strToByte(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

func byteToStr(b []byte) (s string) {
	data := make([]byte, len(b))
	for i, c := range b {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(b)
	return s
}
