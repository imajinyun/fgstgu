package c01

import "github.com/easierway/concurrent_map"

// ConcurrentMapBenchmarkAdapter struct.
type ConcurrentMapBenchmarkAdapter struct {
	cm *concurrent_map.ConcurrentMap
}

// Set method for ConcurrentMapBenchmarkAdapter.
func (c *ConcurrentMapBenchmarkAdapter) Set(key, value interface{}) {
	c.cm.Set(concurrent_map.StrKey(key.(string)), value)
}

// Get method for ConcurrentMapBenchmarkAdapter.
func (c *ConcurrentMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return c.cm.Get(concurrent_map.StrKey(key.(string)))
}

// Del method for ConcurrentMapBenchmarkAdapter.
func (c *ConcurrentMapBenchmarkAdapter) Del(key interface{}) {
	c.cm.Del(concurrent_map.StrKey(key.(string)))
}

// CreateConcurrentMapBenchmarkAdapter func.
func CreateConcurrentMapBenchmarkAdapter(numOfPartitions int) *ConcurrentMapBenchmarkAdapter {
	cm := concurrent_map.CreateConcurrentMap(numOfPartitions)
	return &ConcurrentMapBenchmarkAdapter{cm: cm}
}
