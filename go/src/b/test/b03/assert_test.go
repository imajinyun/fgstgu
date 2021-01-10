package b03

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	Assert(t, true, "a", "b")
	Assert(t, false, "a", "b", "c", "d")
	Assert(t, false, 1, 2, 3, 4, 5)
}

func TestAssertFunction(t *testing.T) {
	AssertFunc(t, func() error {
		return errors.New("test assert function1")
	})

	fn := func() error {
		return errors.New("test assert function2")
	}
	AssertFunc(t, fn)
}
