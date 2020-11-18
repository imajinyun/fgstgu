package main

import (
	"fmt"
	"reflect"
)

// Enum type.
type Enum int

// Zero Enum.
const (
	Zero Enum = 0
)

func main() {
	type Dog struct{}

	typeOfDog := reflect.TypeOf(Dog{})
	fmt.Printf("Name: %v, Kind: %v\n", typeOfDog.Name(), typeOfDog.Kind())

	typeOfZero := reflect.TypeOf(Zero)
	fmt.Printf("Name: %v, Kind: %v\n", typeOfZero.Name(), typeOfZero.Kind())
}
