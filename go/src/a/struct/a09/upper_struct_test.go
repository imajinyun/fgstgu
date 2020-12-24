package a09

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// UpperWriter struct.
type UpperWriter struct {
	io.Writer
}

func (u *UpperWriter) Write(data []byte) (n int, err error) {
	return u.Writer.Write(bytes.ToUpper(data))
}

func TestUpperWriter(t *testing.T) {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}
