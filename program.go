package vm

import (
	"fmt"
	"math"
	"strings"
)

type Program []Op

func (p Program) String() string {
	buf := &strings.Builder{}
	prefix := fmt.Sprintf("%%%dd: ", int(1+math.Log10(float64(len(p)))))
	for pc, op := range p {
		fmt.Fprintf(buf, prefix, pc)
		switch opT := op.(type) {
		case OpString:
			fmt.Fprintf(buf, "string %q", opT.Value)
		case OpJump:
			fmt.Fprintf(buf, "jump %d", opT.PC+pc)
		case OpFork:
			fmt.Fprintf(buf, "fork %d", opT.PC+pc)
		case OpRange:
			fmt.Fprintf(buf, "range %q %q", string(opT.Start), string(opT.End))
		case OpFunc:
			fmt.Fprintf(buf, "func %q", opT.Name)
		case OpCall:
			fmt.Fprintf(buf, "call %q", opT.Name)
		case OpHalt:
			fmt.Fprintf(buf, "halt")
		case OpReturn:
			fmt.Fprintf(buf, "return")
		default:
			panic(fmt.Errorf("unknown op: %#v", opT))
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func String(value string) Program {
	return Program{OpString{Value: value}}
}

func Concat(programs ...Program) Program {
	var r Program
	for _, p := range programs {
		r = append(r, p...)
	}
	return r
}

func ZeroOrMore(p Program) Program {
	fork := OpFork{PC: len(p) + 2}
	jmp := OpJump{PC: -(len(p) + 1)}
	return append(append(Program{fork}, p...), jmp)
}

func ZeroOrOne(p Program) Program {
	fork := OpFork{PC: len(p) + 1}
	return append(Program{fork}, p...)
}

func OneOrMore(p Program) Program {
	fork := OpFork{PC: -len(p)}
	return append(p, fork)
}

func Alt(alts ...Program) Program {
	var a Program
	for _, b := range alts {
		if a == nil {
			a = b
			continue
		}
		fork := OpFork{PC: len(a) + 2}
		jmp := OpJump{PC: len(b) + 1}
		a = append(append(append(Program{fork}, a...), jmp), b...)
	}
	return a
}

func Repeat(min, max int, p Program) Program {
	var newP Program
	for i := 0; i < min; i++ {
		newP = Concat(newP, p)
	}
	for i := 0; i < max-min; i++ {
		newP = Concat(newP, ZeroOrOne(p))
	}
	return newP
}

func Range(start byte, end byte) Program {
	return Program{OpRange{Start: start, End: end}}
}

func Func(name string, p Program) Program {
	return append(append(Program{OpFunc{Name: name}}, p...), OpReturn{})
}

func Call(name string) Program {
	return Program{OpCall{Name: name}}
}

func Halt() Program {
	return Program{OpHalt{}}
}
