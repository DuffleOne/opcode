package applications

import (
	"opcode"
)

func makeHalt() *opcode.Application {
	app := opcode.MakeApp(99)

	app.Exec = func(os *opcode.OS, _ *opcode.OPCode, cursor int) (*int, error) {
		return opcode.IntP(cursor + 1), opcode.ErrHalt
	}

	return app
}

var Halt *opcode.Application = makeHalt()
