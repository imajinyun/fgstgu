package qsort

/*
#include <stdlib.h>
typedef int (*qsort_cmp_func_t)(const void *a, const void *b);
*/
import "C"
import "unsafe"

// CompareFunc type.
type CompareFunc C.qsort_cmp_func_t

// Sort func.
func Sort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}
