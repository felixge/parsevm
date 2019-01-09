package vm

import "fmt"

type Program []Op

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
