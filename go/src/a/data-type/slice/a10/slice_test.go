package a10

import (
	"reflect"
	"testing"
)

func TestAddElementToSlice(t *testing.T) {
	var a []int
	b := []int{}
	t.Logf("%v, %v", a, a == nil)
	t.Logf("%v, %v", b, b == nil)
	t.Logf("%v", reflect.DeepEqual(a, b))

	a = append(a, 1)
	t.Logf("%v", a)

	a = append(a, 1, 2, 3)
	t.Logf("%v", a)

	a = append(a, []int{4, 5, 6}...)
	t.Logf("%v", a)

	a = append([]int{7, 8, 9}, a...)
	t.Logf("%v", a)

	c := []int{1, 2, 3}
	d := c[0:2:cap(c)]
	t.Logf("%v, %d, %d", c, len(c), cap(c))
	t.Logf("%v, %d, %d", d, len(d), cap(d))

	e := []int{1, 2, 3}
	e = append(e[:2], append([]int{2}, e[2:]...)...)
	t.Logf("%v", e)

	f := []int{1, 2, 6}
	f = append(f[:2], append([]int{3, 4, 5}, f[2:]...)...)
	t.Logf("%v", f)
}

func TestListSliceElement(t *testing.T) {
	a := []string{"a", "b", "c"}

	for i := range a {
		t.Logf("a[%d] = %s", i, a[i])
	}
	for i, v := range a {
		t.Logf("a[%d] = %s", i, v)
	}
	for i := 0; i < len(a); i++ {
		t.Logf("a[%d] = %s", i, a[i])
	}
}

func TestRemoveElementFromSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a = a[:len(a)-1]
	t.Log(a)

	a = a[:len(a)-3]
	t.Log(a)

	b := []string{"a", "b", "c", "d", "e"}
	b = b[1:]
	t.Log(b)

	c := []int{1, 2, 3}
	t.Log(c[:0])
	c = append(c[:0], c[1:]...)
	t.Log(c)

	d := []int{1, 2, 3}
	d = d[:copy(d, d[1:])]
	t.Log(d)
}

func TestRemoveHeaderElementFromSlice(t *testing.T) {
	// 删除开头的 1 个元素
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:0], a[1:]...)

	b := []int{1, 2, 3, 4, 5}
	b = b[:copy(b, b[1:])]
	t.Log(a, b)

	// 删除开头的 N 个元素，比如：N = 3
	c := []int{1, 2, 3, 4, 5}
	c = append(c[:0], c[3:]...)

	d := []int{1, 2, 3, 4, 5}
	d = d[:copy(d, d[3:])]
	t.Log(c, d)
}

func TestRemoveMiddleElementFromSlice(t *testing.T) {
	// 删除中间 1 个元素
	a := []int{1, 2, 3, 4, 5}
	i := len(a) / 2
	a = append(a[:i], a[i+1:]...)

	b := []int{1, 2, 3, 4, 5}
	b = b[:i+copy(b[i:], b[i+1:])]
	t.Log(a, b)

	// 删除中间 N 个元素，比如：N = 3
	c := []int{1, 2, 3, 4, 5}
	i, n := 1, 3
	c = append(c[:i], c[i+n:]...)

	d := []int{1, 2, 3, 4, 5}
	d = d[:i+copy(d[i:], d[i+n:])]
	t.Log(c, d)
}

func TestRemoveFooterElementFromSlice(t *testing.T) {
	// 删除尾部的 1 个元素
	a := []int{1, 2, 3, 4, 5}
	a = a[:len(a)-1]

	b := []int{1, 2, 3, 4, 5}
	b = b[:copy(b, b[:len(b)-1])]
	t.Log(a, b)

	// 删除尾部的 N 个元素，比如：N = 3
	c := []int{1, 2, 3, 4, 5}
	n := 3
	c = c[:len(c)-n]

	d := []int{1, 2, 3, 4, 5}
	d = d[:copy(d[:n], d[:len(b)+1-n])]
	t.Log(c, d)
}
