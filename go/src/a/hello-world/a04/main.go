package main

// -> GOOS=linux GOARCH=amd64 go tool compile -S src/a/hello-world/a04/main.go
func main() {
	str := "hello"
	println([]byte(str))
}
