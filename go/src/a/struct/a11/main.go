package main

import (
	"fmt"
	"testing"
)

// TB struct.
type TB struct {
	testing.TB
}

// Fatal method for TB.
func (t *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled")
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("hello", "world", "playground")
}
