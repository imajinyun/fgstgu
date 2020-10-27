package main

import (
	"a/a0002/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
