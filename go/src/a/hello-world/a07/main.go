package main

// -> go build main.go
// -> go tool objdump -S -s "main" ./main
func main() {
	example(make([]string, 2, 4), "hello", 10)
}

func example(slice []string, str string, i int) {
	panic("Want stack trace")
}
