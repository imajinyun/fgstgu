package main

import "fmt"

func main() {
	students := map[int]string{
		1001: "Jack",
		1002: "Tom",
		1003: "Amy",
		1004: "Peter",
	}
	fmt.Println(students)
	newStudents := students
	newStudents[1004] = "Tony"
	fmt.Printf("elements: %v, length: %v\n", students, len(students))
	fmt.Printf("elements: %v, length: %v\n", newStudents, len(newStudents))
}
