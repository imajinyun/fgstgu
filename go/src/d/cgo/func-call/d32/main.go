package main

import "fmt"

// -> go build -o py3-config py3-config.go
// -> PKG_CONFIG=./py3-config go build -buildmode=c-shared -o gopkg.so main.go
func main() {
	fmt.Println("🎉 Hello pkg-config")
}
