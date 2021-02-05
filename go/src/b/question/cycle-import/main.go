package main

import "fgstgu/go/src/b/question/cycle-import/a"

// -> go run ./src/b/question/cycle-import/main.go
// package command-line-arguments
//         imports fgstgu/go/src/b/question/cycle-import/a
//         imports fgstgu/go/src/b/question/cycle-import/b
//         imports fgstgu/go/src/b/question/cycle-import/a: import cycle not allowed
func main() {
	o := a.NewA()
	o.PrintA()
}
