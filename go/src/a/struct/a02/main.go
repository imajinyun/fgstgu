package main

import "fmt"

// Output interface.
type Output interface {
	OutputHelloWorld() string
}

// GoLanguage struct.
type GoLanguage struct {
	Name string
}

// OutputHelloWorld func.
func (p *GoLanguage) OutputHelloWorld() string {
	return "🙏 Hello World!"
}

func main() {
	a := GoLanguage{Name: "Go Programming Language"}
	fmt.Println(a.OutputHelloWorld())
	fmt.Println(a.Name)

	var b Output
	b = new(GoLanguage)
	fmt.Println(b.OutputHelloWorld())
}
