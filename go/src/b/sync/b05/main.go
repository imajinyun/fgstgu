package main

import (
	"fmt"
	"sync/atomic"
)

var seq int64

func genID() int64 {
	return atomic.AddInt64(&seq, 1)
}

// 竞态检测
func main() {
	for i := 0; i < 10; i++ {
		go genID()
	}
	fmt.Println(genID())
}
