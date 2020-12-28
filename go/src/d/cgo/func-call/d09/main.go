package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1
#if defined(CGO_OS_WINDOWS)
    const char* os = "Windows";
#elif defined(CGO_OS_DARWIN)
	const char* os = "Darwin";
#elif defined(CGO_OS_LINUX)
    const char* os = "Linux";
#else
	const char* os = "Unknow OS";
#endif
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.os))
}
