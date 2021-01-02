package main

import "C"
import (
	_ "fgstgu/go/src/d/cgo/func-call/d31/number"
	"fmt"
)

// 需要先注释掉 mina.c 代码
// -> go build -buildmode=c-archive -o main.a
func main() {
	println("Done...")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln:", C.GoString(s))
}
