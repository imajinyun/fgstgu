package main

import "fmt"

// Person struct.
type Person struct {
	firstName, lastName string
}

func (p Person) fullName() {
	fmt.Printf("%s·%s\n", p.firstName, p.lastName)
}

func main() {
	p := Person{"Peter", "Chen"}
	defer p.fullName()
	fmt.Print("🎉 Welcome, ")
}
