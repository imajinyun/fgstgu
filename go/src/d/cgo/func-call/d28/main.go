package main

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include <number.h>
import "C"
import "fmt"

// -> cd ./number
// -> gcc -shared -o libnumber.so number.c
// -> cd ../
// -> go build -o main
// -> ./main
func main() {
	fmt.Println(C.average(117, 48))
}
