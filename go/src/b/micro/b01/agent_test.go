package b01

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (dc *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("Initialize collector", dc.name)
	dc.evtReceiver = evtReceiver
	return nil
}

func (dc *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("Start collector", dc.name)
	for {
		select {
		case <-agtCtx.Done():
			dc.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			dc.evtReceiver.OnEvent(Event{dc.name, dc.content})
		}
	}
}

func (dc *DemoCollector) Stop() error {
	fmt.Println("Stop collector", dc.name)
	select {
	case <-dc.stopChan:
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("Failed to stop for timeout")
	}
}

func (dc *DemoCollector) Destroy() error {
	fmt.Println("Release resources", dc.name)
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(9999)
	a := NewCollect("A", "1")
	b := NewCollect("B", "2")
	agt.RegisterCollector("c1", a)
	agt.RegisterCollector("c2", b)
	if err := agt.Start(); err != nil {
		t.Logf("Start error %v\n", err)
	}
	t.Log(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destroy()
}
