package a05

import "testing"

const numOfElems = 1000

// Content struct.
type Content struct {
	Detail [10000]int
}

func withValue(arr [numOfElems]Content) int {
	return 0
}

func withReference(arr *[numOfElems]Content) int {
	return 0
}

func TestFunc(t *testing.T) {
	var arr [numOfElems]Content
	withValue(arr)
	withReference(&arr)
}

// -> GODEBUG=gctrace=1 go test -bench=BenchmarkArrayWithValue
func BenchmarkArrayWithValue(b *testing.B) {
	var arr [numOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

// -> GODEBUG=gctrace=1 go test -bench=BenchmarkArrayWithReference
func BenchmarkArrayWithReference(b *testing.B) {
	var arr [numOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}
