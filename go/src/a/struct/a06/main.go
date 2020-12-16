package main

import (
	"fgstgu/go/src/a/struct/a06/base"
	_ "fgstgu/go/src/a/struct/a06/cls1"
	_ "fgstgu/go/src/a/struct/a06/cls2"
)

func main() {
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}
