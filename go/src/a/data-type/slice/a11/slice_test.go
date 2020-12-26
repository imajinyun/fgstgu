package a11

import "testing"

func TestSlice(t *testing.T) {
	var n int = 5
	a, b := make([]int, n), make([]int, n)
	t.Log(a)
	t.Log(b)

	for i := 0; i < n; i++ {
		a[i]++
		b[i]++
	}
	t.Log(a)
	t.Log(b)
}
