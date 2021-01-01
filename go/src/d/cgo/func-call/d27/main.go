package main

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include <number.h>
import "C"
import "fmt"

// -> cd ./number
// -> gcc -c -o number.o number.c
// -> ar rcs libnumber.a number.o
// -> cd ../
// -> go build -o main
// -> ./main
func main() {
	fmt.Println(C.number_add_mod(2, 3, 3))
}
