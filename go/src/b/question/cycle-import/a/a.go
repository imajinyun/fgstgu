package a

import (
	"fgstgu/go/src/b/question/cycle-import/b"
	"fmt"
)

type A struct{}

func (a A) PrntB() {
	fmt.Println(a)
}

func NewB() *A {
	return new(A)
}

func RequireB() {
	o := b.NewB()
	o.PrintB()
}
