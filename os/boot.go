package os

import (
	"fmt"
	"opcode"
	"strconv"
	"strings"
)

func BootFromString(debug bool, ih *opcode.InputHandler, apps []opcode.Application, startMemory string) (*OS, error) {
	var mem []int

	parts := strings.Split(startMemory, ",")

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		mem = append(mem, n)
	}

	return Boot(debug, ih, apps, mem), nil
}

func Boot(debug bool, ih *opcode.InputHandler, apps []opcode.Application, startMemory []int) *OS {
	var err error

	ms := opcode.NewMemStore(startMemory, opcode.IntP(2048))

	maps := map[int]opcode.Application{}

	if ih == nil {
		ih, err = opcode.NewInputHandler(opcode.ImmediateInputMode, nil)
		if err != nil {
			panic(fmt.Errorf("cannot make default input handler: %w", err))
		}
	}

	os := &OS{
		debug:        debug,
		memory:       ms,
		stdOut:       []string{},
		inputHandler: ih,
	}

	for _, app := range apps {
		code := app.Opcode()
		maps[code] = app
	}

	os.Applications = maps

	return os
}
