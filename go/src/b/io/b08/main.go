package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "/etc/hosts"
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	fmt.Printf("%T\n", reader)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err != nil {
			fmt.Println("Read finished!")
			break
		}
	}
	file.Close()
}
