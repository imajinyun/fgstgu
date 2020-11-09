package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// 所有子 Goroutine 运行结束以后主 Goroutine 才退出
func main() {
	var wg sync.WaitGroup
	fmt.Printf("wg: %T, %v\n", wg, wg)
	wg.Add(3)
	rand.Seed(time.Now().UnixNano())
	go dump(&wg, 1)
	go dump(&wg, 2)
	go dump(&wg, 3)
	wg.Wait()
	defer fmt.Println("The main func is over!")
}

func dump(wg *sync.WaitGroup, num int) {
	for i := 1; i <= 3; i++ {
		pre := strings.Repeat("\t", num-1)
		fmt.Printf("%s NO.%d goroutine num: %d\n", pre, num, i)
		time.Sleep(time.Second)
	}
	wg.Done()
}
