package a10

import (
	"reflect"
	"testing"
)

func TestAddElementToSlice(t *testing.T) {
	var a []int
	b := []int{}
	t.Logf("%v, %v", a,  a == nil)
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
		t.Logf("a[%d] = %s",i, a[i])
	}
}
