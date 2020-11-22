package b08

import (
	"errors"
	"time"
)

// ObjectResuable struct.
type ObjectResuable struct{}

// ObjectPool struct.
type ObjectPool struct {
	bufChan chan *ObjectResuable
}

// NewObjectPool func.
func NewObjectPool(numOfObject int) *ObjectPool {
	op := ObjectPool{}
	op.bufChan = make(chan *ObjectResuable, numOfObject)
	for i := 0; i < numOfObject; i++ {
		op.bufChan <- &ObjectResuable{}
	}
	return &op
}

// GetObject func.
func (op *ObjectPool) GetObject(timeout time.Duration) (*ObjectResuable, error) {
	select {
	case ret := <-op.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("Timeout")
	}
}

// ReleaseObject func.
func (op *ObjectPool) ReleaseObject(object *ObjectResuable) error {
	select {
	case op.bufChan <- object:
		return nil
	default:
		return errors.New("Overflow")
	}
}
