package applications

import (
	"fmt"
	"opcode"
)

func makeLessThan() *opcode.Application {
	app := opcode.MakeApp(7)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		val1 := os.Memory.GetAt(cursor+1, c.Param1Mode)
		val2 := os.Memory.GetAt(cursor+2, c.Param2Mode)
		ptr := os.Memory.GetIndex(cursor+3, c.Param3Mode)

		if val1 < val2 {
			if os.Debug {
				fmt.Printf("%02d (lt): %d < %d\n", c.Code, val1, val2)
				fmt.Printf("\t%d was %d, now %d\n", ptr, os.Memory.GetIndex(ptr, opcode.PositionMode), 1)
			}
			os.Memory.Set(ptr, 1)
		} else {
			if os.Debug {
				fmt.Printf("%02d (lt): %d >= %d\n", c.Code, val1, val2)
				fmt.Printf("\t%d was %d, now %d\n", ptr, os.Memory.GetIndex(ptr, opcode.PositionMode), 0)
			}
			os.Memory.Set(ptr, 0)
		}

		return opcode.IntP(cursor + 4), nil
	}

	return app
}

var LessThan *opcode.Application = makeLessThan()
