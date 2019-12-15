package opcode

import (
	"errors"
	"opcode/memory"
	"opcode/opcode"
)

var ErrHalt = errors.New("halt")

func IntP(i int) *int {
	return &i
}

type Application interface {
	Opcode() int
	Exec(os OS, opcode *opcode.OPCode, cursor int) (*int, error)
}

type OS interface {
	Debug(string, ...interface{})
	Memory() memory.Memory
	Run() error
	Dump(*int) string
	GetInput() (int, bool)
	GetStdOut(string) string
	Println(int)
	Printf(string, ...interface{})
}
