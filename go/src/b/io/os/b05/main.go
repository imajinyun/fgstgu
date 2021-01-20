package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/tmp/test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	n, err := file.Write([]byte("🍎 Hello World, Hello Go!\n"))
	if err != nil {
		fmt.Printf("✖️ File write failure, reason is %v\n", err.Error())
	} else {
		fmt.Printf("✔️ File write successfully, bytes is %v\n", n)
	}
	n, err = file.WriteString("🎉 我和我的祖国\n")
	if err != nil {
		fmt.Printf("✖️ File write failure, reason is %v\n", err.Error())
	} else {
		fmt.Printf("✔️ File write successfully, byte is %v!\n", n)
	}
}
