package applications

import (
	"opcode"
)

func makeMul() *opcode.Application {
	app := opcode.MakeApp(2)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		p1 := os.Memory.GetAt(cursor+1, c.Param1Mode)
		p2 := os.Memory.GetAt(cursor+2, c.Param2Mode)

		val := p1 * p2

		os.Memory.Set(os.Memory.Get(cursor+3), val)

		return opcode.IntP(cursor + 4), nil
	}

	return app
}

var Mul *opcode.Application = makeMul()
