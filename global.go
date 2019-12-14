package opcode

import (
	"errors"
)

var ErrHalt = errors.New("halt")

func IntP(i int) *int {
	return &i
}

type Application interface {
	Opcode() int
	Exec(os OS, opcode *OPCode, cursor int) (*int, error)
}

type OS interface {
	GetInput() (int, bool)
	Debug(string, ...interface{})
	Memory() *MemoryStore
	Println(interface{})
	StdOut(string) string
	WriteOut()
	Run() error
	Dump(*int) string
}
