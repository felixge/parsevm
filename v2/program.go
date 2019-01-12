package vm

import (
	"fmt"
	"math"
	"strings"
)

type Program []Op

// TODO show fork/jmp with absolute pc
func (p Program) String() string {
	buf := &strings.Builder{}
	prefix := fmt.Sprintf("%%%dd: ", int(math.Log10(float64(len(p)))))
	fmt.Println(prefix)
	for pc, op := range p {
		fmt.Fprintf(buf, prefix, pc)
		switch opT := op.(type) {
		case OpString:
			fmt.Fprintf(buf, "string %q", opT.Value)
		case OpJump:
			fmt.Fprintf(buf, "jump %d", opT.PC+pc)
		case OpFork:
			fmt.Fprintf(buf, "fork %d", opT.PC+pc)
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
