package main

import "fmt"

func main() {
	array := [3]string{"Hello", "World", "PHP"}
	slice := array[:]
	slice[2] = "Go"
	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println()

	array[0] = "Java"
	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println()

	s1 := slice[1:len(slice)]
	fmt.Println(s1)

	fmt.Println()

	s1[0] = "Python"
	fmt.Println("s1 =", s1)
	fmt.Println("array =", array)
	fmt.Println("slice =", slice)
}
