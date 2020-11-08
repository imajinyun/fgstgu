package main

import (
	"encoding/json"
	"fmt"
)

// Point struct.
type Point struct{ X, Y int }

// Circle struct.
type Circle struct {
	Point
	Radius int
}

func main() {
	data := `{"X":50,"Y":50,"Radius":25}`
	circle := Circle{}
	json.Unmarshal([]byte(data), &circle)
	fmt.Println(circle)
}
