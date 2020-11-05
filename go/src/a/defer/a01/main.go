package main

import "fmt"

func main() {
	defer a()
	b()
	defer c()
	fmt.Println("main func is over...")
}

func a() {
	fmt.Println("This is a func.")
}

func b() {
	fmt.Println("This is b func.")
}

func c() {
	fmt.Println("This is c func.")
}
