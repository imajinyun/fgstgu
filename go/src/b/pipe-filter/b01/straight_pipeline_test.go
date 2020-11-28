package b01

import "testing"

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p", spliter, converter, sum)
	ret, err := sp.Process("1,2,3,4,5")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 15 {
		t.Fatalf("\nActual: %v\nExpcect: %v", ret, 15)
	}
}
