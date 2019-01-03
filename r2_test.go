package r2

import (
	"fmt"
	"strings"
	"testing"
)

func TestMatch(t *testing.T) {
	str := Match(String("abc"))
	strConcat := Match(Concat(String("ab"), String("c")))
	strAlt := Match(Alt(String("abc"), String("def")))
	strPlus := Match(Plus(String("abc")))
	strStar := Match(Star(String("abc")))
	strQuestionMark := Match(QuestionMark(String("abc")))
	strRepeat := Match(Repeat(2, 3, String("abc")))
	strAlpha := Match(Alpha())

	tests := []struct {
		Name    string
		Program []Ins
		Input   string
		Want    bool
	}{
		{"str", str, "abc", true},
		{"str", str, "abd", false},
		{"str", str, "dbc", false},
		{"str", str, "ab", false},
		{"str", str, "bc", false},
		{"str", str, "", false},

		{"strConcat", strConcat, "abc", true},
		{"strConcat", strConcat, "abd", false},
		{"strConcat", strConcat, "dbc", false},
		{"strConcat", strConcat, "ab", false},
		{"strConcat", strConcat, "bc", false},
		{"strConcat", strConcat, "", false},

		{"strAlt", strAlt, "def", true},
		{"strAlt", strAlt, "abc", true},
		{"strAlt", strAlt, "abd", false},
		{"strAlt", strAlt, "dbc", false},
		{"strAlt", strAlt, "ab", false},
		{"strAlt", strAlt, "bc", false},
		{"strAlt", strAlt, "", false},

		{"strPlus", strPlus, "abcabc", true},
		{"strPlus", strPlus, "abcabcabc", true},
		{"strPlus", strPlus, "abc", true},
		{"strPlus", strPlus, "abd", false},
		{"strPlus", strPlus, "dbc", false},
		{"strPlus", strPlus, "ab", false},
		{"strPlus", strPlus, "bc", false},
		{"strPlus", strPlus, "", false},

		{"strStar", strStar, "abcabc", true},
		{"strStar", strStar, "abcabcabc", true},
		{"strStar", strStar, "abc", true},
		{"strStar", strStar, "abd", false},
		{"strStar", strStar, "dbc", false},
		{"strStar", strStar, "ab", false},
		{"strStar", strStar, "bc", false},
		{"strStar", strStar, "", true},

		{"strQuestionMark", strQuestionMark, "abcabc", false},
		{"strQuestionMark", strQuestionMark, "abcabcabc", false},
		{"strQuestionMark", strQuestionMark, "abc", true},
		{"strQuestionMark", strQuestionMark, "abd", false},
		{"strQuestionMark", strQuestionMark, "dbc", false},
		{"strQuestionMark", strQuestionMark, "ab", false},
		{"strQuestionMark", strQuestionMark, "bc", false},
		{"strQuestionMark", strQuestionMark, "", true},

		{"strRepeat", strRepeat, "abc", false},
		{"strRepeat", strRepeat, "abcabc", true},
		{"strRepeat", strRepeat, "abcabcabc", true},
		{"strRepeat", strRepeat, "abcabcabcabc", false},
		{"strRepeat", strRepeat, "abd", false},
		{"strRepeat", strRepeat, "dbc", false},
		{"strRepeat", strRepeat, "ab", false},
		{"strRepeat", strRepeat, "bc", false},
		{"strRepeat", strRepeat, "", false},

		{"strAlpha", strAlpha, "a", true},
		{"strAlpha", strAlpha, "b", true},
		{"strAlpha", strAlpha, "y", true},
		{"strAlpha", strAlpha, "z", true},
		{"strAlpha", strAlpha, "0", false},
		{"strAlpha", strAlpha, "-", false},
		{"strAlpha", strAlpha, " ", false},
		{"strAlpha", strAlpha, "", false},
	}

	for _, test := range tests {
		//Print(test.Program)
		got := Run(test.Input, test.Program) != nil
		if got != test.Want {
			t.Errorf("test=%s input=%q got=%t want=%t", test.Name, test.Input, got, test.Want)
		}
	}
}

func TestJSON(t *testing.T) {
	t.Skip()
	whitespace := Star(Alt(Alt(Alt(
		String(" "),
		String("\n"),
	), String("\t")),
		String("\r"),
	))

	token := func(p []Ins) []Ins {
		return append(append(whitespace, p...), whitespace...)
	}

	str := token(Concat(
		String(`"`),
		Capture("string", Star(Range('a', 'z'))),
		String(`"`),
	))
	num := token(Capture("number", Plus(Range('0', '9'))))
	val := Alt(str, num)
	pair := Concat(
		Capture("key", str),
		String(":"),
		Capture("value", val),
	)
	object := token(Capture("object", Concat(
		String("{"),
		Alt(
			Concat(
				Plus(Concat(pair, Concat(String(",")))),
				pair,
			),
			Repeat(0, 1, pair),
		),
		String("}"),
	)))
	doc := Star(object)
	p := Match(doc)
	Print(p)

	tests := []string{
		`{}`,
		`{"foo": "bar"}`,
		`{"foo": 123}`,
		`{"foo": "bar", "hello": "world"}`,
		`{"foo": "bar", {"hello": "world"}}`,
	}
	for _, test := range tests {
		a := Run(test, p)
		fmt.Printf("%s:\n", test)
		for _, c := range a.Captures {
			fmt.Printf("%s%s %s\n", strings.Repeat("  ", c.Depth), c.Name, c.Value)
		}
		fmt.Printf("\n\n")
	}
}
