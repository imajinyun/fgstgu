package main

import (
	"fmt"
	"time"
)

func main() {
	go dump1()
	go dump2()
	time.Sleep(3 * time.Second)
	fmt.Println("\n🙏 Main func is over!")
}

func dump1() {
	for i := 1; i <= 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func dump2() {
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}
