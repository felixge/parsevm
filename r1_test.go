package r1

import (
	"testing"
)

func TestMatch(t *testing.T) {
	// a+b+
	ab := Program{
		{Op: OpChar, C: 'a'},
		{Op: OpSplit},
		{Op: OpChar, C: 'b'},
		{Op: OpSplit},
		{Op: OpMatch},
	}
	ab[1].X = ab[0:]
	ab[1].Y = ab[2:]
	ab[3].X = ab[2:]
	ab[3].Y = ab[4:]

	tests := []struct {
		Program Program
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
	}

	for _, test := range tests {
		got := test.Program.Match(test.Input)
		if got != test.Want {
			t.Errorf("%q got=%t want=%t", test.Input, got, test.Want)
		}
	}
}
