package main

import "fmt"

// Teacher struct.
type Teacher struct {
	name    string
	age     int
	married bool
	gender  int8
}

// 函数传结构体
func main() {
	t := Teacher{"T1", 36, true, 1}
	fmt.Printf("1. struct variable t memory address: %p, value: %v\n", &t, t)
	change(t)
	fmt.Printf("3. struct variable t memory address: %p, value: %v\n", &t, t)
	update(&t)
	fmt.Printf("5. struct variable t memory address: %p, value: %v\n", &t, t)
}

func change(t Teacher) {
	fmt.Printf("2. struct variable t memory address: %p, value: %v\n", &t, t)
	t.name = "T2"
}

func update(t *Teacher) {
	fmt.Printf("4. struct variable t memory address: %p, value: %v\n", &t, *t)
	(*t).name = "T3"
}
