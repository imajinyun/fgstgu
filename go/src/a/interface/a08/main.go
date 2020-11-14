package main

import (
	"fmt"
)

// Dict struct
type Dict struct {
	data map[interface{}]interface{}
}

// Get func.
func (d *Dict) Get(key interface{}) interface{} {
	return d.data[key]
}

// Set func.
func (d *Dict) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// Visit func.
func (d *Dict) Visit(fn func(k, v interface{}) bool) {
	if fn == nil {
		return
	}
	for k, v := range d.data {
		if !fn(k, v) {
			return
		}
	}
}

// Clear func.
func (d *Dict) Clear() {
	d.data = make(map[interface{}]interface{})
}

// NewDict func.
func NewDict() *Dict {
	d := &Dict{}
	d.Clear()
	return d
}

func main() {
	dict := NewDict()
	dict.Set("Item 1", 100)
	dict.Set("Item 2", 80)
	dict.Set("Item 3", 60)
	dict.Set("Item 4", 50)
	dict.Set("Item 5", 30)

	value := dict.Get("Item 3")
	fmt.Printf("%v\n", value)

	dict.Visit(func(k, v interface{}) bool {
		if v.(int) >= 60 {
			fmt.Printf("Value: %v\n", v)
			return true
		}
		fmt.Printf("Value: %v is passed\n", v)
		return false
	})
}
