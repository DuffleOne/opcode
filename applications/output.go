package applications

import (
	"fmt"
	"opcode"
)

func makeOutput() *opcode.Application {
	app := opcode.MakeApp(4)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		val := os.Memory.GetAt(cursor+1, c.Param1Mode)

		fmt.Println(val)

		return opcode.IntP(cursor + 2), nil
	}

	return app
}

var Output *opcode.Application = makeOutput()
