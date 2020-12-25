package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

// Publisher struct.
type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

// NewPublisher returns a Publisher.
func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Publish to publish a topic.
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// Remove to delete the specified subscriber.
func (p *Publisher) Remove(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// Close to close all subscribers.
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// Subscribe add a new subscriber to subscribers.
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic add a new filter topic to subscribers.
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}
