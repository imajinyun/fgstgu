package main

import (
	"fmt"
)

// 将 panic 转换成 error 并返回
func division(x int) (res int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%s", p)
		}
	}()

	div := 0
	res = x / div

	return
}

// go run src/a/a0010/main.go
func main() {
	_, err := division(100)
	if err != nil {
		fmt.Printf("Error: `%s`\n", err)
	}
}
