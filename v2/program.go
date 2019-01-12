package vm

import "fmt"

type Program []Op

// TODO show fork/jmp with absolute pc
func (p Program) String() string {
	pcWidth := fmt.Sprintf("%%%dd", len(p)/10+1)
	var r string
	for pc, op := range p {
		r += fmt.Sprintf(pcWidth+": %s\n", pc, op)
	}
	return r
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
	jmp := OpJmp{PC: -(len(p) + 1)}
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
		jmp := OpJmp{PC: len(b) + 1}
		a = append(append(append(Program{fork}, a...), jmp), b...)
	}
	return a
}
