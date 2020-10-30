package main

import (
	"a/package-path/a02/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

// go run main.go -name=World
func main() {
	flag.Parse()
	lib.Hello(name)
}
