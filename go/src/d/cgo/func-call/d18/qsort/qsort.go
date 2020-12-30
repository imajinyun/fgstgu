package qsort

/*
#include <stdlib.h>
typedef int (*qsort_cmp_func_t) (const void *a, const void *b);
extern int _cgo_qsort_compare(void *a, void *b);
*/
import "C"
import (
	"sync"
	"unsafe"
)

var goQsortCompareInfo struct {
	fn func(a, b unsafe.Pointer) int
	sync.Mutex
}

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	return C.int(goQsortCompareInfo.fn(a, b))
}

// Sort func.
func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int) {
	goQsortCompareInfo.Lock()
	defer goQsortCompareInfo.Unlock()
	goQsortCompareInfo.fn = cmp
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(C._cgo_qsort_compare))
}
