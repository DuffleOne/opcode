package applications

import (
	"opcode"
	opc "opcode/opcode"
)

var JumpIfTrue = &JumpIfTrueApp{}

type JumpIfTrueApp struct{}

func (a *JumpIfTrueApp) Opcode() int {
	return 5
}

func (a *JumpIfTrueApp) Exec(os opcode.OS, c *opc.OPCode, cursor int) (*int, error) {
	val := os.Memory().GetAt(cursor+1, c.Param1Mode)
	ptr := os.Memory().GetAt(cursor+2, c.Param2Mode)

	if val != 0 {
		os.Debug("%02d (jit): %d is truthy; jump to %d\n", c.Code, val, ptr)

		return opcode.IntP(ptr), nil
	}

	os.Debug("%02d (jit): %d is falsy; continue to %d\n", c.Code, val, cursor+3)

	return opcode.IntP(cursor + 3), nil
}
