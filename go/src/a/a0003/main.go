package main

import (
	"a/a0003/lib"

	// use of internal package a/a0003/lib/internal not allowed
	in "a/a0003/lib/internal"
	// "flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	// in.Hello(name)
}
