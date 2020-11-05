package main

import "fmt"

func main() {
	// string to rune.
	s := "我和我的祖国"
	for _, c := range s {
		fmt.Printf("%[1]c %[1]x\n", c)
	}
}
