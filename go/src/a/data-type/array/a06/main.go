package main

import "fmt"

func main() {
	var a [3]int = [3]int{}
	var b = [3]int{}
	c := [3]int{}
	d := [3]int{1, 2, 3}
	e := [...]int{2: 3, 1: 2}
	f := [...]int{1, 2, 4: 5, 6}
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
