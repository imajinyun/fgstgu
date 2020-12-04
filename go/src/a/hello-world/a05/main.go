package main

// -> GOOS=linux GOARCH=amd64 go tool compile -S -N -l src/a/hello-world/a04/main.go
func main() {
	myFunc(99, 11)
}

func myFunc(x, y int) (int, int) {
	return x + y, x - y
}
