package main

import "fmt"

func anyFunc() {
	var any interface{}

	any = 1
	fmt.Printf("%v\n", any)

	any = "Hello World"
	fmt.Printf("%v\n", any)

	any = false
	fmt.Printf("%v\n\n", any)
}

func interfaceAssignFunc() {
	var a int = 100
	var i interface{} = a

	// var b int = i -> cannot use i (type interface {}) as type int in assignment: need type assertion
	fmt.Printf("%v\n", i)

	var b int = i.(int)
	fmt.Printf("%v\n\n", b)
}

func interfaceFunc1() {
	var a interface{} = 100
	var b interface{} = "100"
	var c interface{} = 100
	fmt.Printf("%v\n", a == b)
	fmt.Printf("%v\n\n", a == c)
}

func interfaceFunc2() {
	var a interface{} = []int{10}
	var b interface{} = []int{10}
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)

	// panic: runtime error: comparing uncomparable type []int
	// fmt.Printf("%v\n", a == b)
	fmt.Println()
}

func main() {
	anyFunc()
	interfaceAssignFunc()
	interfaceFunc1()
	interfaceFunc2()
}
