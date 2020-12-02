package main

import "fmt"

/**
 * 1. 将 Go 语言的源代码编译成汇编语言，然后通过汇编语言分析程序具体的执行过程：
 * -> go build -gcflags -S src/a/hello-world/a01/main.go
 *
 * 2. 了解 Go 语言更详细的编译过程，可以通过下面的命令获取汇编指令的优化过程：
 * -> GOSSAFUNC=main go build src/a/hello-world/a01/main.go
 */
func main() {
	fmt.Println("Hello World!")
}
