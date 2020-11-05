package main

import "fmt"

func main() {
	a()
	b()
	c()
	fmt.Println("The main func is exited!")
}

func a() {
	fmt.Println("🎉 This is func `a`.")
}

func b() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("b func recover, return value is %v", msg)
		}
	}()
	fmt.Println("🔥 This is func `b`")
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d ", i)
		if i == 8 {
			fmt.Println()
			panic("Panic...\n")
		}
	}
	fmt.Println()
}

func c() {
	defer func() {
		fmt.Println("Execution delay function")
		msg := recover()
		fmt.Printf("Get recover value is %v\n", msg)
	}()
	fmt.Println("🍎 This is func `c`")
}
