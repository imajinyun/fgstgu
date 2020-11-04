package main

import "fmt"

func main() {
	var a = [3][4]int{
		{10, 11, 12, 13},
		{20, 21, 22, 23},
		{30, 31, 32, 33},
	}

	m, n := len(a), len(a[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	fmt.Println()

	b := [...]string{"Tencent", "Alibaba", "Batedance"}
	c := b
	c[0] = "Meituan"
	fmt.Println(b)
	fmt.Println(c)
}
