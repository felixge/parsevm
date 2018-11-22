package r1

type Ins struct {
	Op OpCode
	C  byte
	X  Program
	Y  Program
}

type OpCode string

const (
	OpChar  OpCode = "char"
	OpMatch OpCode = "match"
	OpJmp   OpCode = "jmp"
	OpSplit OpCode = "split"
)

type Program []*Ins

func (p Program) Match(s string) bool {
	ins := p[0]
	switch ins.Op {
	case OpChar:
		if len(s) == 0 || s[0] != ins.C {
			return false
		}
		p = p[1:]
		s = s[1:]
		return p.Match(s)
	case OpMatch:
		return true
	case OpJmp:
		return ins.X.Match(s)
	case OpSplit:
		if ins.X.Match(s) {
			return true
		}
		return ins.Y.Match(s)
	}
	panic("unreachable")
}
