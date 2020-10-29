package main

import "fmt"

func main()  {
	var a int = 10
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a type: %T\n", a)
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("c type: %T\n", c)

	ptr = &a
	fmt.Printf("a = %d\n", a)
	fmt.Printf("ptr = %d\n", *ptr)
}
