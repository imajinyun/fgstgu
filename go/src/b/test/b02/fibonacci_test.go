// -> go test -bench=.
// -> go test -bench='Sequence$' -cpu=2,4,8,16,32,64,128,256,512,1024,2048,4096
// -> go test -bench='Sequence$' -benchtime=5s .
// -> go test -bench='Sequence$' -benchtime=1000x .
// -> go test -bench='Sequence$' -benchtime=1000x -count=5 .

package b02

import "testing"

func fibonacciSequence(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacciSequence(n-1) + fibonacciSequence(n-2)
}

func BenchmarkFibonacciSequence(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacciSequence(30)
	}
}
