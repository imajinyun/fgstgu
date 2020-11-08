package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	ch := make(chan int)
	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	go producer(ch)
	go consumer(1, ch, a)
	go consumer(2, ch, b)
	go consumer(3, ch, c)
	<-a
	<-b
	<-c
	defer fmt.Println("The func main is over!")
}

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("ID: ", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	defer close(ch)
}

func consumer(num int, ch chan int, bchan chan bool) {
	for data := range ch {
		pre := strings.Repeat("🦋", num)
		fmt.Printf("%s %d %d\n", pre, num, data)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	bchan <- true
	defer close(bchan)
}
