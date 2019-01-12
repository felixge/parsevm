package vm

import "fmt"

type Op interface {
	String() string
}

type OpString struct{ Value string }

func (o OpString) String() string { return fmt.Sprintf("string %q", o.Value) }

type OpJmp struct{ PC int }

func (o OpJmp) String() string { return fmt.Sprintf("jmp %+d", o.PC) }

type OpFork struct{ PC int }

func (o OpFork) String() string { return fmt.Sprintf("fork %+d", o.PC) }
