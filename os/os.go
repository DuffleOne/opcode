package os

import (
	"fmt"
	"strconv"
	"strings"

	"opcode"
	"opcode/memory"
	opc "opcode/opcode"
)

type OS struct {
	debug         bool
	memory        memory.Memory
	Applications  map[int]opcode.Application
	inputHandler  *InputHandler
	outputHandler *OutputHandler
}

func (os *OS) Memory() memory.Memory {
	return os.memory
}

func (os *OS) GetInput() (int, bool) {
	return os.inputHandler.GetInput()
}

func (os *OS) GetStdOut(seperator string) string {
	return os.outputHandler.GetStdOut(seperator)
}

func (os *OS) Printf(format string, args ...interface{}) {
	os.outputHandler.Printf(format, args...)
}

func (os *OS) Println(arg int) {
	os.outputHandler.Println(arg)
}

func (os *OS) Debug(format string, args ...interface{}) {
	if os.debug {
		fmt.Printf(format, args...)
	}
}

func (os *OS) Run() error {
	for {
		addr, b := os.memory.Next()
		opc, err := opc.BuildOPCode(os.memory.GetIndex(addr, memory.PositionMode))
		if err != nil {
			return err
		}

		os.Debug("cur: %d - current opcode: %02d, %s\n", addr, opc.Code, opc)

		if app, ok := os.Applications[opc.Code]; ok {
			jumpto, err := app.Exec(os, opc, addr)
			if err != nil {
				if err == opcode.ErrHalt {
					break
				}

				return err
			}

			if jumpto != nil {
				os.memory.Jump(*jumpto)
			}
		} else {
			return fmt.Errorf("cannot find application for opcode %d", opc.Code)
		}

		if !b {
			return fmt.Errorf("eof without halt")
		}
	}

	return nil
}

func (os *OS) Dump(trimBy *int) string {
	lastFilledByte := 0
	var use []string
	var all []string

	for k, i := range os.memory.All() {
		if i != 0 {
			lastFilledByte = k
		}
		all = append(all, strconv.Itoa(i))
	}

	if trimBy != nil {
		use = copyTo(all, lastFilledByte+*trimBy)
	}

	return strings.Join(use, ",")
}

func copyTo(in []string, to int) []string {
	var new []string

	if to+1 > len(in) || to+1 == len(in) {
		return in
	}

	for i := 0; i < to+1; i++ {
		v := in[i]
		new = append(new, v)
	}

	return new
}
