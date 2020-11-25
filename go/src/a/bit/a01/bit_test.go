package a01

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 3, 2}
	c := [...]int{1, 2, 3, 4}
	d := [...]int{1, 2, 3}
	t.Log(a == b)
	t.Log(a == d)
	t.Log(b == d)

	// invalid operation: a == c (mismatched types [3]int and [4]int)
	// t.Log(a == c)
	_ = c
}

func TestRWX(t *testing.T) {
	t.Log(Readable)   // 1
	t.Log(Writable)   // 2
	t.Log(Executable) // 4
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	t.Log("a&Readable == Readable is", a&Readable == Readable)
	t.Log("a&Writable == Writable is", a&Writable == Writable)
	t.Log("a&Executable == Executable is", a&Executable == Executable)

	a = a &^ Readable // 0111 &^ 0001 = 0110 => 6
	t.Log(a)
	a = a &^ Writable // 0110 &^ 0010 = 0100 => 4
	t.Log(a)
	a = a &^ Executable // 0101 &^ 0110 = 0000 => 0
	t.Log(a)

	t.Log("a&Readable == Readable is", a&Readable == Readable)
	t.Log("a&Writable == Writable is", a&Writable == Writable)
	t.Log("a&Executable == Executable is", a&Executable == Executable)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)
