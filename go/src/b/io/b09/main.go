package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	writer()
}

func writer() {
	filename1 := "/tmp/test.txt"
	file1, _ := os.Open(filename1)
	reader := bufio.NewReader(file1)
	filename2 := "/tmp/demo.txt"
	file2, _ := os.OpenFile(filename2, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	writer := bufio.NewWriter(file2)
	for {
		bs, err := reader.ReadBytes(' ')
		writer.Write(bs)
		writer.Flush()
		if err == io.EOF {
			fmt.Println("Finished reader the file!")
			break
		}
	}
	file1.Close()
	file2.Close()
}
