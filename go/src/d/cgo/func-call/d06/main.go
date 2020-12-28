package main

//void Hello(_GoString_ s);
import "C"
import "fmt"

func main() {
	C.Hello("🎉 Hello World, Hello C, Hello Go\n")
}

//export Hello
func Hello(s string) {
	fmt.Print(s)
}
