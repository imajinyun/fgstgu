package a01

import (
	"regexp"
	"testing"
)

var matchTests = []struct {
	pattern  string
	b        []byte
	expected bool
}{
	{"^\\d{3,5}$", []byte("12345"), true},
	{"^\\d{3,}$", []byte("789"), true},
	{"^[\u4e00-\u9fa5]$", []byte("中"), true},
	{"^\\d{3,}\\D{2}$", []byte("123ab"), true},
	{"^\\d{3,}\\D{2}$", []byte("12345"), false},
}

func TestMatch(t *testing.T) {
	for _, test := range matchTests {
		matched, _ := regexp.Match(test.pattern, test.b)
		if matched != test.expected {
			t.Errorf("Match(%s, %v) = %v want %v",
				test.pattern, test.b, !matched, test.expected)
		}
	}
}
