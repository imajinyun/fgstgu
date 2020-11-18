package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string

	float32
	bool

	next *dummy
}

func main() {
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})
	fmt.Println("Num Field:", d.NumField())

	floatField := d.Field(2)
	fmt.Println("Field:", floatField)

	fmt.Println("FieldByName(\"b\").Type()", d.FieldByName("b").Type())

	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 1}).Type())
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 2}).Type())
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 3}).Type())
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 4}).Type())

	// panic: reflect: Field index out of range
	// fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 5}).Type())
}
