package b07

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetSingletonInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Create instance")
		instance = new(Singleton)
	})
	return instance
}

func TestGetSingletonInstance(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			object := GetSingletonInstance()
			fmt.Printf("%x, %v\n", unsafe.Pointer(object), unsafe.Pointer(object))
			wg.Done()
		}()
	}
	wg.Wait()
}
