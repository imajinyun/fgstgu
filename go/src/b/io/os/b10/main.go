package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	filePath1 := "/etc/hosts"
	dump(filePath1)

	filePath1 = "./.editorconfig"
	dump(filePath1)
}

func dump(filePath string) {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("File type: %T\n", fileStat)
		fmt.Printf("File name: %v\n", fileStat.Name())
		fmt.Printf("File size: %v\n", fileStat.Size())
		fmt.Printf("File is dir: %v\n", fileStat.IsDir())
		fmt.Printf("File mode: %v\n", fileStat.Mode())
		fmt.Printf("File modify time: %v\n", fileStat.ModTime())
		fmt.Printf("File path is absolute: %v\n", filepath.IsAbs(filePath))
		fmt.Printf("File path base: %v\n", filepath.Base(filePath))
		fmt.Printf("Path join using `..`: %v\n", path.Join(filePath, ".."))
		fmt.Printf("Path join using `.`: %v\n", path.Join(filePath, "."))
		fmt.Printf("Path join with other: %v\n", path.Join("/usr/local/src", filePath))
	}
	fmt.Println()
}
