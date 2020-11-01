package a07

import (
	"strings"
	"testing"
)

var compareTests = []struct {
	a, b string
	i    int
}{
	{"", "", 0},
	{"a", "", 1},
	{"", "a", -1},
	{"abc", "abc", 0},
	{"ab", "abc", -1},
	{"abc", "ab", 1},
	{"x", "ab", 1},
	{"ab", "x", -1},
	{"x", "a", 1},
	{"b", "x", -1},
	{"abcdefgh", "abcdefgh", 0},
	{"abcdefghi", "abcdefghi", 0},
	{"abcdefghi", "abcdefghj", -1},
}

func TestCompare(t *testing.T) {
	for _, v := range compareTests {
		cmp := strings.Compare(v.a, v.b)
		if cmp != v.i {
			t.Errorf("`Compare(%q, %q) = %v`", v.a, v.b, v.i)
		}
	}
}

var equalFoldTests = []struct {
	a, b     string
	expected bool
}{
	{"Go", "go", true},
	{" ", " ", true},
}

func TestEqualFold(t *testing.T) {
	for _, v := range equalFoldTests {
		cmp := strings.EqualFold(v.a, v.b)
		if cmp != v.expected {
			t.Errorf("`EqualFold(%s, %s) = %v`", v.a, v.b, v.expected)
		}
	}
}

var repeatTests = []struct {
	a, b string
	l, r int
}{
	{"o", "oooooooo", 8, 0},
	{"test ", "test test ", 2, 0},
	{"", "", 20, 0},
	{"😭", "😭😭😭", 3, 0},
	{"!#?=<>", "!#?=<>", 0, -1},
}

func TestRepeat(t *testing.T) {
	for _, v := range repeatTests {
		a := strings.Repeat(v.a, v.l)
		cmp := strings.Compare(a, v.b)
		if cmp != v.r {
			t.Errorf("`Repeat(%s, %v) = %v`, want %v", v.a, v.l, v.r, cmp)
		}
	}
}

var replaceTests = []struct {
	s, old, new string
	n           int
	res         string
	expected    int
}{
	{"This is a test", "test", "demo", 1, "This is a demo", 0},
	{"中国", "中", "美", -1, "中国", 1},
}

func TestReplace(t *testing.T) {
	for _, v := range replaceTests {
		a := strings.Replace(v.s, v.old, v.new, v.n)
		cmp := strings.Compare(a, v.res)
		if cmp != v.expected {
			t.Errorf("`Replace(%s, %s, %s, %v)` = %v, want %v",
				v.s, v.old, v.new, v.n, a, cmp)
		}
	}
}

var joinTests = []struct {
	elems    []string
	sep      string
	expected string
	i        int
}{
	{[]string{"a", "b", "c"}, "", "abc", 0},
	{[]string{}, ",", "", 0},
	{[]string{"1", "2", "3"}, ",", "1,2,3", 0},
}

func TestJoin(t *testing.T) {
	for _, v := range joinTests {
		res := strings.Join(v.elems, v.sep)
		cmp := strings.Compare(res, v.expected)
		if cmp != v.i {
			t.Errorf("Join(%p, %s) = %v, want %v", v.elems, v.sep, res, cmp)
		}
	}
}
