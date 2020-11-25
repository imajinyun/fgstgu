package b01

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 999
	f := *((*float64)(unsafe.Pointer(&i)))
	t.Log(unsafe.Pointer(&i))
	t.Log((*float64)(unsafe.Pointer(&i)))
	t.Log(f)
}

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(a)
	t.Log(unsafe.Pointer(&a))
	t.Log((*[]MyInt)(unsafe.Pointer(&a)))
	t.Log(b)
}

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writerFn := func() {
		data := []int{}
		for i := 0; i < 30; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readerFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		t.Log(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writerFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writerFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readerFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

type Customer struct {
	Name string
	Age  int
}

type MyInt int
