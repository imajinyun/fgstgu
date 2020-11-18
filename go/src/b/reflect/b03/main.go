package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Cat struct{}
	ins := &Cat{}
	typeOfCat := reflect.TypeOf(ins)
	fmt.Printf("Name: %v, Kind: %v\n", typeOfCat.Name(), typeOfCat.Kind())

	typeOfCat = typeOfCat.Elem()
	fmt.Printf("Elem: %v, Name: %v, Kind: %v\n", typeOfCat, typeOfCat.Name(), typeOfCat.Kind())
}
