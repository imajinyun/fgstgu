package main

import "fmt"

// AA struct.
type AA struct{ name string }

// BB struct.
type BB struct{ name string }

// CC struct.
type CC struct {
	AA
	BB
}

func main() {
	c := &CC{}
	c.AA.name = "This is name of AA struct"
	c.BB.name = "This is name of BB struct"
	fmt.Printf("%+v\n", c)
}
