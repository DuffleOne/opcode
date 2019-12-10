package applications

import (
	"opcode"
)

func makeAdjustRelativeBase() *opcode.Application {
	app := opcode.MakeApp(9)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		param := os.Memory.GetAt(cursor+1, c.Param1Mode)
		os.Memory.RelativeBase += param

		return opcode.IntP(cursor + 2), nil
	}

	return app
}

var AdjustRelativeBase *opcode.Application = makeAdjustRelativeBase()
