package c01

import "sync"

// RWLockMap struct.
type RWLockMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

// Get method for RWLockMap.
func (m *RWLockMap) Get(key interface{}) (interface{}, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()
	return v, ok
}

// Set method for RWLockMap.
func (m *RWLockMap) Set(key, value interface{}) {
	m.lock.Lock()
	m.m[key] = value
	m.lock.Unlock()
}

// Del method for RWLockMap.
func (m *RWLockMap) Del(key interface{}) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}

// CreateRWLockMap func.
func CreateRWLockMap() *RWLockMap {
	return &RWLockMap{m: make(map[interface{}]interface{}, 0)}
}
