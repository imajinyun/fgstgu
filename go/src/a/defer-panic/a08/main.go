package main

import (
	"fmt"
	"time"
)

func main() {
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()
	time.Sleep(time.Second)
}
