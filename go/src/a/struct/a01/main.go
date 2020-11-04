package main

import "fmt"

// Teacher struct.
type Teacher struct {
	name   string
	age    uint8
	gender byte
}

func main() {
	var a Teacher
	fmt.Println(a)
	fmt.Printf("a -> %T, %v, %q, %p\n\n", a, a, a, &a)
	a.name = "Storm"
	a.age = 18
	a.gender = 1
	fmt.Println(a)
	fmt.Printf("a -> %T, %v, %q, %p\n\n", a, a, a, &a)

	b := Teacher{}
	b.name = "Jack"
	b.age = 28
	b.gender = 1
	fmt.Println(b)
	fmt.Printf("b -> %T, %v, %q, %p\n\n", b, b, b, &b)

	c := Teacher{name: "Tom", age: 38, gender: 1}
	fmt.Println(c)
	fmt.Printf("c -> %T, %v, %q, %p\n\n", c, c, c, &c)

	c = Teacher{name: "Acme", age: 48}
	fmt.Println(c)
	fmt.Printf("c -> %T, %v, %q, %p\n\n", c, c, c, &c)

	d := Teacher{name: "John", age: 58, gender: 1}
	fmt.Println(d)
	fmt.Printf("d -> %T, %v, %q, %p\n\n", d, d, d, &d)
}
