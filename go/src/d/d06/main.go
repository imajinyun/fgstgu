package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

// -> go run main.go > cpu.pprof
// -> go run main.go
// -> go tool pprof -http=:8080 cpu.pprof
func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := genWithoutCap(n)
		bubbleSort(nums)
		n *= 10
	}
}

func genWithoutCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
