package r2

import (
	"testing"
)

func TestMatch(t *testing.T) {
	str := Match(String("abc"))
	strConcat := Match(Concat(String("ab"), String("c")))
	strAlt := Match(Alt(String("abc"), String("def")))
	strPlus := Match(Plus(String("abc")))

	tests := []struct {
		Program []*Ins
		Input   string
		Want    bool
	}{
		{str, "abc", true},
		{str, "abd", false},
		{str, "dbc", false},
		{str, "ab", false},
		{str, "bc", false},
		{str, "", false},

		{strConcat, "abc", true},
		{strConcat, "abd", false},
		{strConcat, "dbc", false},
		{strConcat, "ab", false},
		{strConcat, "bc", false},
		{strConcat, "", false},

		{strAlt, "def", true},
		{strAlt, "abc", true},
		{strAlt, "abd", false},
		{strAlt, "dbc", false},
		{strAlt, "ab", false},
		{strAlt, "bc", false},
		{strAlt, "", false},

		{strPlus, "abcabc", true},
		{strPlus, "abcabcabc", true},
		{strPlus, "abc", true},
		{strPlus, "abd", false},
		{strPlus, "dbc", false},
		{strPlus, "ab", false},
		{strPlus, "bc", false},
		{strPlus, "", false},
	}

	for _, test := range tests {
		got := Run(test.Input, test.Program, 0)
		if got != test.Want {
			t.Errorf("%q got=%t want=%t", test.Input, got, test.Want)
		}
	}
}
