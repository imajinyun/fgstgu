package main

import "fmt"

// Animal interface.
type Animal interface {
}

// Cat struct.
type Cat struct {
	name string
	age  int
}

// Person struct.
type Person struct {
	name string
	age  int
}

func main() {
	var c1 Animal = Cat{"Dog1", 1}
	var c2 Animal = Cat{"Dog2", 2}
	var p1 Animal = Person{"Zhangsan", 18}
	var p2 Animal = Person{"Steven Tom", 28}
	var a1 Animal = "This is only a test, dont't worry!"
	var a2 Animal = 1000
	var a3 Animal = 3.1415926
	var a4 map[string]interface{} = make(map[string]interface{})
	var a5 map[string]interface{} = map[string]interface{}{}
	var a6 []interface{} = make([]interface{}, 0)
	var a7 []interface{} = []interface{}{}

	res := []Animal{c1, c2, p1, p2, a1, a2, a3, a4, a5, a6, a7}
	for _, v := range res {
		dump(v)
	}

	// map value is any type.
	m := make(map[string]interface{})
	m["name"] = "Tony"
	m["age"] = 39
	m["height"] = 1.78
	m["gender"] = true
	dump(m)
	fmt.Println()

	n := make([]interface{}, 0, 5)
	n = append(n, c1, p1, a1, a4, a6)
	dump(n)
}

func dump(v Animal) {
	fmt.Printf("%T, %v\n", v, v)
}
