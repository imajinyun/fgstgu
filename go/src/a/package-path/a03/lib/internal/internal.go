package internal

import (
	"fmt"
	"io"
)

// Hello func.
func Hello(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s!\n", name)
}
