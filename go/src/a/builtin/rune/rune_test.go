package rune_test

import (
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	t.Logf("%#v\n", []rune("世界"))
	t.Logf("%#v\n", string([]rune{'世', '界'}))
}

func TestStrToRune(t *testing.T) {
	b := []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0x47, 0x6f}
	r := strToRune(b)
	t.Logf("\n%#v\n%v\n", r, r)
}

func TestRuneToStr(t *testing.T) {
	r := []int32{20320, 22909, 71, 111}
	s := runeToStr(r)
	t.Logf("\n%#v\n%v\n", s, s)
}

func strToRune(s []byte) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

func runeToStr(b []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range b {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
