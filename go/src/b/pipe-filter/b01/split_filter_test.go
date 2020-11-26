package b01

import (
	"reflect"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	sf := NewSplitFilter(",")
	req, err := sf.Process("1,2,3,4,5")
	if err != nil {
		t.Fatal(err)
	}

	a := strings.Split("1,2,3", ",")
	b := []string{"1", "2", "3"}
	t.Log(a, b)

	parts, ok := req.([]string)
	if !ok {
		t.Fatalf("Response type is %T, but the expected type is string", parts)
	}
	expect := []string{"1", "2", "3", "4", "5"}
	if !reflect.DeepEqual(parts, expect) {
		t.Errorf("\nActual: %v\nExpect: %v", parts, expect)
	}
}

func TestWrongInput(t *testing.T) {
	sf := NewSplitFilter(",")
	_, err := sf.Process(999)
	if err == nil {
		t.Fatal("An error is expected")
	}
}
