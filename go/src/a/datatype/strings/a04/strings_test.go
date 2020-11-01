package a04

import (
	"reflect"
	"strings"
	"testing"
)

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var abcd = "abcd"
var faces = "☺☻☹"
var commas = "1,2,3,4"
var dots = "1....2....3....4"

type SplitTest struct {
	s   string
	sep string
	n   int
	a   []string
}

var splitTests = []SplitTest{
	{"", "", -1, []string{}},
	{abcd, "", 2, []string{"a", "bcd"}},
	{abcd, "", 4, []string{"a", "b", "c", "d"}},
	{abcd, "", -1, []string{"a", "b", "c", "d"}},
	{faces, "", -1, []string{"☺", "☻", "☹"}},
	{faces, "", 3, []string{"☺", "☻", "☹"}},
	{faces, "", 17, []string{"☺", "☻", "☹"}},
	{"☺�☹", "", -1, []string{"☺", "�", "☹"}},
	{abcd, "a", 0, nil},
	{abcd, "a", -1, []string{"", "bcd"}},
	{abcd, "z", -1, []string{"abcd"}},
	{commas, ",", -1, []string{"1", "2", "3", "4"}},
	{dots, "...", -1, []string{"1", ".2", ".3", ".4"}},
	{faces, "☹", -1, []string{"☺☻", ""}},
	{faces, "~", -1, []string{faces}},
	{"1 2 3 4", " ", 3, []string{"1", "2", "3 4"}},
	{"1 2", " ", 3, []string{"1", "2"}},
}

func TestSplit(t *testing.T) {
	for _, v := range splitTests {
		a := strings.SplitN(v.s, v.sep, v.n)
		if !eq(a, v.a) {
			t.Errorf("SplitN(%q, %q, %d) = %v; want %v",
				v.s, v.sep, v.n, a, v.a)
		}
		if v.n == 0 {
			continue
		}
		s := strings.Join(a, v.sep)
		if s != v.s {
			t.Errorf("Join(Split(%q, %q, %d), %q) = %q", v.s, v.sep, v.n, v.sep, s)
		}
		if v.n < 0 {
			b := strings.Split(v.s, v.sep)
			if !reflect.DeepEqual(a, b) {
				t.Errorf("Split disagrees with SplitN(%q, %q, %d) = %v; want %v",
					v.s, v.sep, v.n, b, a)
			}
		}
	}
}

var splitAfterTests = []SplitTest{
	{abcd, "a", -1, []string{"a", "bcd"}},
	{abcd, "z", -1, []string{"abcd"}},
	{abcd, "", -1, []string{"a", "b", "c", "d"}},
	{commas, ",", -1, []string{"1,", "2,", "3,", "4"}},
	{dots, "...", -1, []string{"1...", ".2...", ".3...", ".4"}},
	{faces, "☹", -1, []string{"☺☻☹", ""}},
	{faces, "~", -1, []string{faces}},
	{faces, "", -1, []string{"☺", "☻", "☹"}},
	{"1 2 3 4", " ", 3, []string{"1 ", "2 ", "3 4"}},
	{"1 2 3", " ", 3, []string{"1 ", "2 ", "3"}},
	{"1 2", " ", 3, []string{"1 ", "2"}},
	{"123", "", 2, []string{"1", "23"}},
	{"123", "", 17, []string{"1", "2", "3"}},
}

func TestSplitAfter(t *testing.T) {
	for _, tt := range splitAfterTests {
		a := strings.SplitAfterN(tt.s, tt.sep, tt.n)
		if !eq(a, tt.a) {
			t.Errorf("Split(%q, %q, %q) = %v; want %v",
				tt.s, tt.sep, tt.n, a, tt.a)
			continue
		}
		s := strings.Join(a, "")
		if s != tt.s {
			t.Errorf("Join(Split(%q, %q, %d), %q) = %q",
				tt.s, tt.sep, tt.n, a, tt.a)
		}
		if tt.n < 0 {
			b := strings.SplitAfter(tt.s, tt.sep)
			if !reflect.DeepEqual(a, b) {
				t.Errorf("SplitAfter disagrees with SplitAfterN(%q, %q, %d) = %v; want %v",
					tt.s, tt.sep, tt.n, b, a)
			}
		}
	}
}
