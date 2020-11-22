package b09

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Print("Create a new object: ")
			return 666
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(33)
	runtime.GC()
	v, _ = pool.Get().(int)
	fmt.Println(v)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Print("Create a new object: ")
			return 555
		},
	}
	pool.Put(111)
	pool.Put(222)
	pool.Put(333)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
