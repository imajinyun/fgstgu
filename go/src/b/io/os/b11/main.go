package main

import (
	"fmt"
	"os"
)

func main() {
	dirname := "/tmp/test"
	err := os.Mkdir(dirname, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Directory %s created successfully!\n", dirname)
	}

	dirname = "/tmp/a/b/c/d/test"
	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Directory %s created successfully!\n", dirname)
	}

	filename := "/tmp/test/acme.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v %s created successfully!\n", file, filename)
	}
}
