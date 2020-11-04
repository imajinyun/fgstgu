package main

import (
	"fmt"
	"math"
)

// Shape interface.
type Shape interface {
	perimeter() float64
	area() float64
}

// Rectangle struct.
type Rectangle struct {
	a, b float64
}

// Triangle struct.
type Triangle struct {
	a, b, c float64
}

// Circle struct.
type Circle struct {
	radius float64
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.a + r.b)
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getType(s Shape) {
	if instance, ok := s.(Rectangle); ok {
		fmt.Printf("Rectangle(weight=%.2f, height=%.2f)\n", instance.a, instance.b)
	} else if instance, ok := s.(Triangle); ok {
		fmt.Printf("Triangle(edge=%.2f, edge=%.2f, edge=%.2f)\n",
			instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("Circle(radius=%.2f)\n", instance.radius)
	}
}

func getType2(s Shape) {
	switch instance := s.(type) {
	case Rectangle:
		fmt.Printf("Rectangle(width=%.2f, height=%.2f)\n", instance.a, instance.b)
	case Triangle:
		fmt.Printf("Triangle(edge=%.2f, edge=%.2f, edge=%.2f)\n",
			instance.a, instance.b, instance.c)
	case Circle:
		fmt.Printf("Circle(radius=%.2f)\n", instance.radius)
	}
}

func getResult(s Shape) {
	getType2(s)
	fmt.Printf("Perimeter: %.2f, Area: %.2f\n", s.perimeter(), s.area())
	fmt.Printf("%T, %v, %p\n\n", s, s, &s)
}

// Realize the system interface and change the print result.
func (t Triangle) String() string {
	return fmt.Sprintf("Triangle properties: %.2f, %.2f, %.2f", t.a, t.b, t.c)
}

func main() {
	var s Shape
	s = Rectangle{3, 5}
	getResult(s)

	s = Triangle{3, 4, 5}
	getResult(s)

	s = Circle{6}
	getResult(s)
}
