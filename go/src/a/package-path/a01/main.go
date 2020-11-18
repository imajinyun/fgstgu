package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

/**
 * -> go run main.go main_lib.go -name=World
 * or
 * -> cd /path/to/go/a/package-path/a01
 * -> go build -o main
 * -> ./main
 */
func main() {
	flag.Parse()
	hello(name)
}
