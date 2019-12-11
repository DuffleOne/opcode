package applications

import (
	"fmt"
	"opcode"
)

func makeJumpIfTrue() *opcode.Application {
	app := opcode.MakeApp(5)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		val := os.Memory.GetAt(cursor+1, c.Param1Mode)
		ptr := os.Memory.GetAt(cursor+2, c.Param2Mode)

		if val != 0 {
			if os.Debug {
				fmt.Printf("%02d (jit): %d is truthy; jump to %d\n", c.Code, val, ptr)
			}
			return opcode.IntP(ptr), nil
		}

		if os.Debug {
			fmt.Printf("%02d (jit): %d is falsy; continue to %d\n", c.Code, val, cursor+3)
		}

		return opcode.IntP(cursor + 3), nil
	}

	return app
}

var JumpIfTrue *opcode.Application = makeJumpIfTrue()
