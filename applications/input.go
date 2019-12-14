package applications

import (
	"errors"
	"opcode/memory"

	"opcode"
)

var Input = &InputApp{}

type InputApp struct{}

func (a *InputApp) Opcode() int {
	return 3
}

func (a *InputApp) Exec(os opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
	val, ok := os.GetInput()

	if !ok {
		return nil, errors.New("no input provided")
	}

	p := os.Memory().GetIndex(cursor+1, c.Param1Mode)

	os.Debug("%02d (input): val: %d, ptr: %d\n", c.Code, val, p)
	os.Debug("\t%d was %d, now %d\n", p, os.Memory().GetIndex(p, memory.PositionMode), val)

	os.Memory().Set(p, val)

	return opcode.IntP(cursor + 2), nil
}
