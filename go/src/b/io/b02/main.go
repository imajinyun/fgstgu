package main

import (
	"fmt"
	"os"
)

func main() {
	name := "/tmp/test"
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Directory %s created successfully!\n", name)
	}

	name = "/tmp/a/b/c/d/test"
	err = os.MkdirAll(name, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Directory %s created successfully!\n", name)
	}
}
