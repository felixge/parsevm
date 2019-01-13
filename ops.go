package vm

type Op interface{}

type OpString struct {
	Value string
}

type OpJump struct {
	PC int
}

type OpFork struct {
	PC int
}

type OpRange struct {
	Start byte
	End   byte
}

type OpFunc struct {
	Name string
}

type OpReturn struct{}

type OpCall struct {
	Name string
}

type OpHalt struct{}
