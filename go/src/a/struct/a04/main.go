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

// RefitCar struct.
type RefitCar struct {
	Wheel
	Engine struct {
		Power int
		Type  string
	}
}

func main() {
	c1 := Car{
		Wheel:  Wheel{Size: 23},
		Engine: Engine{Power: 115, Type: "1.8T"},
	}
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%v\n", c1)
	fmt.Println()

	c2 := RefitCar{
		Wheel: Wheel{Size: 20},
		Engine: struct {
			Power int
			Type  string
		}{
			Power: 150,
			Type:  "2.0T",
		},
	}
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%v\n", c2)
}
