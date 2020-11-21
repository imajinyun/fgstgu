package b01

import (
	"fmt"
	"testing"
	"time"
)

func message() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func task() {
	fmt.Println("Working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestMessage(t *testing.T) {
	fmt.Println(message())
	task()
}

func asyncMessage() (retChan chan string) {
	retChan = make(chan string, 1)
	go func() {
		ret := message()
		fmt.Println("Return result")
		retChan <- ret
		fmt.Println("Message exited")
	}()
	return
}

func TestAsyncMessage(t *testing.T) {
	retChan := asyncMessage()
	task()
	fmt.Println(<-retChan)
	time.Sleep(time.Second)
}
