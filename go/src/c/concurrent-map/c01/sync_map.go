package c01

import "sync"

// SyncMapBenchmarkAdapter struct.
type SyncMapBenchmarkAdapter struct {
	m sync.Map
}

// Set method for SyncMapBenchmarkAdapter.
func (s *SyncMapBenchmarkAdapter) Set(key, value interface{}) {
	s.m.Store(key, value)
}

// Get method for SyncMapBenchmarkAdapter.
func (s *SyncMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return s.m.Load(key)
}

// Del method for SyncMapBenchmarkAdapter.
func (s *SyncMapBenchmarkAdapter) Del(key interface{}) {
	s.m.Delete(key)
}

// CreateSyncMapBenchmarkAdapter func.
func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
	return &SyncMapBenchmarkAdapter{}
}
