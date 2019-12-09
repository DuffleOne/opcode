package applications

import (
	"opcode"
)

func makeJumpIfFalse() *opcode.Application {
	app := opcode.MakeApp(6)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		val := os.Memory.GetAt(cursor+1, c.Param1Mode)

		if val != 0 {
			return opcode.IntP(cursor + 3), nil
		}

		newPtr := os.Memory.GetAt(cursor+2, c.Param2Mode)

		return opcode.IntP(newPtr), nil
	}

	return app
}

var JumpIfFalse *opcode.Application = makeJumpIfFalse()
