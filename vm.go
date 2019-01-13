package vm

import (
	"fmt"
	"io"
)

func NewVM(p Program) *VM {
	funcs := map[string]int{}
	for pc, op := range p {
		if opT, ok := op.(OpFunc); ok {
			funcs[opT.Name] = pc + 1
		}
	}

	v := &VM{p: p, funcs: funcs}
	v.threads = v.addThread(v.threads, &thread{})

	return v
}

type VM struct {
	// n is the amount of bytes the VM has accepted via Write()
	n int
	// threads is the list of active threads
	threads []*thread
	stats   Stats
	pcs     map[int]struct{}
	p       Program
	funcs   map[string]int
}

func (v *VM) Write(data []byte) (int, error) {
	for i, c := range data {
		var nextThreads []*thread
		v.pcs = map[int]struct{}{}
		if l := len(v.threads); l > v.stats.Concurrency {
			v.stats.Concurrency = l
		}

		for _, t := range v.threads {
			v.stats.Ops++
			// If the thread is already at the end of the program, an additional
			// char will kill it.
			if t.pc >= len(v.p) {
				continue
			}

			op := v.p[t.pc]
			switch opT := op.(type) {
			case OpString:
				// If char doesn't match, kill this thread.
				if opT.Value[t.oc] != c {
					continue
				}

				// Increment our op counter (offset) into this string.
				t.oc++
				if t.oc >= len(opT.Value) {
					// We're done matching the string, move on to the next op and
					// reset our op counter (offset).
					t.pc++
					t.oc = 0
				}
			case OpRange:
				if c < opT.Start || c > opT.End {
					continue
				}

				t.pc++
			default:
				return v.n + i, fmt.Errorf("unknown op: %T", op)
			}

			nextThreads = v.addThread(nextThreads, t)
			v.pcs[t.pc] = struct{}{}
		}

		v.threads = nextThreads

		if len(nextThreads) == 0 {
			v.n += i
			return i, io.ErrShortWrite
		}
	}
	n := len(data)
	v.n += n
	return n, nil
}

func (v *VM) addThread(threads []*thread, t *thread) []*thread {
loop:
	for t.pc < len(v.p) {
		v.stats.Ops++
		op := v.p[t.pc]
		switch opT := op.(type) {
		case OpFork:
			stack := make([]int, len(t.stack))
			copy(stack, t.stack)
			fork := &thread{pc: t.pc + opT.PC, stack: stack}
			threads = v.addThread(threads, fork)
			t.pc++
			v.stats.Forks++
		case OpJump:
			t.pc += opT.PC
		case OpCall:
			t.stack = append(t.stack, t.pc+1)
			t.pc = v.funcs[opT.Name]
		case OpReturn:
			t.pc = t.stack[len(t.stack)-1]
			t.stack = t.stack[0 : len(t.stack)-1]
		case OpHalt:
			t.pc = len(v.p)
		case OpString:
			if opT.Value != "" {
				break loop
			}
			t.pc++
		default:
			break loop
		}
	}

	if _, ok := v.pcs[t.pc]; ok {
		return threads
	}

	return append(threads, t)
}

func (v *VM) Valid() bool {
	for _, t := range v.threads {
		v.stats.Ops++

		if t.pc == len(v.p) {
			return true
		}
	}
	return false
}

func (v *VM) Stats() Stats {
	return v.stats
}

type thread struct {
	// pc is the program counter, i.e. an offset into the vm's program.
	pc int
	// oc is the operation counter, i.e. an offset into a multi-character op, e.g
	// OpString.
	oc int
	// stack contains the pc values of OpCall ops that haven't been returned to yet.
	stack []int
}
