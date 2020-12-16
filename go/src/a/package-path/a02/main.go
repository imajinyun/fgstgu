package main

import (
	"fgstgu/go/src/a/package-path/a02/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World!")
}

/**
 * -> cd src
 * -> export GOPATH=/Users/tal/Codes/fgstgu/go:$GOPAT
 * -> go install a/package-path/a02/lib
 * -> go run a/package-path/a02/main.go -name=World
 */
func main() {
	flag.Parse()
	lib.Hello(name)
}
