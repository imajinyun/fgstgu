package main

import "fmt"

// Outer inferface.
type Outer interface {
	Output(s string)
}

type aType string

func (a aType) Output(s string) {
	fmt.Printf("%s method output\n", s)
}

type bType string

func (b bType) Output(s string) {
	fmt.Printf("%s method output\n", s)
}

func myPrint(o Outer, s string) {
	o.Output(s)
}

func main() {
	var out Outer
	out = aType("aType")
	myPrint(out, "Hello")
	out = bType("bType")
	myPrint(out, "World")
}
