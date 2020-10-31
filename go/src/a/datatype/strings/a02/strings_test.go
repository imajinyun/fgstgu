package a02_test

import (
	"strings"
	"testing"
)

var dots = "1...2...3...4"

type IndexTest struct {
	s   string
	sep string
	out int
}

var indexTests = []IndexTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"oofofoofooo", "f", 2},
	{"oofofoofooo", "foo", 4},
	{"barfoobarfoo", "foo", 3},
	{"foo", "", 0},
	{"foo", "o", 1},
	{"abcABCabc", "A", 3},
	{"abcdef"[1:], "de", 2},
}

var lastIndexTests = []IndexTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"foo", "f", 0},
	{"oofofoofooo", "f", 7},
	{"oofofoofooo", "foo", 7},
	{"barfoobarfoo", "foo", 9},
	{"foo", "", 3},
	{"foo", "o", 2},
	{"abcABCabc", "A", 3},
	{"abcABCabc", "a", 6},
}

var indexAnyTests = []IndexTest{
	{"", "", -1},
	{"", "a", -1},
	{"", "abc", -1},
	{"a", "", -1},
	{"a", "a", 0},
	{"\x80", "\xffb", 0},
	{"aaa", "a", 0},
	{"abc", "xyz", -1},
	{"abc", "xcz", 2},
	{"ab☺c", "x☺yz", 2},
	{"a☺b☻c☹d", "cx", len("a☺b☻")},
	{"a☺b☻c☹d", "uvw☻xyz", len("a☺b")},
	{"aRegExp*", ".(|)*+?^$[]", 7},
	{dots + dots + dots, " ", -1},
	{"012abcba210", "\xffb", 4},
	{"012\x80bcb\x80210", "\xffb", 3},
	{"0123456\xcf\x80abc", "\xcfb\x80", 10},
}

var lastIndexAnyTests = []IndexTest{
	{"", "", -1},
	{"", "a", -1},
	{"", "abc", -1},
	{"a", "", -1},
	{"a", "a", 0},
	{"\x80", "\xffb", 0},
	{"aaa", "a", 2},
	{"abc", "xyz", -1},
	{"abc", "ab", 1},
	{"ab☺c", "x☺yz", 2},
	{"a☺b☻c☹d", "cx", len("a☺b☻")},
	{"a☺b☻c☹d", "uvw☻xyz", len("a☺b")},
	{"a.RegExp*", ".(|)*+?^$[]", 8},
	{dots + dots + dots, " ", -1},
	{"012abcba210", "\xffb", 6},
	{"012\x80bcb\x80210", "\xffb", 7},
	{"0123456\xcf\x80abc", "\xcfb\x80", 10},
}

func runIndexTests(t *testing.T, fn func(s, sep string) int, name string, cases []IndexTest) {
	for _, test := range cases {
		actual := fn(test.s, test.sep)
		if actual != test.out {
			t.Errorf("%s(\"%s\", \"%s\") = %v; want %v",
				name, test.s, test.sep, actual, test.out)
		}
	}
}

func TestIndex(t *testing.T)     { runIndexTests(t, strings.Index, "Index", indexTests) }
func TestLastIndex(t *testing.T) { runIndexTests(t, strings.LastIndex, "LastIndex", lastIndexTests) }
func TestIndexAny(t *testing.T)  { runIndexTests(t, strings.IndexAny, "IndexAny", indexAnyTests) }
func TestLastIndexAny(t *testing.T) {
	runIndexTests(t, strings.LastIndexAny, "LastIndexAny", lastIndexAnyTests)
}

func TestIndexByte(t *testing.T) {
	for _, tt := range indexTests {
		if len(tt.sep) != 1 {
			continue
		}
		pos := strings.IndexByte(tt.s, tt.sep[0])
		t.Log(tt.s, tt.sep, tt.out, pos)
		if pos != tt.out {
			t.Errorf("IndexByte(\"%q\", \"%q\") = %v; want %v",
				tt.s, tt.sep[0], pos, tt.out)
		}
	}
}

func TestRune(t *testing.T) {
	runes := []struct {
		in   string
		rune rune
		want int
	}{
		{"", 'a', -1},
		{"", '☺', -1},
		{"foo", '☹', -1},
		{"foo", 'o', 1},
		{"foo☺bar", '☺', 3},
		{"foo☺☻☹bar", '☹', 9},
		{"a A x", 'A', 2},
		{"some_text=some_value", '=', 9},
		{"☺a", 'a', 3},
		{"a☻☺b", '☺', 4},
	}
	for _, tt := range runes {
		if got := strings.IndexRune(tt.in, tt.rune); got != tt.want {
			t.Errorf("IndexRune(%q, %q) = %v; want %v",
				tt.in, tt.rune, got, tt.want)
		}
	}
}

