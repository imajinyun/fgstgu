package a06

import "testing"

// [Array vs Slice: accessing speed](https://stackoverflow.com/questions/30525184/array-vs-slice-accessing-speed)

var gs = make([]byte, 1000) // Global slice
var ga [1000]byte           // Global array

func BenchmarkSliceGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, v := range gs {
			gs[j]++
			gs[j] = gs[j] + v + 1
			gs[j] += v
		}
	}
}

func BenchmarkArrayGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, v := range ga {
			ga[j]++
			ga[j] = ga[j] + v + 1
			ga[j] += v
		}
	}
}

func BenchmarkSliceLocal(b *testing.B) {
	var s = make([]byte, 1000)
	for i := 0; i < b.N; i++ {
		for j, v := range s {
			s[j]++
			s[j] = s[j] + v + 1
			s[j] += v
		}
	}
}

func BenchmarkArrayLocal(b *testing.B) {
	var a [1000]byte
	for i := 0; i < b.N; i++ {
		for j, v := range a {
			a[j]++
			a[j] = a[j] + v + 1
			a[j] += v
		}
	}
}
