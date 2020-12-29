package main

//extern int go_qsort_compare(void *a, void *b);
import "C"
import (
	"fmt"
	"unsafe"
)

func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	values := []int32{18, 3, 0, 1, 89, 23, 78, 5, 23, 9}
	qsort.Sort(unsafe.Pointer(&values[0]),
		len(values),
		int(unsafe.Sizeof(values[0])),
		qsort.CompareFunc(C.go_qsort_compare),
	)
	fmt.Println(values)
}
