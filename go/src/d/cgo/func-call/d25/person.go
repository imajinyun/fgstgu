package main

// Person struct.
type Person struct {
	name string
	age  int
}

// NewPerson returns a new Pserson.
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Set to set person name and age.
func (p *Person) Set(name string, age int) {
	p.name = name
	p.age = age
}

// Get to get person name and age.
func (p *Person) Get() (name string, age int) {
	return p.name, p.age
}
