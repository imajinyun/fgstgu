package a03

import (
	"strings"
	"testing"
	"unicode"
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

type FieldsTest struct {
	s string
	a []string
}

var fieldsTests = []FieldsTest{
	{"", []string{}},
	{" ", []string{}},
	{" \t ", []string{}},
	{"\u2000", []string{}},
	{"  abc  ", []string{"abc"}},
	{"1 2 3 4", []string{"1", "2", "3", "4"}},
	{"1  2  3  4", []string{"1", "2", "3", "4"}},
	{"1\t\t2\t\t3\t4", []string{"1", "2", "3", "4"}},
	{"1\u20002\u20013\u20024", []string{"1", "2", "3", "4"}},
	{"\u2000\u2001\u2002", []string{}},
	{"\n™\t™\n", []string{"™", "™"}},
	{"\n\u20001™2\u2000 \u2001 ™", []string{"1™2", "™"}},
	{"\n1\uFFFD \uFFFD2\u20003\uFFFD4", []string{"1\uFFFD", "\uFFFD2", "3\uFFFD4"}},
	{"1\xFF\u2000\xFF2\xFF \xFF", []string{"1\xFF", "\xFF2\xFF", "\xFF"}},
	{faces, []string{faces}},
}

func TestFields(t *testing.T) {
	for _, tt := range fieldsTests {
		a := strings.Fields(tt.s)
		if !eq(a, tt.a) {
			t.Errorf("Fields(%q) = %v; want %v", tt.s, a, tt.s)
			continue
		}
	}
}

var FieldsFuncTests = []FieldsTest{
	{"", []string{}},
	{"XX", []string{}},
	{"XXhiXXX", []string{"hi"}},
	{"aXXbXXXcX", []string{"a", "b", "c"}},
}

func TestFieldsFunc(t *testing.T) {
	for _, tt := range fieldsTests {
		a := strings.FieldsFunc(tt.s, unicode.IsSpace)
		if !eq(a, tt.a) {
			t.Errorf("FieldsFunc(%q, unicode.IsSpace) = %v; want %v", tt.s, a, tt.a)
			continue
		}
	}
	pred := func(c rune) bool { return c == 'X' }
	for _, tt := range FieldsFuncTests {
		a := strings.FieldsFunc(tt.s, pred)
		if !eq(a, tt.a) {
			t.Errorf("FieldsFunc(%q) = %v; want %v", tt.s, a, tt.a)
		}
	}
}
