package b01

import "testing"

func TestUnit(t *testing.T) {
	t.Log("Hello World")
}

func TestFail(t *testing.T) {
	t.Log("Before fail message")
	t.Fail()
	t.Log("After fail message")
}

func TestFailNow(t *testing.T) {
	t.Log("Before")
	t.FailNow()
	t.Log("After")
}
