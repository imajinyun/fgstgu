package b08

import (
	"fmt"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	op := NewObjectPool(10)
	// if err := op.ReleaseObject(&ObjectResuable{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := op.GetObject(time.Second); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := op.ReleaseObject(v); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("Done")
}
