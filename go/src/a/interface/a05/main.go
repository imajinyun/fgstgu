package main

import (
	"fmt"
	"io"
)

type device struct{}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func main() {
	var wc io.WriteCloser = new(device)
	fmt.Printf("%+v\n", wc)
	wc.Write(nil)
	wc.Close()

	var writer io.Writer = new(device)
	writer.Write(nil)
	fmt.Printf("%+v\n", writer)
}
