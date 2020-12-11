package c01

import (
	"strconv"
	"sync"
	"testing"
)

// Const define.
const (
	NumOfReader = 100
	NumOfWriter = 10
)

// Map interface.
type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, m Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < NumOfWriter; j++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					m.Set(strconv.Itoa(i), i*i)
					m.Set(strconv.Itoa(i), i*i)
					m.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for k := 0; k < NumOfReader; k++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					m.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	b.Run("Map with RWLock", func(b *testing.B) {
		m := CreateRWLockMap()
		benchmarkMap(b, m)
	})
}
