package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "/etc/hosts"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bs := make([]byte, 1024*8, 1024*8)
		n := -1
		for {
			n, err = file.Read(bs)
			if n == 0 || err == io.EOF {
				fmt.Println("✔️ The file has been read to the end!")
				break
			}
			fmt.Printf(string(bs[:n]))
		}
	}
	file.Close()
}
