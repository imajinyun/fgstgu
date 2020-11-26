package b01

import "testing"

func TestSumFilter(t *testing.T) {
	sf := NewSumFilter()
	actual, err := sf.Process([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Fatal(err)
	}
	expect := 15
	if actual != expect {
		t.Fatalf("\nActual: %d\nExpect: %d", expect, actual)
	}
}

func TestWrongInputForSumFilter(t *testing.T) {
	sf := NewSumFilter()
	_, err := sf.Process([]float32{1.01, 2.02, 3.03})
	if err == nil {
		t.Fatal("An error is expected")
	}
}
