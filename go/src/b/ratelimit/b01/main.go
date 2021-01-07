package main

import (
	"fmt"
	"time"
)

var capacity = 100
var tokenBucket = make(chan struct{}, capacity)

func main() {
	var fillInterval = time.Millisecond * 10
	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Printf("current token count: %3d, time: %s\n", len(tokenBucket), time.Now())
			}
		}
	}
	go fillToken()
	time.Sleep(time.Hour)
}
