package b03

import (
	"fmt"
	"testing"
)

type testingHelper interface {
	Helper()
}

// Assert func.
func Assert(tb testing.TB, condition bool, args ...interface{}) {
	if x, ok := tb.(testingHelper); ok {
		x.Helper()
	}
	if !condition {
		if msg := fmt.Sprint(args...); msg != "" {
			tb.Fatalf("assert failed: %s", msg)
		} else {
			tb.Fatalf("assert failed")
		}
	}
}

// AssertFunc func.
func AssertFunc(tb testing.TB, fn func() error) {
	if x, ok := tb.(testingHelper); ok {
		x.Helper()
	}
	if err := fn(); err != nil {
		tb.Fatalf("assert func failed: %v", err)
	}
}
