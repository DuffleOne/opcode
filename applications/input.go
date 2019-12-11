package applications

import (
	"bufio"
	"fmt"
	"opcode"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func makeInput() *opcode.Application {
	app := opcode.MakeApp(3)

	app.Exec = func(os *opcode.OS, c *opcode.OPCode, cursor int) (*int, error) {
		fmt.Print(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		input = strings.TrimSpace(input)

		val, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}

		p := os.Memory.GetIndex(cursor+1, c.Param1Mode)

		if os.Debug {
			fmt.Printf("%02d (input): val: %d, ptr: %d\n", c.Code, val, p)
			fmt.Printf("\t%d was %d, now %d\n", p, os.Memory.GetIndex(p, opcode.PositionMode), val)
		}

		os.Memory.Set(p, val)

		return opcode.IntP(cursor + 2), nil
	}

	return app
}

var Input *opcode.Application = makeInput()
