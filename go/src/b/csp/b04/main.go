package main

import (
	"fgstgu/go/src/b/csp/b04/pubsub"
	"fmt"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	filterTopics := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("🎉 hello world")
	p.Publish("🔥 hello new world")
	p.Publish("🐶 hello golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()
	go func() {
		for msg := range filterTopics {
			fmt.Println("topics:", msg)
		}
	}()
	time.Sleep(3 * time.Second)
}
