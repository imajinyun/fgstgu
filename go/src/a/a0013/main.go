package main

import (
	"fmt"
	"strings"
)

// 函数变量
func main() {
	a := strToCase("ABCDEFGHIJKLMNOPQRSTUVWXYZ", toLower)
	fmt.Printf("String to lower: %s\n", a)

	b := strToCase2("abcdefghijklmnopqrstuvwxyz", toUpper)
	fmt.Printf("String to upper: %s\n", b)
}

func toLower(str string) string {
	result := ""
	for _, value := range str {
		result += strings.ToLower(string(value))
	}
	return result
}

func toUpper(str string) string {
	result := ""
	for _, value := range str {
		result += strings.ToUpper(string(value))
	}
	return result
}

func strToCase(str string, fn func(string) string) string {
	fmt.Printf("fn type: %T\n", fn)
	return fn(str)
}

type caseFunc func(string) string

func strToCase2(str string, fn caseFunc) string {
	fmt.Printf("fn type: %T\n", fn)
	return fn(str)
}
