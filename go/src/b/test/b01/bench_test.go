package b01

import "testing"

/**
 * -> go test -v -bench=. bench_test.go
 * -> go test -v -bench=. -benchtime=5s bench_test.go
 */
func BenchmarkSimple(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func BenchmarkAddTimerControl(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
