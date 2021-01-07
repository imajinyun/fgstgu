package main

import (
	"context"
	"fgstgu/go/src/b/grpc/b13/api"
	"fgstgu/go/src/b/grpc/b13/pubsub"
	"fmt"
	"strings"
	"time"
)

// PublishService struct.
type PublishService struct {
	pub *pubsub.Publisher
}

// NewPublisherService returns a new PublishService.
func NewPublisherService() *PublishService {
	return &PublishService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

// Publish method for PublishService.
func (p *PublishService) Publish(ctx context.Context, arg *api.String) (*api.String, error) {
	p.pub.Publish(arg.GetValue())
	return &api.String{}, nil
}

// Subscribe method for PublishService.
func (p *PublishService) Subscribe(arg *api.String, stream api.PublishService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&api.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hello")
	go p.Publish("world")
	go p.Publish("golang: https://golang.org/")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(1)
	go func() { fmt.Println("golang topic:", <-golang) }()
	go func() { fmt.Println("docker topic:", <-docker) }()
	<-make(chan bool)
}