var ContainsTests = []struct {
	str      string
	substr   string
	expected bool
}{
	{"abc", "bc", true},
	{"abc", "bcd", false},
	{"abc", "", true},
	{"", "a", false},
	{"", "", true},
	{"你好 Go 语言", "语言", true},
	{"xxxxxx", "01", false},
	{"01xxxx", "01", true},
	{"xx01xx", "01", true},
	{"xxxx01", "01", true},
	{"01xxxxx"[1:], "01", false},
	{"xxxxx01"[:6], "01", false},
}

func TestContains(t *testing.T) {
	for _, ct := range ContainsTests {
		if strings.Contains(ct.str, ct.substr) != ct.expected {
			t.Errorf("Contains(\"%s\", \"%s\") = %v, want %v",
				ct.str, ct.substr, !ct.expected, ct.expected)
		}
	}
}

var ContainsAnyTests = []struct {
	str      string
	substr   string
	expected bool
}{
	{"", "", false},
	{"", "a", false},
	{"", "abc", false},
	{"a", "", false},
	{"a", "a", true},
	{"aaa", "a", true},
	{"abc", "xyz", false},
	{"abc", "xcz", true},
	{"a☺b☻c☹d", "uvw☻xyz", true},
	{"aRegExp*", ".(|)*+?^$[]", true},
	{dots + dots + dots, " ", false},
}

func TestContainsAny(t *testing.T) {
	for _, ct := range ContainsAnyTests {
		if strings.ContainsAny(ct.str, ct.substr) != ct.expected {
			t.Errorf("ContainsAny(\"%s\", \"%s\") = %v, want %v",
				ct.str, ct.substr, !ct.expected, ct.expected)
		}
	}
}

var ContainsRuneTests = []struct {
	str      string
	r        rune
	expected bool
}{
	{"", 'a', false},
	{"a", 'a', true},
	{"aaa", 'a', true},
	{"abc", 'y', false},
	{"abc", 'c', true},
	{"a☺b☻c☹d", 'x', false},
	{"a☺b☻c☹d", '☻', true},
	{"aRegExp*", '*', true},
}

func TestContainsRune(t *testing.T) {
	for _, ct := range ContainsRuneTests {
		if strings.ContainsRune(ct.str, ct.r) != ct.expected {
			t.Errorf("ContainsRune(\"%q\", \"%q\") = %v, want %v",
				ct.str, ct.r, !ct.expected, ct.expected)
		}
	}
}

var CountTests = []struct {
	s   string
	sep string
	num int
}{
	{"", "", 1},
	{"", "notempty", 0},
	{"notempty", "", 9},
	{"smaller", "not smaller", 0},
	{"12345678987654321", "6", 2},
	{"611161116", "6", 3},
	{"notequal", "NotEqual", 0},
	{"equal", "equal", 1},
	{"abc1231231123q", "123", 3},
	{"11111", "11", 2},
}

func TestCount(t *testing.T) {
	for _, ct := range CountTests {
		if num := strings.Count(ct.s, ct.sep); num != ct.num {
			t.Errorf("Count(\"%s\", \"%s\") = %d, want %d",
				ct.s, ct.sep, num, ct.num)
		}
	}
}

var HasPrefixTests = []struct {
	s        string
	substr   string
	expected bool
}{
	{"", "", true},
	{"this is a test", "this", true},
	{"this is a test", "This", false},
	{"abc", "a", true},
	{"😭gogogo", "😭", true},
}

func TestHasPrefix(t *testing.T) {
	for _, ct := range HasPrefixTests {
		if strings.HasPrefix(ct.s, ct.substr) != ct.expected {
			t.Errorf("HasPrefix(\"%s\", \"%s\") = %v, want %v",
				ct.s, ct.substr, !ct.expected, ct.expected)
		}
	}
}

var HasSuffixTests = []struct {
	s        string
	substr   string
	expected bool
}{
	{"", "", true},
	{"this is a test", "test", true},
	{"this is a test", "Test", false},
	{"abc", "c", true},
	{"gogogo😭", "😭", true},
}

func TestHasSuffix(t *testing.T) {
	for _, ct := range HasSuffixTests {
		if strings.HasSuffix(ct.s, ct.substr) != ct.expected {
			t.Errorf("HasSuffix(\"%s\", \"%s\") = %v, want %v",
				ct.s, ct.substr, !ct.expected, ct.expected)
		}
	}
}
