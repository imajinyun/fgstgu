package main

import (
	"time"

	"github.com/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	if err = l.Lock(); err != nil {
		panic(err)
	}
	println("lock successfully, start your business logic")
	time.Sleep(time.Second * 10)
	l.Unlock()
	println("unlock successfully, finished your business logic")
}
