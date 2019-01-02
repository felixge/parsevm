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

type OpCapture struct {
	Name  string
	Start bool
}

type CaptureGroup struct {
	Name  string
	Value string
	Depth int
	Done  bool
}

type Thread struct {
	S        string
	P        []Ins
	PC       int
	Captures []CaptureGroup
	Match    bool
}

func (t *Thread) Clone() *Thread {
	captures := make([]CaptureGroup, len(t.Captures))
	copy(captures, t.Captures)
	return &Thread{S: t.S, P: t.P, PC: t.PC, Captures: captures}
}

func Run(s string, p []Ins) *Thread {
	threads := []*Thread{{P: p, S: s}}
	for len(threads) > 0 {
		newThreads := make([]*Thread, 0, len(threads))
		for _, t := range threads {
			children := t.next()
			for _, child := range children {
				if child.Match {
					return child
				}
			}
			newThreads = append(newThreads, children...)
		}
		threads = newThreads
	}
	return nil
}

func (t *Thread) next() []*Thread {
	self := []*Thread{t}
	switch op := t.P[t.PC].(type) {
	case OpCapture:
		t.PC += 1
		if op.Start {
			t.startCapture(op.Name)
		} else {
			t.endCapture()
		}
		return self
	case OpRange:
		if len(t.S) == 0 {
			return nil
		} else if t.S[0] < op.Start || t.S[0] > op.End {
			return nil
		}
		t.PC += 1
		t.capture(string(t.S[0]))
		t.S = t.S[1:]
		return self
	case OpString:
		if len(t.S) < len(op.Value) {
			return nil
		} else if t.S[0:len(op.Value)] != op.Value {
			return nil
		}
		t.PC += 1
		t.capture(op.Value)
		t.S = t.S[len(op.Value):]
		return self
	case OpMatch:
		if len(t.S) == 0 {
			t.Match = true
			return self
		}
		return nil
	case OpJmp:
		t.PC += op.N
		return self
	case OpSplit:
		clone := t.Clone()
		t.PC += op.A
		clone.PC += op.B
		return append(self, clone)
	}
	panic("unreachable")
}
func (t *Thread) startCapture(name string) {
	depth := 0
	for _, capture := range t.Captures {
		if !capture.Done {
			depth++
		}
	}
	t.Captures = append(t.Captures, CaptureGroup{Name: name, Depth: depth})
}

func (t *Thread) endCapture() {
	for i := len(t.Captures) - 1; i >= 0; i-- {
		if !t.Captures[i].Done {
			t.Captures[i].Done = true
			break
		}
	}
}

func (t *Thread) capture(s string) {
	for i := range t.Captures {
		if !t.Captures[i].Done {
			t.Captures[i].Value += s
		}
	}
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
	start := OpCapture{name, true}
	stop := OpCapture{name, false}
	return append(append([]Ins{start}, p...), stop)
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
		case OpCapture:
			start := "start"
			if !op.Start {
				start = "end"
			}
			fmt.Printf("capture %s %s", start, op.Name)
		case OpMatch:
			fmt.Printf("match")
		default:
			panic("unreachable")
		}
		fmt.Print("\n")
	}
}
