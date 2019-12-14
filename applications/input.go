package applications

import (
	"errors"
	"fmt"

	"opcode"
)

var Input = &InputApp{}

type InputApp struct{}

func (a *InputApp) Opcode() int {
	return 3
}

func (a *InputApp) Exec(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
	val, ok := os.InputHandler.GetInput()

	if !ok {
		return nil, errors.New("no input provided")
	}

	p := os.Memory.GetIndex(cursor+1, c.Param1Mode)

	if os.Debug {
		fmt.Printf("%02d (input): val: %d, ptr: %d\n", c.Code, val, p)
		fmt.Printf("\t%d was %d, now %d\n", p, os.Memory.GetIndex(p, opcode.PositionMode), val)
	}

	os.Memory.Set(p, val)

	return opcode.IntP(cursor + 2), nil
}
