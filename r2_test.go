package r2

import (
	"testing"
)

func TestMatch(t *testing.T) {
	ab := Match(Concat(Plus(String("a")), Plus(String("b"))))
	hello := Match(String("hello"))
	helloOrWorld := Match(Alt(String("hello"), String("world")))
	Print(helloOrWorld)

	tests := []struct {
		Program []*Ins
		Input   string
		Want    bool
	}{
		{ab, "ab", true},
		{ab, "aab", true},
		{ab, "abb", true},
		{ab, "aabb", true},

		{ab, "ac", false},
		{ab, "cb", false},
		{ab, "aaaac", false},

		{hello, "hello", true},
		{hello, "world", false},
		{hello, "ello", false},
		{hello, "hell", false},
		{hello, "hella", false},

		{helloOrWorld, "hello", true},
		{helloOrWorld, "world", true},
		{helloOrWorld, "ello", false},
		{helloOrWorld, "hell", false},
		{helloOrWorld, "hella", false},
	}

	for _, test := range tests {
		got := Run(test.Input, test.Program, 0)
		if got != test.Want {
			t.Errorf("%q got=%t want=%t", test.Input, got, test.Want)
		}
	}
}
