package qsort

/*
#include <stdlib.h>
typedef int (*qsort_cmp_func_t) (const void *a, const void *b);
extern int _cgo_qsort_compare(void *a, void *b);
*/
import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var goQsortCompareInfo struct {
	base     unsafe.Pointer
	elemnum  int
	elemsize int
	less     func(a, b int) bool
	sync.Mutex
}

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	var (
		base     = uintptr(goQsortCompareInfo.base)
		elemsize = uintptr(goQsortCompareInfo.elemsize)
	)
	i := int((uintptr(a) - base) / elemsize)
	j := int((uintptr(b) - base) / elemsize)
	switch {
	case goQsortCompareInfo.less(i, j): // v[i] < v[j]
		return -1
	case goQsortCompareInfo.less(j, i): // v[i] > v[j]
		return +1
	default:
		return 0
	}
}

// Slice func.
func Slice(slice interface{}, less func(a, b int) bool) {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if rv.Len() == 0 {
		return
	}
	goQsortCompareInfo.Lock()
	defer goQsortCompareInfo.Unlock()
	defer func() {
		goQsortCompareInfo.base = nil
		goQsortCompareInfo.elemnum = 0
		goQsortCompareInfo.elemsize = 0
		goQsortCompareInfo.less = nil
	}()
	goQsortCompareInfo.base = unsafe.Pointer(rv.Index(0).Addr().Pointer())
	goQsortCompareInfo.elemnum = rv.Len()
	goQsortCompareInfo.elemsize = int(rv.Type().Elem().Size())
	goQsortCompareInfo.less = less
	C.qsort(
		goQsortCompareInfo.base,
		C.size_t(goQsortCompareInfo.elemnum),
		C.size_t(goQsortCompareInfo.elemsize),
		C.qsort_cmp_func_t(C._cgo_qsort_compare),
	)
}
