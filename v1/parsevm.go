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
	N    int
	Cond string
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
	PC   int
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
	I        int
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
	return &Thread{S: t.S, I: t.I, P: t.P, PC: t.PC, Captures: captures, Stack: stack}
}

type Program []Ins

func NewVM(p *Program) *VM {
	return &VM{p: p}
}

type VM struct {
	p *Program
}

func (v *VM) Write(data []byte) (int, error) {
	return 0, nil
}

type Stats struct {
	ThreadsExecuted int
	ThreadsDeduped  int
	ThreadsParallel int
	Cycles          int
}

func Run(s string, p []Ins) *Thread {
	threadID := 1
	threads := []*Thread{{ID: threadID, P: p, S: s}}
	var stats Stats

	for len(threads) > 0 {
		if l := len(threads); l > stats.ThreadsParallel {
			stats.ThreadsParallel = l
		}

		newThreads := make([]*Thread, 0, len(threads))

		for _, t := range threads {
			children := t.next()
			stats.Cycles++

		childloop:
			for _, child := range children {
			dedupe:
				for _, newThread := range newThreads {
					if newThread.PC != child.PC || newThread.I != child.I || len(newThread.Stack) != len(child.Stack) {
						continue dedupe
					}
					for i := range child.Stack {
						if child.Stack[i] != newThread.Stack[i] {
							continue dedupe
						}
					}
					stats.ThreadsDeduped++
					continue childloop
				}

				if child.ID == 0 {
					threadID++
					child.ID = threadID
				}
				if child.Match {
					stats.ThreadsExecuted = threadID
					fmt.Printf("%#v\n", stats)
					return child
				}
				newThreads = append(newThreads, child)
			}
		}
		threads = newThreads
	}
	stats.ThreadsExecuted = threadID
	return nil
}

func DecodeStack(stack []int, p []Ins) []string {
	d := make([]string, len(stack))
	for i, pc := range stack {
		d[i] = p[pc-1].(OpCall).Name
	}
	return d
}

