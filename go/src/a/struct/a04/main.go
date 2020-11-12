package main

import "fmt"

// Wheel struct.
type Wheel struct {
	Size int
}

// Engine struct.
type Engine struct {
	Power int
	Type  string
}

// Car struct.
type Car struct {
	Wheel
	Engine
}

func main() {
	c := Car{
		Wheel:  Wheel{Size: 23},
		Engine: Engine{Power: 115, Type: "1.8T"},
	}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%v\n", c)
}
