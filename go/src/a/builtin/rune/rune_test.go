package rune_test

import "testing"

func TestRune(t *testing.T) {
	t.Logf("%#v\n", []rune("世界"))
	t.Logf("%#v\n", string([]rune{'世', '界'}))
}
