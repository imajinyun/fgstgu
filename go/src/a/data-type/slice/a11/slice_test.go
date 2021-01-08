package a11

import (
	"bytes"
	"testing"
)

func TestSlice(t *testing.T) {
	var n int = 5
	a, b := make([]int, n), make([]int, n)
	t.Log(a)
	t.Log(b)

	for i := 0; i < n; i++ {
		a[i]++
		b[i]++
	}
	t.Log(a)
	t.Log(b)
}

func TestSliceAppend(t *testing.T) {
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)

	t.Logf("a: len(a) = %d, cap(a) = %d", len(a), cap(a))
	t.Logf("b: len(b) = %d, cap(b) = %d", len(b), cap(b))
}

func TestSliceErrorShareData(t *testing.T) {
	path := []byte("1234/567890")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	t.Logf("dir1: %v", string(dir1))
	t.Logf("dir2: %v", string(dir2))

	dir1 = append(dir1, "suffix"...)
	t.Logf("dir1: %v", string(dir1))
	t.Logf("dir2: %v", string(dir2))
}

func TestSliceFixedShareData(t *testing.T) {
	path := []byte("1234/567890")
	sepIndex := bytes.IndexByte(path, '/')

	// 使用了 Full Slice Expression，其中最后一个参数叫「Limited Capacity」，
	// 于是，后续的 append() 操作将会导致重新分配内存。
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	t.Logf("dir1: %v", string(dir1))
	t.Logf("dir2: %v", string(dir2))

	dir1 = append(dir1, "suffix"...)
	t.Logf("dir1: %v", string(dir1))
	t.Logf("dir2: %v", string(dir2))
}
