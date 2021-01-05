package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// 通过输出当前栈帧的信息来输出 goid
func main() {
	var buf = make([]byte, 64)
	var stk = buf[:runtime.Stack(buf, false)]
	fmt.Println(getGoid())
	fmt.Println(string(stk))
}

// 从 runtime.Stack() 获取的字符串中解析出 goid
func getGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	field := strings.Fields(stk)[0]
	id, err := strconv.Atoi(field)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
