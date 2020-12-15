package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	scandir("./")
}

func scandir(dirname string) error {
	fmt.Println(dirname)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filepath := filepath.Join(dirname, file.Name())
		if file.IsDir() {
			scandir(filepath)
		} else {
			fmt.Println(filepath)
		}
	}
	return nil
}
