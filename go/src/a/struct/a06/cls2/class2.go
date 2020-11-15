package cls2

import (
	"a/struct/a06/base"
	"fmt"
)

// Class2 struct.
type Class2 struct{}

// Do func.
func (c *Class2) Do() {
	fmt.Println("Class2")
}

func init() {
	base.Register("Class2", func() base.Class {
		return new(Class2)
	})
}
