package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 0)
	s = append(s, 1)
	fmt.Printf("%T, %v", s, s)
}
