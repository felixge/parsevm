package vm

import (
	"fmt"
	"io"
)

func NewVM(p Program) *VM {
	v := &VM{}
	v.threads = v.addThread(v.threads, &thread{p: p})
	return v
}

type VM struct {
	// n is the amount of bytes the VM has accepted via Write()
	n int
	// threads is the list of active threads
	threads []*thread
	stats   Stats
	pcs     map[int]struct{}
}

func (v *VM) Write(data []byte) (int, error) {
	for i, c := range data {
		var nextThreads []*thread
		v.pcs = map[int]struct{}{}

		for i := 0; i < len(v.threads); i++ {
			v.stats.Ops++
			t := v.threads[i]
			// If the thread is already at the end of the program, an additional
			// char will kill it.
			if t.pc >= len(t.p) {
				continue
			}

			op := t.p[t.pc]
			switch opT := op.(type) {
			case OpString:
				// If char doesn't match, kill this thread.
				if opT.Value[t.oc] != c {
					continue
				}

				// Increment our op counter (offset) into this string.
				t.oc++
				if t.oc == len(opT.Value) {
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
				return v.n + i, fmt.Errorf("unknown op: %s", op)
			}

			nextThreads = v.addThread(nextThreads, t)
			v.pcs[t.pc] = struct{}{}
		}

		v.threads = nextThreads

		if len(nextThreads) == 0 {
			return v.n + i, io.ErrShortWrite
		}
	}
	n := len(data)
	v.n += n
	return n, nil
}

func (v *VM) addThread(threads []*thread, t *thread) []*thread {
loop:
	for t.pc < len(t.p) {
		v.stats.Ops++
		op := t.p[t.pc]
		switch opT := op.(type) {
		case OpFork:
			fork := &thread{p: t.p, pc: t.pc + opT.PC}
			threads = v.addThread(threads, fork)
			t.pc++
			v.stats.Forks++
		case OpJump:
			t.pc += opT.PC
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

		if t.pc == len(t.p) {
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
	// p is the program executed by the thread. The program is a set of ops, and
	// it's always the same for all threads of the same vm.
	p Program
}
