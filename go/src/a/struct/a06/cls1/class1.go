package cls1

import (
	"a/struct/a06/base"
	"fmt"
)

// Class1 struct.
type Class1 struct{}

// Do func.
func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
