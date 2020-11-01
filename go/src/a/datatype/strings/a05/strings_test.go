package a05

import (
	"strings"
	"testing"
	"unicode"
)

type StringTest struct {
	in, out string
}

func runStringTests(t *testing.T, f func(string) string, name string, cases []StringTest) {
	for _, test := range cases {
		actual := f(test.in)
		if actual != test.out {
			t.Errorf("%s(%q) = %q, want %q", name, test.in, actual, test.out)
		}
	}
}

var TitleTests = []struct {
	in, out string
}{
	{"", ""},
	{"a", "A"},
	{" aaa aaa aaa ", " Aaa Aaa Aaa "},
	{" Aaa Aaa Aaa ", " Aaa Aaa Aaa "},
	{"123a456", "123a456"},
	{"double-blind", "Double-Blind"},
	{"ÿøû", "Ÿøû"},
	{"with_underscore", "With_underscore"},
	{"unicode \xe2\x80\xa8 line separator", "Unicode \xe2\x80\xa8 Line Separator"},
}

func TestTitle(t *testing.T) {
	for _, tt := range TitleTests {
		if s := strings.Title(tt.in); s != tt.out {
			t.Errorf("Title(%q) = %q, want %q", tt.in, s, tt.out)
		}
	}
}

var lowerTests = []StringTest{
	{"", ""},
	{"abc", "abc"},
	{"AbC123", "abc123"},
	{"azAZ09_", "azaz09_"},
	{"longStrinGwitHmixofsmaLLandcAps", "longstringwithmixofsmallandcaps"},
	{"LONG\u2C6FSTRING\u2C6FWITH\u2C6FNONASCII\u2C6FCHARS", "long\u0250string\u0250with\u0250nonascii\u0250chars"},
	{"\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D", "\u0251\u0251\u0251\u0251\u0251"},
	{"A\u0080\U0010FFFF", "a\u0080\U0010FFFF"},
}

var upperTests = []StringTest{
	{"", ""},
	{"ONLYUPPER", "ONLYUPPER"},
	{"abc", "ABC"},
	{"AbC123", "ABC123"},
	{"azAZ09_", "AZAZ09_"},
	{"longStrinGwitHmixofsmaLLandcAps", "LONGSTRINGWITHMIXOFSMALLANDCAPS"},
	{"long\u0250string\u0250with\u0250nonascii\u2C6Fchars", "LONG\u2C6FSTRING\u2C6FWITH\u2C6FNONASCII\u2C6FCHARS"},
	{"\u0250\u0250\u0250\u0250\u0250", "\u2C6F\u2C6F\u2C6F\u2C6F\u2C6F"},
	{"a\u0080\U0010FFFF", "A\u0080\U0010FFFF"},
}

func TestToLower(t *testing.T) { runStringTests(t, strings.ToLower, "ToLower", lowerTests) }
func TestToUpper(t *testing.T) { runStringTests(t, strings.ToUpper, "ToUpper", upperTests) }

func TestSpecialCase(t *testing.T) {
	lower := "abcçdefgğhıijklmnoöprsştuüvyz"
	upper := "ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ"
	u := strings.ToUpperSpecial(unicode.TurkishCase, upper)
	if u != upper {
		t.Errorf("Upper(upper) is %s not %s", u, upper)
	}
	u = strings.ToUpperSpecial(unicode.TurkishCase, lower)
	if u != upper {
		t.Errorf("Upper(lower) is %s not %s", u, upper)
	}

	l := strings.ToLowerSpecial(unicode.TurkishCase, lower)
	if l != lower {
		t.Errorf("Lower(lower) is %s not %s", l, lower)
	}
	l = strings.ToLowerSpecial(unicode.TurkishCase, upper)
	if l != lower {
		t.Errorf("Lower(upper) is %s not %s", l, lower)
	}
}
