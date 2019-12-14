package applications

import (
	"opcode"
)

var Output = &OutputApp{}

type OutputApp struct{}

func (a *OutputApp) Opcode() int {
	return 4
}

func (a *OutputApp) Exec(os opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
	val := os.Memory().GetAt(cursor+1, c.Param1Mode)

	os.Println(val)

	os.Debug("%02d (output): val: %d, ptr: %d\n", c.Code, val, cursor+1)

	return opcode.IntP(cursor + 2), nil
}
