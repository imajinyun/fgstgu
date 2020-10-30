package main

import (
	"fmt"
)

// Teacher struct.
type Teacher struct {
	name    string
	age     int
	married bool
	gender  int8
}

func main() {
	var t1 = Teacher{"T1", 35, true, 1}
	var t2 = Teacher{"T2", 28, false, 2}
	var a *Teacher = &t1
	var b *Teacher = &t2

	fmt.Printf("t1 type: %T, value: %v\n", t1, t1)
	fmt.Printf("t2 type: %T, value: %v\n", t2, t2)
	fmt.Printf("a type: %T, value: %v\n", a, a)
	fmt.Printf("b type: %T, value: %v\n", b, b)
	fmt.Printf("*a type: %T, value: %v\n", *a, *a)
	fmt.Printf("*b type: %T, value: %v\n", *b, *b)
	fmt.Println()

	fmt.Printf("[%s, %d, %v, %d]\n", t1.name, t1.age, t1.married, t1.gender)
	fmt.Printf("[%s, %d, %v, %d]\n", a.name, a.age, a.married, a.gender)
	fmt.Printf("[%s, %d, %v, %d]\n", t2.name, t2.age, t2.married, t2.gender)
	fmt.Printf("[%s, %d, %v, %d]\n", b.name, b.age, b.married, b.gender)
	fmt.Println()

	fmt.Printf("[%s, %d, %v, %d]\n", (*a).name, (*a).age, (*a).married, (*a).gender)
	fmt.Printf("[%s, %d, %v, %d]\n", (*b).name, (*b).age, (*b).married, (*b).gender)
	fmt.Printf("&a type: %T\n", &a)
	fmt.Println(&a.name, &a.age, &a.married, &a.gender)
	fmt.Printf("&b type: %T\n", &b)
	fmt.Println(&b.name, &b.age, &b.married, &b.gender)
}
