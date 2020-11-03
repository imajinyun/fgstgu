package main

import "fmt"

func main() {
	// Remove all elements.
	// To remove all elements, simply set the slice to nil.
	// This will release the underlying array to the gargage collector (assuming
	// there are no other references).
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Printf("a = %v, len = %d, cap = %d\n", a, len(a), cap(a))

	// Keep allocated memory.
	// To keep the underlying array, slice the slice to zero length.
	b := []string{"A", "B", "C", "D", "E"}
	b = b[:0]
	fmt.Printf("b = %v, len = %d, cap = %d\n", b, len(b), cap(b))

	// Empty slice vs. nil slice
	var c []int = nil
	fmt.Printf("c = %v, len = %d, cap = %d\n", c, len(c), cap(c))

	var d []int = nil
	var e []int = make([]int, 0)
	fmt.Println("d is nil: ", d == nil)
	fmt.Println("e is nil: ", e == nil)
	fmt.Printf("d = %#v, e = %#v\n", d, e)
}
