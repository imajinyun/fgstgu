package main

import "fmt"

func main() {
	s := "我和我的祖国。"
	fmt.Printf("Original string: %s\n", s)
	fmt.Print("Flipped string: ")
	reverseString(s)
}

func reverseString(s string) {
	for _, v := range []rune(s) {
		defer fmt.Printf("%c", v)
	}
}
