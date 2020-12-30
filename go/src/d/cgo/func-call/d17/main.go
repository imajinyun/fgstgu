package main

//extern int goQsortCompare(void *a, void *b);
import "C"
import (
	"fgstgu/go/src/d/cgo/func-call/d17/qsort"
	"fmt"
	"unsafe"
)

//export goQsortCompare
func goQsortCompare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

// -> go run main.go
func main() {
	values := []int32{18, 101, 0, 1, 89, 23, 78, 2, 23, 9}
	qsort.Sort(unsafe.Pointer(&values[0]),
		len(values),
		int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(C.goQsortCompare),
	)
	fmt.Println(values)
}
