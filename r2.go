package r2

import "fmt"

type Ins struct {
	Op OpCode
	C  byte
	X  int
	Y  int
}

type OpCode string

const (
	OpChar  OpCode = "char"
	OpMatch OpCode = "match"
	OpJmp   OpCode = "jmp"
	OpSplit OpCode = "split"
)

func Run(s string, p []*Ins, pc int) bool {
	ins := p[pc]
	switch ins.Op {
	case OpChar:
		if len(s) == 0 || s[0] != ins.C {
			return false
		}
		return Run(s[1:], p, pc+1)
	case OpMatch:
		return true
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
	ins := make([]*Ins, len(s))
	for i := 0; i < len(s); i++ {
		ins[i] = &Ins{Op: OpChar, C: s[i]}
	}
	return ins
}

func Concat(a, b []*Ins) []*Ins {
	return append(a, b...)
}

func Alt(a, b []*Ins) []*Ins {
	split := &Ins{Op: OpSplit, X: 1, Y: len(a) + 2}
	jmp := &Ins{Op: OpJmp, X: len(b) + 1}
	return append(append(append([]*Ins{split}, a...), jmp), b...)
}

func Plus(p []*Ins) []*Ins {
	return append(p, &Ins{Op: OpSplit, X: -len(p), Y: 1})
}

func Print(p []*Ins) {
	for i, ins := range p {
		fmt.Printf("% 3d: %s", i, ins.Op)
		switch ins.Op {
		case OpChar:
			fmt.Print(" " + string(ins.C))
		case OpSplit:
			fmt.Printf(" %d %d", i+ins.X, i+ins.Y)
		case OpJmp:
			fmt.Printf(" %d", i+ins.X)
		}
		fmt.Print("\n")
	}
}
