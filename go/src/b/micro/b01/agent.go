package b01

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	// Waiting state.
	Waiting = iota

	// Running state.
	Running
)

// ErrWrongState var.
var ErrWrongState = errors.New("Can not take the operation in the current state")

// CollectionError struct.
type CollectionError struct {
	CollectionErrors []error
}

func (ce CollectionError) Error() string {
	var strs []string
	for _, err := range ce.CollectionErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

// Event struct.
type Event struct {
	Source  string
	Content string
}

// EventReceiver interface.
type EventReceiver interface {
	OnEvent(evt Event)
}

// Collector interface.
type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

// Agent struct.
type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int
}

// EventProcessGroutine func.
func (a *Agent) EventProcessGroutine() {
	var evtSeq [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeq[i] = <-a.evtBuf:
			case <-a.ctx.Done():
				return
			}
		}
		fmt.Println("🗳", evtSeq)
	}
}

// NewAgent func.
func NewAgent(sizeEvtBuf int) *Agent {
	agent := Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
	return &agent
}

// RegisterCollector func.
func (a *Agent) RegisterCollector(name string, collector Collector) error {
	if a.state != Waiting {
		return ErrWrongState
	}
	a.collectors[name] = collector
	return collector.Init(a)
}

func (a *Agent) startCollectors() error {
	var (
		err   error
		errs  CollectionError
		mutex sync.Mutex
	)
	for name, collector := range a.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs.CollectionErrors = append(errs.CollectionErrors,
					errors.New(name+":"+err.Error()))
			}
		}(name, collector, a.ctx)
	}
	if len(errs.CollectionErrors) == 0 {
		return nil
	}
	return errs
}

func (a *Agent) stopCollectors() error {
	var err error
	var errs CollectionError
	for name, collector := range a.collectors {
		if err = collector.Stop(); err != nil {
			errs.CollectionErrors = append(errs.CollectionErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectionErrors) == 0 {
		return nil
	}
	return errs
}

func (a *Agent) destoryCollectors() error {
	var err error
	var errs CollectionError
	for name, collector := range a.collectors {
		if err = collector.Destroy(); err != nil {
			errs.CollectionErrors = append(errs.CollectionErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectionErrors) == 0 {
		return nil
	}
	return errs
}

// Start func.
func (a *Agent) Start() error {
	if a.state != Waiting {
		return ErrWrongState
	}
	a.state = Running
	a.ctx, a.cancel = context.WithCancel(context.Background())
	go a.EventProcessGroutine()
	return a.startCollectors()
}

// Stop func.
func (a *Agent) Stop() error {
	if a.state != Running {
		return ErrWrongState
	}
	a.state = Waiting
	a.cancel()
	return a.stopCollectors()
}

// Destroy func.
func (a *Agent) Destroy() error {
	if a.state != Waiting {
		return ErrWrongState
	}
	return a.destoryCollectors()
}

// OnEvent func.
func (a *Agent) OnEvent(evt Event) {
	a.evtBuf <- evt
}
