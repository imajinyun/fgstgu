package main

import (
	"a/package-path/a03/lib"
	// use of internal package a/package-path/a03/lib/internal not allowed
	// in "a/package-path/a03/lib/internal"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World")
}

func main() {
	flag.Parse()
	lib.Hello()
	// in.Hello(name)
}
