package r2

import (
	"fmt"
	"io/ioutil"
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

func TestRecursive(t *testing.T) {
	value := Func("value", Alt(Alt(Alt(Call("number"), Call("array")), Call("object")), Call("string")))

	object := Func("object", Concat(
		String("{"),
		Alt(Call("ws"), Call("members")),
		String("}"),
	))

	members := Func("members", Alt(
		Call("member"),
		Concat(Call("member"), String(","), Call("members")),
	))

	member := Func("member", Concat(
		Call("ws"),
		Call("string"),
		Call("ws"),
		String(":"),
		Call("element"),
	))

	array := Func("array", Concat(
		String("["),
		Alt(Call("ws"), Call("elements")),
		String("]"),
	))

	elements := Func("elements", Alt(
		Call("element"),
		Concat(Call("element"), String(","), Call("elements")),
	))

	element := Func("element", Concat(Call("ws"), Call("value"), Call("ws")))

	str := Func("string", Concat(
		String(`"`),
		Star(Range('a', 'z')),
		String(`"`),
	))

	number := Func("number", Range('0', '9'))

	// TODO more convenient Alt
	ws := Func("ws", Alt(Alt(Alt(Alt(
		String(""),
		Concat(String("\r"), Call("ws"))),
		Concat(String("\n"), Call("ws"))),
		Concat(String("\t"), Call("ws"))),
		Concat(String(" "), Call("ws"))),
	)
	p := Concat(
		Call("element"),
		Match(nil),
		value,
		object,
		members,
		member,
		array,
		elements,
		element,
		str,
		number,
		ws,
	)
	Print(p)
	dot := Graphviz(p)
	if err := ioutil.WriteFile("json.dot", dot, 0666); err != nil {
		t.Fatal(err)
	}

	tests := []string{
		`"foo"`,
		`""`,

		"[]",
		"[1]",
		"[1,2]",
		"[1, 2]",
		"[1,2,3]",
		"[[]]",
		"[[1]]",
		"[[1,2]]",
		"[[1,2,3]]",
		"[[],4]",
		"[[1],4]",
		"[[1,2],4]",
		"[[1,2,3],4]",
		"[[[]]]",

		`{}`,
		`{ }`,
		`{"a":"b"}`,
		`{"a": "b"}`,
		`{"a":"b","c":"d"}`,
		`{"a": "b", "c": "d"}`,
		`{"a": "b", "c": {"d": "e"}}`,
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			thr := Run(test, p)
			if thr == nil {
				t.Fail()
			}
		})
	}

	//Match(Func("array", Concat(String("[")

	//p := Match(Func("foo", Concat(String("("), Alt(Call("foo"), String("bar")), String(")"))))
	//Print(p)
	//thread := Run(`(((bar)))`, p)
	//fmt.Printf("%#v\n", thread)
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
	number := token(Capture("number", Plus(Range('0', '9'))))
	val := Alt(str, number)
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
