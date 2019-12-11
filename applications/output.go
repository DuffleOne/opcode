package applications

import (
	"fmt"
	"opcode"
)

func makeOutput() *opcode.Application {
	app := opcode.MakeApp(4)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		val := os.Memory.GetAt(cursor+1, c.Param1Mode)

		os.Println(val)

		if os.Debug {
			fmt.Printf("%02d (output): val: %d, ptr: %d\n", c.Code, val, cursor+1)
		}

		return opcode.IntP(cursor + 2), nil
	}

	return app
}

var Output *opcode.Application = makeOutput()
