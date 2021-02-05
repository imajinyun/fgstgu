package main

import "fmt"

func main() {
	defer fmt.Println("panic: in main func")
	defer func() {
		defer func() {
			panic("panic #3")
		}()
		panic("panic #2")
	}()
	panic("panic #1")
}
