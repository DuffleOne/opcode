package os

import (
	"fmt"

	"opcode"
	"opcode/memory"
)

const DefaultMemorySize = 2048

type OSBootParams struct {
	Debug        bool
	Memory       memory.Memory
	InputHandler *InputHandler
	Applications []opcode.Application
}

func Boot(params OSBootParams) (*OS, error) {
	var err error
	mapApps := map[int]opcode.Application{}

	for _, app := range params.Applications {
		code := app.Opcode()
		mapApps[code] = app
	}

	ih := params.InputHandler

	if ih == nil {
		ih, err = NewInputHandler(ImmediateInputMode, nil)
		if err != nil {
			return nil, fmt.Errorf("cannot make default input handler: %w", err)
		}
	}

	return &OS{
		debug:        params.Debug,
		memory:       params.Memory,
		stdOut:       []string{},
		inputHandler: params.InputHandler,
		Applications: mapApps,
	}, nil
}
