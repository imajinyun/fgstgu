package main

import (
	"fmt"
	"reflect"
)

func main() {
	type dog struct {
		legCount1 int
		LegCount2 int
	}
	valueOfDog := reflect.ValueOf(dog{})
	legCount1 := valueOfDog.FieldByName("legCount1")
	fmt.Printf("%T, %v\n", legCount1, legCount1)

	// panic: reflect: reflect.Value.SetInt using value obtained using unexported field
	// legCount.SetInt(4)

	valueOfDog = reflect.ValueOf(&dog{})
	valueOfDog = valueOfDog.Elem()
	LegCount2 := valueOfDog.FieldByName("LegCount2")
	LegCount2.SetInt(4)
	fmt.Printf("%T, %v\n", LegCount2, LegCount2)
	fmt.Println(LegCount2.Int())
}
