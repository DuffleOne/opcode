package applications

import (
	"opcode"
	opc "opcode/opcode"
)

var AdjustRelativeBase = &ARBApp{}

type ARBApp struct{}

func (a *ARBApp) Opcode() int {
	return 9
}

func (a *ARBApp) Exec(os opcode.OS, c *opc.OPCode, cursor int) (*int, error) {
	param := os.Memory().GetAt(cursor+1, c.Param1Mode)

	os.Debug("%02d (arb): %d (crb) + %d\n", c.Code, os.Memory().GetRelativeBase(), param)

	os.Memory().IncRelativeBase(param)

	os.Debug("%02d (arb): relativeBase now %d\n", c.Code, os.Memory().GetRelativeBase())

	return opcode.IntP(cursor + 2), nil
}
