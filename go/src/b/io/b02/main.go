package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "/tmp/test"
	err := os.Mkdir(filename, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Directory %s created successfully!\n", filename)
	}

	filename = "/tmp/a/b/c/d/test"
	err = os.MkdirAll(filename, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Directory %s created successfully!\n", filename)
	}
}
