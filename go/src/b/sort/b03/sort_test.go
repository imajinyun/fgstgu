package b03

import (
	"reflect"
	"sort"
	"testing"
	"unsafe"
)

func TestSortFloat64Value1(t *testing.T) {
	var a = []float64{4, 2, 78, 1, 128, 10, 88, 23}
	sortFloat64Value1(a)
	t.Log(a)
}

func TestSortFloat64Value2(t *testing.T) {
	var a = []float64{4, 2, 78, 1, 128, 10, 88, 23}
	sortFloat64Value2(a)
	t.Log(a)
}

func sortFloat64Value1(a []float64) {
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
}

func sortFloat64Value2(a []float64) {
	var b []int
	ahdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bhdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	*bhdr = *ahdr
	sort.Ints(b)
}
