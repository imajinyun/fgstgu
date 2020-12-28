package main

/*
#include <stdio.h>
void dumpint(int v) {
	printf("dumpint: %d\n", v);
}
*/
import "C"

// -> go run main.go
// 通过虚拟的 C 包导入的 C 语言符号并不需要以大写字母开头，它们不受 Go 语言的导出规则约束
func main() {
	v := 42
	C.dumpint(C.int(v))
}
