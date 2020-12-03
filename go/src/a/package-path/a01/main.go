package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

/**
 * -> go run main.go main_lib.go -name=World
 * or
 * -> cd /path/to/fgstgu/go/src/a/package-path/a01
 * -> go build -o main
 * -> ./main -name=World
 */
func main() {
	flag.Parse()
	hello(name)
}
