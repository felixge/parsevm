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
}

func (v *VM) Write(data []byte) (int, error) {
	for i, c := range data {
		var nextThreads []*thread
	currentThreads:
		for i := 0; i < len(v.threads); i++ {
			t := v.threads[i]
			// Execute the thread until it consumes c or dies.
			for {
				// If the thread is already at the end of the program, an additional
				// char will kill it.
				if t.pc >= len(t.p) {
					continue currentThreads
				}

				op := t.p[t.pc]
				v.stats.Ops++
				switch opT := op.(type) {
				case OpString:
					// If char doesn't match, kill this thread.
					if opT.Value[t.oc] != c {
						continue currentThreads
					}

					// Increment our op counter (offset) into this string.
					t.oc++
					if t.oc == len(opT.Value) {
						// We're done matching the string, move on to the next op and
						// reset our op counter (offset).
						t.pc++
						t.oc = 0
					}

					nextThreads = v.addThread(nextThreads, t)
					continue currentThreads
				default:
					return v.n + i, fmt.Errorf("unknown op: %s", op)
				}
			}
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
		op := t.p[t.pc]

		switch opT := op.(type) {
		case OpFork:
			fork := &thread{p: t.p, pc: t.pc + opT.PC}
			threads = v.addThread(threads, fork)
			t.pc++
			v.stats.Forks++
		case OpJmp:
			t.pc += opT.PC
		default:
			break loop
		}
		v.stats.Ops++
	}
	for _, existingT := range threads {
		v.stats.Ops++
		if existingT.pc == t.pc {
			return threads
		}
	}
	return append(threads, t)
}

func (v *VM) Valid() bool {
	for _, t := range v.threads {
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
