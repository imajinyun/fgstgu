package a06

import (
	"strings"
	"testing"
)

const space = "\t\v\r\f\n\u0085\u00a0\u2000\u3000"

type StringTest struct {
	in, out string
}

func runStringTests(t *testing.T, f func(string) string, name string, cases []StringTest) {
	for _, v := range cases {
		actual := f(v.in)
		if actual != v.out {
			t.Errorf("%s(%q) = %q, want %q", name, v.in, actual, v.out)
		}
	}
}

var trimSpaceTests = []StringTest{
	{"", ""},
	{"abc", "abc"},
	{space + "abc" + space, "abc"},
	{" ", ""},
	{" \t\r\n \t\t\r\r\n\n ", ""},
	{" \t\r\n x\t\t\r\r\n\n ", "x"},
	{" \u2000\t\r\n x\t\t\r\r\ny\n \u3000", "x\t\t\r\r\ny"},
	{"1 \t\r\n2", "1 \t\r\n2"},
	{" x\x80", "x\x80"},
	{" x\xc0", "x\xc0"},
	{"x \xc0\xc0 ", "x \xc0\xc0"},
	{"x \xc0", "x \xc0"},
	{"x \xc0 ", "x \xc0"},
	{"x \xc0\xc0 ", "x \xc0\xc0"},
	{"x ☺\xc0\xc0 ", "x ☺\xc0\xc0"},
	{"x ☺ ", "x ☺"},
}

func TestTrimSpace(t *testing.T) { runStringTests(t, strings.TrimSpace, "TrimSpace", trimSpaceTests) }

var trimTests = []struct {
	f, in, arg, out string
}{
	{"Trim", "abba", "a", "bb"},
	{"Trim", "abba", "ab", ""},
	{"TrimLeft", "abba", "ab", ""},
	{"TrimRight", "abba", "ab", ""},
	{"TrimLeft", "abba", "a", "bba"},
	{"TrimRight", "abba", "a", "abb"},
	{"Trim", "<tag>", "<>", "tag"},
	{"Trim", "* listitem", " *", "listitem"},
	{"Trim", `"quote"`, `"`, "quote"},
	{"Trim", "\u2C6F\u2C6F\u0250\u0250\u2C6F\u2C6F", "\u2C6F", "\u0250\u0250"},
	{"Trim", "\x80test\xff", "\xff", "test"},
	{"Trim", " Ġ ", " ", "Ġ"},
	{"Trim", " Ġİ0", "0 ", "Ġİ"},
	//empty string tests
	{"Trim", "abba", "", "abba"},
	{"Trim", "", "123", ""},
	{"Trim", "", "", ""},
	{"TrimLeft", "abba", "", "abba"},
	{"TrimLeft", "", "123", ""},
	{"TrimLeft", "", "", ""},
	{"TrimRight", "abba", "", "abba"},
	{"TrimRight", "", "123", ""},
	{"TrimRight", "", "", ""},
	{"TrimRight", "☺\xc0", "☺", "☺\xc0"},
	{"TrimPrefix", "aabb", "a", "abb"},
	{"TrimPrefix", "aabb", "b", "aabb"},
	{"TrimSuffix", "aabb", "a", "aabb"},
	{"TrimSuffix", "aabb", "b", "aab"},
}

func TestTrim(t *testing.T) {
	for _, v := range trimTests {
		name := v.f
		var f func(string, string) string
		switch name {
		case "Trim":
			f = strings.Trim
		case "TrimLeft":
			f = strings.TrimLeft
		case "TrimRight":
			f = strings.TrimRight
		case "TrimPrefix":
			f = strings.TrimPrefix
		case "TrimSuffix":
			f = strings.TrimSuffix
		default:
			t.Errorf("Undefined trim function %s", name)
		}
		actual := f(v.in, v.arg)
		if actual != v.out {
			t.Errorf("%s(%q, %q) = %q; want %q", name, v.in, v.arg, actual, v.out)
		}
	}
}
