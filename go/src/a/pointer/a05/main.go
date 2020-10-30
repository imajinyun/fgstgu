package main

import "fmt"

// LENGTH represents the length of the array.
const LENGTH int = 5

// 指针数组
func main() {
	a := [LENGTH]string{"a", "b", "c", "d", "e"}
	i := 0

	var ptr [LENGTH]*string
	fmt.Printf("ptr type: %T, value: %v\n", ptr, ptr)

	for i = 0; i < LENGTH; i++ {
		ptr[i] = &a[i]
	}
	fmt.Printf("ptr type: %T, value: %v\n", ptr, ptr)
	fmt.Println(*ptr[0], ptr[0])

	for i = 0; i < LENGTH; i++ {
		fmt.Printf("a[%d] = %s\n", i, *ptr[i])
	}
}
