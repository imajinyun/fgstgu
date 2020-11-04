package main

import "fmt"

// IComputer interface.
type IComputer interface {
	DumpInfo() string
}

// DellComputer struct.
type DellComputer struct {
	name string
}

// AcerComputer struct.
type AcerComputer struct {
	name string
}

// DumpInfo output Dell.
func (dell DellComputer) DumpInfo() string {
	return dell.name + ""
}

// DumpInfo output Acer.
func (acer AcerComputer) DumpInfo() string {
	return acer.name
}

func main() {
	dell := DellComputer{"This is Dell brand computer."}
	acer := AcerComputer{"This is Acer brand computer."}
	fmt.Println(dell.DumpInfo())
	fmt.Println(acer.DumpInfo())

	var ic IComputer
	ic = dell
	fmt.Printf("%T, %v, %p\n", ic, ic, &ic)
	fmt.Println(ic.DumpInfo())

	ic = acer
	fmt.Printf("%T, %v, %p\n", ic, ic, &ic)
	fmt.Println(ic.DumpInfo())
}
