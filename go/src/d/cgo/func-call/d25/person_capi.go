package main

/*
#cgo CXXFLAGS: -std=c++17
#include "person_capi.h"
*/
import "C"
import "unsafe"

//export newPerson
func newPerson(name *C.char, age C.int) C.person_t {
	id := NewObjectID(NewPerson(C.GoString(name), int(age)))
	return C.person_t(id)
}

//export deletePerson
func deletePerson(p C.person_t) {
	ObjectID(p).Free()
}

//export setPerson
func setPerson(t C.person_t, name *C.char, age C.int) {
	p := ObjectID(t).Get().(*Person)
	p.Set(C.GoString(name), int(age))
}

//export getPersonName
func getPersonName(t C.person_t, buf *C.char, size C.int) *C.char {
	p := ObjectID(t).Get().(*Person)
	name, _ := p.Get()
	n := int(size) - 1
	tmp := ((*[1 << 31]byte)(unsafe.Pointer(buf)))[0:n:n]
	n = copy(tmp, []byte(name))
	tmp[n] = 0
	return buf
}

//export getPersonAge
func getPersonAge(t C.person_t) C.int {
	p := ObjectID(t).Get().(*Person)
	_, age := p.Get()
	return C.int(age)
}
