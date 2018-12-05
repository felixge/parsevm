package r2

import "fmt"

type Ins struct {
	Op OpCode
	S  string
	X  int
	Y  int
}

type OpCode string

const (
	OpMatch  OpCode = "match"
	OpJmp    OpCode = "jmp"
	OpSplit  OpCode = "split"
	OpString OpCode = "string"
	OpRange  OpCode = "range"
)

func Run(s string, p []*Ins, pc int) bool {
	ins := p[pc]
	switch ins.Op {
	case OpRange:
		if len(s) == 0 {
			return false
		} else if s[0] < ins.S[0] || s[0] > ins.S[1] {
			return false
		}
		return Run(s[1:], p, pc+1)
	case OpString:
		if len(s) < len(ins.S) {
			return false
		} else if s[0:len(ins.S)] != ins.S {
			return false
		}
		return Run(s[len(ins.S):], p, pc+1)
	case OpMatch:
		return len(s) == 0
	case OpJmp:
		return Run(s, p, pc+ins.X)
	case OpSplit:
		if Run(s, p, pc+ins.X) {
			return true
		}
		return Run(s, p, pc+ins.Y)
	}
	panic("unreachable")
}

func Match(p []*Ins) []*Ins {
	return append(p, &Ins{Op: OpMatch})
}

func String(s string) []*Ins {
	return []*Ins{&Ins{Op: OpString, S: s}}
}

func Concat(parts ...[]*Ins) []*Ins {
	var newP []*Ins
	for _, p := range parts {
		newP = append(newP, p...)
	}
	return newP
}

func Alt(a, b []*Ins) []*Ins {
	split := &Ins{Op: OpSplit, X: 1, Y: len(a) + 2}
	jmp := &Ins{Op: OpJmp, X: len(b) + 1}
	return append(append(append([]*Ins{split}, a...), jmp), b...)
}

func Plus(p []*Ins) []*Ins {
	return append(p, &Ins{Op: OpSplit, X: -len(p), Y: 1})
}

func Star(p []*Ins) []*Ins {
	split := &Ins{Op: OpSplit, X: 1, Y: len(p) + 2}
	jmp := &Ins{Op: OpJmp, X: -(len(p) + 1)}
	return append(append([]*Ins{split}, p...), jmp)
}

func QuestionMark(p []*Ins) []*Ins {
	split := &Ins{Op: OpSplit, X: 1, Y: len(p) + 1}
	return append([]*Ins{split}, p...)
}

func Repeat(min, max int, p []*Ins) []*Ins {
	var newP []*Ins
	for i := 0; i < min; i++ {
		newP = Concat(newP, p)
	}
	for i := 0; i < max-min; i++ {
		newP = Concat(newP, QuestionMark(p))
	}
	return newP
}

func Range(start byte, end byte) []*Ins {
	return []*Ins{&Ins{Op: OpRange, S: string(start) + string(end)}}
}

func Alpha() []*Ins {
	return Alt(Range('a', 'z'), Range('A', 'Z'))
}

func Capture(name string, p []*Ins) []*Ins {
	return p
}

func Print(p []*Ins) {
	for i, ins := range p {
		fmt.Printf("% 3d: %s", i, ins.Op)
		switch ins.Op {
		case OpRange:
			fmt.Printf(" %q %q", string(ins.S[0]), string(ins.S[1]))
		case OpString:
			fmt.Printf(" %q", ins.S)
		case OpSplit:
			fmt.Printf(" %d %d", i+ins.X, i+ins.Y)
		case OpJmp:
			fmt.Printf(" %d", i+ins.X)
		}
		fmt.Print("\n")
	}
}
