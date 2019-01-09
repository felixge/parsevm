package vm

import "fmt"

type Op interface {
	String() string
	//Execute(v *VM, t *Thread)
}

type OpString struct{ Value string }

func (o OpString) String() string { return fmt.Sprintf("string %q", o.Value) }
