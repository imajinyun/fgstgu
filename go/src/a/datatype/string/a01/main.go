package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "你好世界，你好 Go！"
	fmt.Printf("s length: %v\n", len(s))
	fmt.Println(s)

	count := 0
	for i, ch := range s {
		fmt.Printf("%d:%c ", i, ch)
		count++
	}
	fmt.Println()
	fmt.Println()

	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x\n", i, ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("%d:%x\n", i, ch)
	}
	fmt.Println()

	fmt.Printf("s length using len func: %v\n", len(s))
	fmt.Printf("s length using utf8.RuneCountInString func: %v\n", utf8.RuneCountInString(s))
}
