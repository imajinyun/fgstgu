// [Calling C code from go](https://karthikkaranth.me/blog/calling-c-code-from-go/)
// [Calling Go Functions from Other Languages using C Shared Libraries](https://github.com/vladimirvivien/go-cshared-examples)

package main

/*
#include <stdio.h>
static void Hello(const char* s) {
	puts(s);
}
*/
import "C"

func main() {
	C.Hello(C.CString("🎉 Hello World, Hello C, Hello Go\n"))
}
