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