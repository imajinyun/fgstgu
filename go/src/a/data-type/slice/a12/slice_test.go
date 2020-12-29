package a12

import "testing"

func TestSlice(t *testing.T) {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}
	t.Log(table)
}

func TestMulitSlice(t *testing.T) {
	m, n := 3, 4
	table := make([][]bool, m+1)
	for i := range table {
		table[i] = make([]bool, n+1)
	}
	t.Logf("\n%v", table)
	for i := 0; i <= m; i++ {
		table[i][0] = true
	}
	t.Logf("\n%v", table)

	for j := 0; j <= n; j++ {
		table[0][j] = true
	}
	t.Logf("\n%v", table)
}
