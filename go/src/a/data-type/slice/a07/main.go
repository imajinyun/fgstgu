package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	slice1 := []*int{}
	for _, v := range array {
		slice1 = append(slice1, &v)
	}
	for _, v := range slice1 {
		fmt.Printf("%v ", *v)
	}
	fmt.Println()

	slice2 := []*int{}
	for i, _ := range array {
		slice2 = append(slice2, &array[i])
	}
	for _, v := range slice2 {
		fmt.Printf("%v ", *v)
	}
	fmt.Println()
}
