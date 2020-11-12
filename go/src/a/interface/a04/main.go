package main

import "fmt"

// MyWriter interface
type MyWriter interface {
	Write(string) string
}

type file struct{}

func (f *file) Write(s string) string {
	fmt.Println("Writer file content:\n", s)
	return s
}

func main() {
	f := new(file)
	var w MyWriter
	w = f
	fmt.Println("Return writer content\n", w.Write("🎉 Hello World, Hello Go!\n"))
}
