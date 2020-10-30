package main

import "fmt"

// 闭包实现累加和
func main() {
	fn := sumer()
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\t", i)
		fmt.Println(fn(i))
	}
}

func sumer() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum1 = %2d\t", sum)
		sum += x
		fmt.Printf("sum2 = %2d\t", sum)
		return sum
	}
}
