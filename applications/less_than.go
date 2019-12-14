package applications

import (
	"opcode"
)

var LessThan = &LessThanApp{}

type LessThanApp struct{}

func (a *LessThanApp) Opcode() int {
	return 7
}

func (a *LessThanApp) Exec(os opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
	val1 := os.Memory().GetAt(cursor+1, c.Param1Mode)
	val2 := os.Memory().GetAt(cursor+2, c.Param2Mode)
	ptr := os.Memory().GetIndex(cursor+3, c.Param3Mode)

	if val1 < val2 {
		os.Debug("%02d (lt): %d < %d\n", c.Code, val1, val2)
		os.Debug("\t%d was %d, now %d\n", ptr, os.Memory().GetIndex(ptr, opcode.PositionMode), 1)

		os.Memory().Set(ptr, 1)
	} else {

		os.Debug("%02d (lt): %d >= %d\n", c.Code, val1, val2)
		os.Debug("\t%d was %d, now %d\n", ptr, os.Memory().GetIndex(ptr, opcode.PositionMode), 0)

		os.Memory().Set(ptr, 0)
	}

	return opcode.IntP(cursor + 4), nil
}
