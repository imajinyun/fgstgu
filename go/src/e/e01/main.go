package main

import (
	"github.com/labstack/gommon/log"
	"e/e01/cmd"
)

/**
 * -> export GOPATH=/Users/tal/Codes/fgstgu/go:$GOPATH
 * -> go install e/e01
 * -> go run e/e01/main.go help word
 * -> go run e/e01/main.go word -s=hello_world -m=1
 */
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute: %v", err)
	}
}
