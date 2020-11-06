package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := "abcdefghijkl🙏mnopqrstuvwxyz 666 China Go!"
	ioReader := strings.NewReader(str)
	reader := bufio.NewReader(ioReader)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "q!" {
			break
		}
	}
}
