package b01

import (
	"reflect"
	"testing"
)

func TestConvertToInt(t *testing.T) {
	tif := NewToIntFilter()
	actual, err := tif.Process([]string{"1", "2", "3"})
	if err != nil {
		t.Fatal(err)
	}
	expect := []int{1, 2, 3}
	if !reflect.DeepEqual(actual, expect) {
		t.Fatalf("\nActual: %v\nExpect: %v", actual, expect)
	}
}

func TestWrongInputToIntFilter(t *testing.T) {
	tif := NewToIntFilter()
	_, err := tif.Process([]string{"hello", "world", "hello", "go"})
	if err == nil {
		t.Fatal("An error is expected for wrong input")
	}
	_, err = tif.Process([]int{1, 2, 3})
	if err == nil {
		t.Fatal("An error is expected for wrong input")
	}
}
