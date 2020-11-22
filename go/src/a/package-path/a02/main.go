package main

import (
	"a/package-path/a02/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

/**
 * -> cd /path/to/go
 * -> export GOPATH=/path/to/go/
 * -> cd ./src
 * -> go build -o main a/package-path/a02
 * -> ./main
 */
func main() {
	flag.Parse()
	lib.Hello(name)
}
