package main

import (
	"fmt"
	"runtime"
)

type panicContext struct {
	pContext string
}

func allowRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("Runtime error:", err)
		default:
			fmt.Println("Error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("Main func is running...")
	allowRun(func() {
		fmt.Println("Before manual downtime...")
		panic(&panicContext{"Manual trigger error"})
		// fmt.Println("After manual downtime...")
	})

	allowRun(func() {
		fmt.Println("Assignment before downtime...")
		var a *int
		*a = 1
		fmt.Println("Assignment after downtime...")
	})
	fmt.Println("Main func is over!")
}
