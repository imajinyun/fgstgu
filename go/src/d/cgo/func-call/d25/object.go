package main

import "sync"

// ObjectID type.
type ObjectID int32

var refs struct {
	sync.Mutex
	objs map[ObjectID]interface{}
	next ObjectID
}

func init() {
	refs.Lock()
	defer refs.Unlock()
	refs.objs = make(map[ObjectID]interface{})
	refs.next = 1000
}

// NewObjectID returns a new ObjectID.
func NewObjectID(obj interface{}) ObjectID {
	refs.Lock()
	defer refs.Unlock()
	id := refs.next
	refs.next++
	refs.objs[id] = obj
	return id
}

// IsNil to determine if the specified ID is zero.
func (id ObjectID) IsNil() bool {
	return id == 0
}

// Get to get the specified ID from refs.
func (id ObjectID) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()
	return refs.objs[id]
}

// Free to delete the specified ID from refs.
func (id *ObjectID) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()
	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0
	return obj
}
