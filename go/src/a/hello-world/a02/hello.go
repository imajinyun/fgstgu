package a02

/**
 * 将代码编译成汇编指令：
 * -> GOOS=linux GOARCH=amd64 go tool compile -S src/a/hello-world/a02/hello.go
 */
func hello(x int) (y int) {
	y = x + 10
	return
}
