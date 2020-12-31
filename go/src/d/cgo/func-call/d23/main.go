package main

/*
#include <stdio.h>
extern int *getGoPtr();
static void Main() {
	int *p = getGoPtr();
	*p = 88;
	printf("%d\n", *p);
}
*/
import "C"

// GODEBUG=cgocheck=0 go run main.go
func main() {
	C.Main()
}

//export getGoPtr
func getGoPtr() *C.int {
	return new(C.int)
}
