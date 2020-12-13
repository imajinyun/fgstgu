package main

import "fmt"

// MyInterface interface.
type MyInterface interface {
	WithoutParameter()
	WithParameter(float64)
	WithReturnValue() string
}

// MyType type.
type MyType int

// WithoutParameter method for MyType.
func (m MyType) WithoutParameter() {
	fmt.Println("method with parmeter called")
}

// WithParameter method for MyType.
func (m MyType) WithParameter(f float64) {
	fmt.Println("method without parameter called")
}

// WithReturnValue method for MyType.
func (m MyType) WithReturnValue() string {
	return fmt.Sprintf("mehtod with return value called")
}

func main() {
	var value MyInterface
	value = MyType(5)
	value.WithoutParameter()
	value.WithParameter(99.9999)
	fmt.Println(value.WithReturnValue())
}
