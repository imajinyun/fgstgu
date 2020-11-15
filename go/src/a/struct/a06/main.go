package main

import (
	"a/struct/a06/base"
	_ "a/struct/a06/cls1"
	_ "a/struct/a06/cls2"
)

func main() {
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}
