package a06

// -> GOSSAFUNC=newSlice go build src/a/data-type/slice/a06/slice.go
func newSlice() []int {
	array := [3]int{1, 2, 3}
	slice := array[0:1]
	return slice
}
