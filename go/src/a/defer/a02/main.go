package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	getLargest(slice)
}

func finished() {
	fmt.Println("Finished!")
}

func getLargest(s []int) {
	defer finished()
	fmt.Println("Find largest value in slice `s`...")
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	fmt.Printf("Largest value is %v\n", max)
}
