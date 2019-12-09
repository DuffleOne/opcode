package applications

import (
	"opcode"
)

func makeEquals() *opcode.Application {
	app := opcode.MakeApp(8)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		val1 := os.Memory.GetAt(cursor+1, c.Param1Mode)
		val2 := os.Memory.GetAt(cursor+2, c.Param2Mode)

		if val1 == val2 {
			os.Memory.Set(os.Memory.Get(cursor+3), 1)
		} else {
			os.Memory.Set(os.Memory.Get(cursor+3), 0)
		}

		return opcode.IntP(cursor + 4), nil
	}

	return app
}

var Equals *opcode.Application = makeEquals()
