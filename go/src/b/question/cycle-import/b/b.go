package b

import (
	"fgstgu/go/src/b/question/cycle-import/a"
	"fmt"
)

type B struct{}

func (b B) PrntB() {
	fmt.Println(b)
}

func NewB() *B {
	return new(B)
}

func RequireA() {
	o := a.NewA()
	o.PrintA()
}
