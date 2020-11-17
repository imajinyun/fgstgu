package main

import "os"

func main() {
	exitChan := make(chan int)
	go server("127.0.0.1:8007", exitChan)
	code := <-exitChan
	os.Exit(code)
}
