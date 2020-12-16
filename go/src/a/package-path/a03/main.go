package main

import (
	"fgstgu/go/src/a/package-path/a03/lib"
	// use of internal package a/package-path/a03/lib/internal not allowed
	// in "a/package-psath/a03/lib/internal"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Guest", "Hello World")
}

/**
 * -> cd src
 * -> export GOPATH=/Users/username/Codes/fgstgu/go:$GOPATH
 * -> go install a/package-path/a03/lib
 * -> go run a/package-path/a03/main.go -name=World
 */
func main() {
	flag.Parse()
	lib.Hello(name)
	// in.Hello(name)
}
