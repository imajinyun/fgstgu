package main

import (
	"./lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "The greeting object.")
}

// go run main.go -name=World
func main() {
	flag.Parse()
	lib.Hello(name)
}
