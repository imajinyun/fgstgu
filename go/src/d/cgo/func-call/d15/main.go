package main

/*
#include <errno.h>

static void noreturn() {}
static int div(int a, int b) {
	if (b == 0) {
		errno = EINVAL;
		return 0;
	}
	return a / b;
}
*/
import "C"
import "fmt"

func main() {
	v0, err0 := C.div(20, 1)
	fmt.Println(v0, err0)

	v1, err1 := C.div(10, 0)
	fmt.Println(v1, err1)

	_, err := C.noreturn()
	fmt.Println(err)

	v, _ := C.noreturn()
	fmt.Printf("%#v, %v\n", v, v)
	fmt.Println(C.noreturn())
}
