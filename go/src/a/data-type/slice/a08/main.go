package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c"}
	s := a[:]
	fmt.Printf("a addr is %p, s addr is %p\n", &a, &s)
	fmt.Printf("a=%v, s=%v\n", a, s)

	a[2] = "e"
	fmt.Printf("a=%v, s=%v\n", a, s)

	s[0] = "d"
	fmt.Printf("a=%v, s=%v\n", a, s)
	fmt.Printf("a type is %T, b type is %T\n", a, s)

	t := []string{"m", "n"}
	s = append(s, t...)
	fmt.Printf("a=%v, s=%v\n", a, s)
	fmt.Printf("a addr is %p, s addr is %p\n", &a, &s)

	a[0] = "a"
	fmt.Printf("a=%v, s=%v\n", a, s)

	s[0] = "d"
	fmt.Printf("a=%v, s=%v\n", a, s)

	p := []string{"r", "s", "t", "u", "v", "w", "x", "y", "z"}
	s = append(s, p...)
	fmt.Printf("a=%v, s=%v\n", a, s)
	fmt.Printf("a addr is %p, s addr is %p\n", &a, &s)
}
