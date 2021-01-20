package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "/tmp/test.txt"
	dstFile := "/tmp/demo.txt"
	total, err := copyFile(srcFile, dstFile)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Copy successfully, byte is %v\n", total)
	}
}

func copyFile(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
