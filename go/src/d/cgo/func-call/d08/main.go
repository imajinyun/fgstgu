package main

//static const char* str = "Hello World";
import "C"
import "fgstgu/go/src/d/cgo/func-call/d08/cgohelper"

// -> go run main.go 这段不能正常工作。
// 原因：
// 在 Go 语言中方法是依附于类型存在的，不同 Go 包中引入的虚拟的 C 包的
// 类型是不同的（main.C 与 cgohelper.C 类型不同），这导致从它们延伸
// 出来的 Go 类型也是不同的类型（*main.C.char 与 *cgo_helper.C.char 类型不同），
// 这最终导致了下面代码不能正常工作。
func main() {
	cgohelper.PrintCString(C.str)
}
