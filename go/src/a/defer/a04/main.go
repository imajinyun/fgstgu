package main

import "fmt"

func main() {
	a := 10
	b := 90
	defer add(a, b, true)

	a = 111
	b = 999
	defer add(a, b, false)
}

func add(a, b int, flag bool) {
	if flag {
		fmt.Printf("The function that is delayed execution is `add`, result is %v\n", a+b)
	} else {
		fmt.Printf("The function that is not deferred execution is `add`, result is %v\n", a+b)
	}
}
