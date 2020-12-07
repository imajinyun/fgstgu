package datatype

import (
	"testing"
	"unsafe"
)

var (
	intNum     = 0
	float64Num = 0.00
	str        = ""
	boolean    = false
	array      = [...]int{1}
	slice      = []string{}
	intStrMap  = map[int]string{}
	intBoolmap = map[int]bool{}
	st         = new(struct{})
	iface      = new(interface{})
)

func TestSize(t *testing.T) {
	t.Logf("type: %T, size: %d\n", intNum, unsafe.Sizeof(intNum))
	t.Logf("type: %T, size: %d\n", float64Num, unsafe.Sizeof(float64Num))
	t.Logf("type: %T, size: %d\n", str, unsafe.Sizeof(str))
	t.Logf("type: %T, size: %d\n", boolean, unsafe.Sizeof(boolean))
	t.Logf("type: %T, size: %d\n", array, unsafe.Sizeof(array))
	t.Logf("type: %T, size: %d\n", slice, unsafe.Sizeof(slice))
	t.Logf("type: %T, size: %d\n", intStrMap, unsafe.Sizeof(intStrMap))
	t.Logf("type: %T, size: %d\n", intBoolmap, unsafe.Sizeof(intBoolmap))
	t.Logf("type: %T, size: %d\n", st, unsafe.Sizeof(st))
	t.Logf("type: %T, size: %d\n", iface, unsafe.Sizeof(iface))
}
