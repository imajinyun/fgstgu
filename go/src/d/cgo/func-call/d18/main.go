package main

import (
	"fgstgu/go/src/d/cgo/func-call/d18/qsort"
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	values := []int{100, 88, 1, 17, 38, 72, 3, 89, 7}
	goSliceSort(values)
	fmt.Println(values)

	values = []int{800, 37, 169, 15, 33, 29, 124, 2, 13}
	qsort.Sort(unsafe.Pointer(&values[0]),
		len(values),
		int(unsafe.Sizeof(values[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)
	fmt.Println(values)
}

func goSliceSort(values []int) {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
}
