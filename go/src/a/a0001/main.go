package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

// go run main.go main_lib.go -name=World
func main() {
	flag.Parse()
	hello(name)
}
