package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var tickets int = 10
var wg sync.WaitGroup
var mutex sync.Mutex

// 互斥锁
func main() {
	wg.Add(4)
	go saleTicket("NO.1 window:", &wg)
	go saleTicket("NO.2 window:", &wg)
	go saleTicket("NO.3 window:", &wg)
	go saleTicket("NO.4 window:", &wg)
	wg.Wait()
	defer fmt.Println("Ticket sales have stopped!")
}

func saleTicket(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(1 * time.Second)
			num, _ := strconv.Atoi(name[:1])
			pre := strings.Repeat("\t", num)
			fmt.Println(pre, name, tickets)
			tickets--
		} else {
			fmt.Printf("%s is over\n", name)
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}
