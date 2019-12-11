package applications

import (
	"fmt"
	"opcode"
)

func makeAdjustRelativeBase() *opcode.Application {
	app := opcode.MakeApp(9)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		param := os.Memory.GetAt(cursor+1, c.Param1Mode)

		if os.Debug {
			fmt.Printf("%02d (arb): %d (crb) + %d\n", c.Code, os.Memory.RelativeBase, param)
		}

		os.Memory.RelativeBase += param

		if os.Debug {
			fmt.Printf("%02d (arb): relativeBase now %d\n", c.Code, os.Memory.RelativeBase)
		}

		return opcode.IntP(cursor + 2), nil
	}

	return app
}

var AdjustRelativeBase *opcode.Application = makeAdjustRelativeBase()
