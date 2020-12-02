package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 1 << 30
	b := math.Pow(2, 30)
	fmt.Println(a, b)
}
