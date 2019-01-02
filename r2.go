package r2

import "fmt"

type Ins interface{}

type OpMatch struct{}

type OpJmp struct {
	N int
}

type OpSplit struct {
	A int
	B int
}

type OpString struct {
	Value string
}

type OpRange struct {
	Start byte
	End   byte
}

type Thread struct {
	P  []Ins
	PC int
}

func (t *Thread) Split() *Thread {
	return &Thread{P: t.P, PC: t.PC}
}

func Run(s string, t *Thread) *Thread {
	switch op := t.P[t.PC].(type) {
	case OpRange:
		if len(s) == 0 {
			return nil
		} else if s[0] < op.Start || s[0] > op.End {
			return nil
		}
		t.PC += 1
		return Run(s[1:], t)
	case OpString:
		if len(s) < len(op.Value) {
			return nil
		} else if s[0:len(op.Value)] != op.Value {
			return nil
		}
		t.PC += 1
		return Run(s[len(op.Value):], t)
	case OpMatch:
		if len(s) == 0 {
			return t
		}
		return nil
	case OpJmp:
		t.PC += op.N
		return Run(s, t)
	case OpSplit:
		t1 := t
		t2 := t.Split()
		t1.PC += op.A
		t2.PC += op.B
		if t1 := Run(s, t1); t1 != nil {
			return t1
		}
		return Run(s, t2)
	}
	panic("unreachable")
}

func Match(p []Ins) []Ins {
	return append(p, OpMatch{})
}

func String(s string) []Ins {
	return []Ins{OpString{s}}
}

func Concat(parts ...[]Ins) []Ins {
	var newP []Ins
	for _, p := range parts {
		newP = append(newP, p...)
	}
	return newP
}

func Alt(a, b []Ins) []Ins {
	split := OpSplit{1, len(a) + 2}
	jmp := OpJmp{len(b) + 1}
	return append(append(append([]Ins{split}, a...), jmp), b...)
}

func Plus(p []Ins) []Ins {
	return append(p, OpSplit{-len(p), 1})
}

func Star(p []Ins) []Ins {
	split := OpSplit{1, len(p) + 2}
	jmp := OpJmp{-(len(p) + 1)}
	return append(append([]Ins{split}, p...), jmp)
}

func QuestionMark(p []Ins) []Ins {
	split := OpSplit{1, len(p) + 1}
	return append([]Ins{split}, p...)
}

func Repeat(min, max int, p []Ins) []Ins {
	var newP []Ins
	for i := 0; i < min; i++ {
		newP = Concat(newP, p)
	}
	for i := 0; i < max-min; i++ {
		newP = Concat(newP, QuestionMark(p))
	}
	return newP
}

func Range(start byte, end byte) []Ins {
	return []Ins{OpRange{start, end}}
}

func Alpha() []Ins {
	return Alt(Range('a', 'z'), Range('A', 'Z'))
}

func Capture(name string, p []Ins) []Ins {
	return p
}

func Print(p []Ins) {
	for i, ins := range p {
		fmt.Printf("% 3d: ", i)
		switch op := ins.(type) {
		case OpRange:
			fmt.Printf("range %q %q", string(op.Start), string(op.End))
		case OpString:
			fmt.Printf("string %q", op.Value)
		case OpSplit:
			fmt.Printf("split %d %d", i+op.A, i+op.B)
		case OpJmp:
			fmt.Printf("jmp %d", i+op.N)
		}
		fmt.Print("\n")
	}
}
