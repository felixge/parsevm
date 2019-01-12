package parsevm

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func compileJSON() []Ins {
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
		Star(Alt(
			Range('a', 'z'),
			Range('A', 'Z'),
			Range('0', '9'),
		)),
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
	return Concat(
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
}

var jsonExample = `
{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
	"key4": "value4",
	"key5": "value5",
	"key6": "value6",
	"key7": "value7",
	"key8": "value8",
	"key9": "value9",
	"key10": {
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
		"key7": "value7",
		"key8": "value8",
		"key9": "value9"
	}
}
`

func BenchmarkJSON(b *testing.B) {
	p := Optimize(compileJSON())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		thr := Run(jsonExample, p)
		if thr == nil {
			b.FailNow()
		}
	}
}

func BenchmarkGoJSON(b *testing.B) {
	in := []byte(jsonExample)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !json.Valid(in) {
			b.FailNow()
		}
	}
}

func TestJSON(t *testing.T) {
	p := Optimize(compileJSON())

	//Print(p)
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
