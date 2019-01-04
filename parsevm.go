package parsevm

import (
	"bytes"
	"fmt"
	"strings"
)

type Ins interface{}

type OpMatch struct{}

type OpJmp struct {
	N int
}

type OpFork struct {
	N int
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

type OpCall struct {
	Name string
}

type OpFunc struct {
	Name string
}

type OpReturn struct{}

type CaptureGroup struct {
	Name  string
	Value string
	Depth int
	Done  bool
}

type Thread struct {
	ID       int
	S        string
	P        []Ins
	PC       int
	Captures []CaptureGroup
	Match    bool
	Stack    []int
}

func (t *Thread) Clone() *Thread {
	captures := make([]CaptureGroup, len(t.Captures))
	copy(captures, t.Captures)
	stack := make([]int, len(t.Stack))
	copy(stack, t.Stack)
	return &Thread{S: t.S, P: t.P, PC: t.PC, Captures: captures, Stack: stack}
}

func Run(s string, p []Ins) *Thread {
	threadID := 1
	threads := []*Thread{{ID: threadID, P: p, S: s}}
	for len(threads) > 0 {
		newThreads := make([]*Thread, 0, len(threads))
		for _, t := range threads {
			children := t.next()
			for _, child := range children {
				if child.ID == 0 {
					threadID++
					child.ID = threadID
				}
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
	op := t.P[t.PC]
	//fmt.Printf("t=% 3d pc=% 3d % 20s %#v\n", t.ID, t.PC, t.S, op)
	switch op := op.(type) {
	case OpFunc:
		t.PC++
		return self
	case OpCall:
		newPC := callPC(t.P, op)
		if newPC == -1 {
			panic(fmt.Sprintf("cannot call unkown func: %s", op.Name))
		}
		t.Stack = append(t.Stack, t.PC+1)
		//fmt.Printf("call %#v\n", t.Stack)
		t.PC = newPC + 1
		return self
	case OpReturn:
		//fmt.Printf("return %#v\n", t.Stack)
		//if len(t.Stack) == 0 {
		//t.PC++
		//return self
		//}
		tail := len(t.Stack) - 1
		t.PC = t.Stack[tail]
		t.Stack = t.Stack[0:tail]
		return self
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
		t.PC++
		t.capture(string(t.S[0]))
		t.S = t.S[1:]
		return self
	case OpString:
		if len(t.S) < len(op.Value) {
			return nil
		} else if t.S[0:len(op.Value)] != op.Value {
			return nil
		}
		t.PC++
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
	case OpFork:
		clone := t.Clone()
		t.PC += 1
		clone.PC += op.N
		return append(self, clone)
	}
	panic(fmt.Sprintf("unknown op: %#v", op))
}

func callPC(p []Ins, c OpCall) int {
	pc := -1
	for i, ins := range p {
		fn, ok := ins.(OpFunc)
		if !ok || fn.Name != c.Name {
			continue
		}
		pc = i
		break
	}
	return pc
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
	fork := OpFork{len(a) + 2}
	jmp := OpJmp{len(b) + 1}
	return append(append(append([]Ins{fork}, a...), jmp), b...)
}

func Plus(p []Ins) []Ins {
	return append(p, OpFork{-len(p)})
}

func Star(p []Ins) []Ins {
	fork := OpFork{len(p) + 2}
	jmp := OpJmp{-(len(p) + 1)}
	return append(append([]Ins{fork}, p...), jmp)
}

func QuestionMark(p []Ins) []Ins {
	fork := OpFork{len(p) + 1}
	return append([]Ins{fork}, p...)
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

func Func(name string, ins []Ins) []Ins {
	//return append(append([]Ins{OpCall{name}, OpJmp{len(ins) + 3}, OpFunc{name}}, ins...), OpReturn{})
	return append(append([]Ins{OpFunc{name}}, ins...), OpReturn{})
}

func Call(name string) []Ins {
	return []Ins{OpCall{name}}
}

func Graphviz(p []Ins) []byte {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "digraph g {\n")
	for i, ins := range p {
		var label string
		switch op := ins.(type) {
		case OpRange:
			label = fmt.Sprintf("range %q %q", string(op.Start), string(op.End))
		case OpString:
			label = fmt.Sprintf("string %q", op.Value)
		case OpFork:
			label = fmt.Sprintf("fork %d", i+op.N)
			fmt.Fprintf(buf, "%d -> %d;", i, i+op.N)
		case OpJmp:
			label = fmt.Sprintf("jmp %d", i+op.N)
			fmt.Fprintf(buf, "%d -> %d;", i, i+op.N)
		case OpCapture:
			//start := "start"
			//if !op.Start {
			//start = "end"
			//}
			//fmt.Printf("capture %s %s", start, op.Name)
		case OpMatch:
			label = fmt.Sprintf("match")
		case OpFunc:
			label = fmt.Sprintf("func %s", op.Name)
			//indent++
		case OpCall:
			label = fmt.Sprintf("call %s", op.Name)
			fmt.Fprintf(buf, "%d -> %d;", i, callPC(p, op))
		case OpReturn:
			label = fmt.Sprintf("return")
		default:
			panic("unreachable")
		}
		label = fmt.Sprintf("[%d] %s", i, label)
		fmt.Fprintf(buf, "%d[label=%q];", i, label)
		if i > 0 {
			fmt.Fprintf(buf, "%d -> %d;", i-1, i)
		}
		//fmt.Print("\n")
	}
	fmt.Fprintf(buf, "}\n")
	return buf.Bytes()
}

func Print(p []Ins) {
	indent := 0
	for i, ins := range p {
		if _, ok := ins.(OpReturn); ok {
			indent--
		}
		fmt.Printf("% 3d: %s", i, strings.Repeat("  ", indent))
		switch op := ins.(type) {
		case OpRange:
			fmt.Printf("range %q %q", string(op.Start), string(op.End))
		case OpString:
			fmt.Printf("string %q", op.Value)
		case OpFork:
			fmt.Printf("fork %d", i+op.N)
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
		case OpFunc:
			fmt.Printf("func %s", op.Name)
			indent++
		case OpCall:
			fmt.Printf("call %s", op.Name)
		case OpReturn:
			fmt.Printf("return")
		default:
			panic("unreachable")
		}
		fmt.Print("\n")
	}
}
