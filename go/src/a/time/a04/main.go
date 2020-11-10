package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := multiValues()
	fmt.Printf("%d, %d\n", a, b)
	fn := timeSpent(execSlowFunc)
	fmt.Printf("%v\n", fn(9999))
}

func multiValues() (int, int) {
	return rand.Int(), rand.Int()
}

func execSlowFunc(x int) int {
	time.Sleep(3 * time.Second)
	return x + 1
}

func timeSpent(fn func(x int) int) func(n int) int {
	return func(x int) int {
		start := time.Now()
		ret := fn(x)
		fmt.Println("Time spent:", time.Since(start).Seconds())
		return ret
	}
}
