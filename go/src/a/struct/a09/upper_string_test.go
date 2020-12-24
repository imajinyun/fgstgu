package a09

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func TestUpperString(t *testing.T) {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
