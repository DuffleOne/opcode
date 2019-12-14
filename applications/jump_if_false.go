package applications

import (
	"opcode"
	opc "opcode/opcode"
)

var JumpIfFalse = &JumpIfFalseApp{}

type JumpIfFalseApp struct{}

func (a *JumpIfFalseApp) Opcode() int {
	return 6
}

func (a *JumpIfFalseApp) Exec(os opcode.OS, c *opc.OPCode, cursor int) (*int, error) {
	val := os.Memory().GetAt(cursor+1, c.Param1Mode)
	ptr := os.Memory().GetAt(cursor+2, c.Param2Mode)

	if val == 0 {
		os.Debug("%02d (jit): %d is falsy; jump to %d\n", c.Code, val, ptr)

		return opcode.IntP(ptr), nil
	}

	os.Debug("%02d (jit): %d is truthy; continue to %d\n", c.Code, val, cursor+3)

	return opcode.IntP(cursor + 3), nil
}
