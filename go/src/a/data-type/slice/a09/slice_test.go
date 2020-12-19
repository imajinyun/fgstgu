package a09

import (
	"math/rand"
	"time"
	"testing"
)

func BenchmarkGenWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genWithCap(1000000)
	}
}

func BenchmarkGenWithoutCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genWithoutCap(1000000)
	}
}

func genWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func genWithoutCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
