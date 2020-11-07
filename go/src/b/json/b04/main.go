package main

import (
	"encoding/json"
	"fmt"
)

// Point struct
type Point struct {
	X, Y int
}

// Circle struct.
type Circle struct {
	Point
	Radius int
}

func main() {
	circle := Circle{Point{50, 50}, 25}
	if data, err := json.MarshalIndent(circle, "", "  "); err == nil {
		fmt.Printf("%s\n", data)
	}
}