func (t *Thread) next() []*Thread {
	self := []*Thread{t}
	op := t.P[t.PC]
	//fmt.Printf("t=% 3d pc=% 3d %#v\n", t.ID, t.PC, op)
	switch op := op.(type) {
	case OpFunc:
		t.PC++
		return self
	case OpCall:
		newPC := op.PC
		if newPC == -1 {
			newPC = callPC(t.P, op)
		}
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
		if len(t.S)-t.I == 0 {
			return nil
		} else if t.S[t.I] < op.Start || t.S[t.I] > op.End {
			return nil
		}
		t.PC++
		t.capture(string(t.S[t.I]))
		t.I++
		return self
	case OpString:
		if len(t.S)-t.I < len(op.Value) {
			return nil
		} else if t.S[t.I:t.I+len(op.Value)] != op.Value {
			return nil
		}
		t.PC++
		t.capture(op.Value)
		t.I += len(op.Value)
		return self
	case OpMatch:
		if len(t.S)-t.I == 0 {
			t.Match = true
			return self
		}
		return nil
	case OpJmp:
		t.PC += op.N
		return self
	case OpFork:
		if strings.HasPrefix(t.S[t.I:], op.Cond) {
			clone := t.Clone()
			t.PC += 1
			clone.PC += op.N
			return append(self, clone)
		}
		t.PC += 1
		return self
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

func Alt(alts ...[]Ins) []Ins {
	var a []Ins
	for _, b := range alts {
		if a == nil {
			a = b
			continue
		}
		fork := OpFork{N: len(a) + 2}
		jmp := OpJmp{len(b) + 1}
		a = append(append(append([]Ins{fork}, a...), jmp), b...)
	}
	return a
}

func Plus(p []Ins) []Ins {
	return append(p, OpFork{N: -len(p)})
}

func Star(p []Ins) []Ins {
	fork := OpFork{N: len(p) + 2}
	jmp := OpJmp{-(len(p) + 1)}
	return append(append([]Ins{fork}, p...), jmp)
}

func QuestionMark(p []Ins) []Ins {
	fork := OpFork{N: len(p) + 1}
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
	return []Ins{OpCall{Name: name, PC: -1}}
}

func Graphviz(p []Ins) []byte {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "digraph g {\n")
	var rows []string
	var conns []string
	for i, ins := range p {
		row := []string{fmt.Sprintf("%d", i)}
		var label string
		switch op := ins.(type) {
		case OpRange:
			row = append(row, "range", string(op.Start), string(op.End))
			//label = fmt.Sprintf("range %q %q", string(op.Start), string(op.End))
		case OpString:
			row = append(row, "string", op.Value)
			//label = fmt.Sprintf("string %q", op.Value)
		case OpFork:
			row = append(row, "fork", fmt.Sprintf("%d", i+op.N))
			//label = fmt.Sprintf("fork %d", i+op.N)
			//fmt.Fprintf(buf, "%d -> %d;", i, i+op.N)
		case OpJmp:
			row = append(row, "jmp", fmt.Sprintf("%d", i+op.N))
			conns = append(conns, fmt.Sprintf("abc:%d:w -> abc:%d:w;", i, i+op.N))
			//label = fmt.Sprintf("jmp %d", i+op.N)
			//fmt.Fprintf(buf, "%d -> %d;", i, i+op.N)
		case OpCapture:
			//start := "start"
			//if !op.Start {
			//start = "end"
			//}
			//fmt.Printf("capture %s %s", start, op.Name)
		case OpMatch:
			row = append(row, "match")
			//label = fmt.Sprintf("match")
		case OpFunc:
			row = append(row, "func", op.Name)
			//label = fmt.Sprintf("func %s", op.Name)
			//indent++
		case OpCall:
			row = append(row, "call", op.Name)
			//label = fmt.Sprintf("call %s", op.Name)
			//conns = append(conns, fmt.Sprintf("abc:%d -> abc:%d;", i, callPC(p, op)))
		case OpReturn:
			row = append(row, "return")
			//label = fmt.Sprintf("return")
		default:
			panic("unreachable")
		}
		//label = strings.Replace(label, "{", "", -1)
		//label = strings.Replace(label, "}", "", -1)
		//label = strings.Replace(label, `"`, "", -1)
		//label = strings.Replace(label, `\`, "", -1)
		//label = fmt.Sprintf("{ %d | %s}", i, label)
		//label = fmt.Sprintf("{ %03d | %s}", i, label)
		//label = strings.Replace(label, "}", "\\}", -1)
		//label = strings.Replace(label, `"`, `\"`, -1)
		for j, col := range row {
			row[j] = fmt.Sprintf(`<TD port="%d" align="left">%s</TD>`, i, col)
		}
		label = `<TR>` + strings.Join(row, "") + `</TR>`
		rows = append(rows, label)
		//label = fmt.Sprintf("[%d] %s", i, label)
		//fmt.Fprintf(buf, "%d[label=%q];", i, label)
		//if i > 0 {
		//fmt.Fprintf(buf, "%d -> %d;", i-1, i)
		//}
	}
	fmt.Fprintf(buf, `
	abc [shape=none, margin=0, label=<
<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" CELLPADDING="4">
%s
</TABLE>>];
%s
`, strings.Join(rows, "\n"), strings.Join(conns, "\n"))
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
			fmt.Printf("fork %d %q", i+op.N, op.Cond)
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
			fmt.Printf("call %s %d", op.Name, op.PC)
		case OpReturn:
			fmt.Printf("return")
		default:
			panic("unreachable")
		}
		fmt.Print("\n")
	}
}

func Optimize(p []Ins) []Ins {
	for i, ins := range p {
		switch op := ins.(type) {
		case OpFork:
			p[i] = optimizeFork(i+op.N, op, p)
		case OpCall:
			op.PC = callPC(p, op)
			p[i] = op
		}
	}
	return p
}

func optimizeFork(pc int, fork OpFork, p []Ins) OpFork {
	for pc < len(p) {
		switch op := p[pc].(type) {
		case OpString:
			fork.Cond = op.Value
			return fork
		case OpCall:
			pc = callPC(p, op)
		case OpFunc:
			pc++
		default:
			return fork
		}
	}
	return fork
}
