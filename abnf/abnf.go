package abnf

import pvm "github.com/felixge/parsevm"

var abnf = compileABNF()

func compileABNF() []pvm.Ins {
	var p []pvm.Ins
	define := func(name string, ins []pvm.Ins) {
		p = append(p, pvm.Func(name, ins)...)
	}

	define("rulelist", pvm.Plus(pvm.Alt(
		pvm.Call("rule"),
		pvm.Concat(pvm.Star(pvm.Call("c-wsp")), pvm.Call("c-nl")),
	)))

	define("rule", pvm.Concat(
		pvm.Call("rulename"),
		pvm.Call("defined-as"),
		pvm.Call("elements"),
		pvm.Call("c-nl"),
	))

	define("rulename", pvm.Concat(pvm.Concat(
		pvm.Call("ALPHA"),
		pvm.Star(pvm.Alt(
			pvm.Call("ALPHA"),
			pvm.Call("DIGIT"),
			pvm.String("-"),
		)),
	)))

	define("defined-as", pvm.Concat(
		pvm.Star(pvm.Call("c-wsp")),
		pvm.Alt(pvm.String("="), pvm.String("=/")),
		pvm.Star(pvm.Call("c-wsp")),
	))

	define("elements", pvm.Concat(
		pvm.Call("alternation"),
		pvm.Star(pvm.Call("c-wsp")),
	))

	define("c-wsp", pvm.Alt(
		pvm.Call("WSP"),
		pvm.Concat(pvm.Call("c-nl"), pvm.Call("WSP")),
	))

	define("c-nl", pvm.Alt(
		pvm.Call("comment"),
		pvm.Call("CRLF"),
	))

	define("comment", pvm.Concat(
		pvm.String(";"),
		pvm.Star(pvm.Alt(pvm.Call("WSP"), pvm.Call("VCHAR"))),
		pvm.Call("CRLF"),
	))

	define("alternation", pvm.Concat(
		pvm.Call("concatenation"),
		pvm.Star(pvm.Concat(
			pvm.Star(pvm.Call("c-wsp")),
			pvm.String("/"),
			pvm.Star(pvm.Call("c-wsp")),
			pvm.Call("concatenation"),
		)),
	))

	define("concatenation", pvm.Concat(
		pvm.Call("repetition"),
		pvm.Star(pvm.Concat(
			pvm.Plus(pvm.Call("c-wsp")),
			pvm.Call("repetition"),
		)),
	))

	define("repetition", pvm.Concat(
		pvm.QuestionMark(pvm.Call("repeat")),
		pvm.Call("element"),
	))

	define("repeat", pvm.Alt(
		pvm.Plus(pvm.Call("DIGIT")),
		pvm.Concat(
			pvm.Star(pvm.Call("DIGIT")),
			pvm.String("*"),
			pvm.Star(pvm.Call("DIGIT")),
		),
		pvm.Call("element"),
	))

	define("element", pvm.Alt(
		pvm.Call("rulename"),
		pvm.Call("group"),
		pvm.Call("option"),
		pvm.Call("char-val"),
		pvm.Call("num-val"),
		pvm.Call("prose-val"),
	))

	define("group", pvm.Concat(
		pvm.String("("),
		pvm.Star(pvm.Call("c-wsp")),
		pvm.Call("alternation"),
		pvm.Star(pvm.Call("c-wsp")),
		pvm.String(")"),
	))

	define("option", pvm.Concat(
		pvm.String("["),
		pvm.Star(pvm.Call("c-wsp")),
		pvm.Call("alternation"),
		pvm.Star(pvm.Call("c-wsp")),
		pvm.String("]"),
	))

	define("char-val", pvm.Concat(
		pvm.Call("DQUOTE"),
		pvm.Star(pvm.Alt(
			pvm.Range('\x20', '\x21'),
			pvm.Range('\x23', '\x7E'),
		)),
		pvm.Call("DQUOTE"),
	))

	define("num-val", pvm.Concat(
		pvm.String("%"),
		pvm.Alt(
			pvm.Call("bin-val"),
			pvm.Call("dec-val"),
			pvm.Call("hex-val"),
		),
	))

	define("bin-val", pvm.Concat(
		pvm.String("b"),
		pvm.Plus(pvm.Call("BIT")),
	))

	define("dec-val", pvm.Concat(
		pvm.String("d"),
		pvm.Plus(pvm.Call("DIGIT")),
		pvm.QuestionMark(pvm.Alt(
			pvm.Plus(pvm.Concat(pvm.String("."), pvm.Plus(pvm.Call("DIGIT")))),
			pvm.Concat(pvm.String("-"), pvm.Plus(pvm.Call("DIGIT"))),
		)),
	))

	define("hex-val", pvm.Concat(
		pvm.String("x"),
		pvm.Plus(pvm.Call("HEXDIG")),
		pvm.QuestionMark(pvm.Alt(
			pvm.Plus(pvm.Concat(pvm.String("."), pvm.Plus(pvm.Call("HEXDIG")))),
			pvm.Concat(pvm.String("-"), pvm.Plus(pvm.Call("HEXDIG"))),
		)),
	))

	define("prose-val", pvm.Concat(
		pvm.String("<"),
		pvm.Star(pvm.Alt(
			pvm.Range('\x20', '\x3d'),
			pvm.Range('\x3f', '\x7e'),
		)),
		pvm.String(">"),
	))

	// builtin
	define("ALPHA", pvm.Alt(
		pvm.Range('\x41', '\x5a'),
		pvm.Range('\x61', '\x7a'),
	))

	define("BIT", pvm.Alt(
		pvm.String("0"),
		pvm.String("1"),
	))

	define("CHAR", pvm.Range('\x01', '\x7f'))

	define("CR", pvm.String("\x0d"))

	define("CRLF", pvm.Alt(
		pvm.Call("CR"),
		pvm.Call("LF"),
	))

	define("CTL", pvm.Range('\x00', '\x1f'))

	define("DIGIT", pvm.Range('\x30', '\x39'))

	define("DQUOTE", pvm.String("\x22"))

	define("HEXDIGIT", pvm.Alt(
		pvm.Call("DIGIT"),
		pvm.String("A"),
		pvm.String("B"),
		pvm.String("C"),
		pvm.String("D"),
		pvm.String("E"),
		pvm.String("F"),
	))

	define("HTAB", pvm.String("\x09"))

	define("LF", pvm.String("\x0a"))

	define("LWSP", pvm.Star(pvm.Alt(
		pvm.Call("WSP"),
		pvm.Concat(pvm.Call("CRLF"), pvm.Call("WSP")),
	)))

	define("OCTET", pvm.Range('\x00', '\xff'))

	define("SP", pvm.String("\x20"))

	define("VCHAR", pvm.Range('\x21', '\x7e'))

	define("WSP", pvm.Alt(
		pvm.Call("SP"),
		pvm.Call("HTAB"),
	))

	return pvm.Concat(pvm.Concat(pvm.Call("rulelist"), pvm.Match(nil)), p)
}

func Validate(s string) bool {
	//pvm.Print(abnf)
	t := pvm.Run(s, abnf)
	return t != nil
}
