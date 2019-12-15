package os

import (
	"fmt"

	"opcode"
	"opcode/memory"
)

const DefaultMemorySize = 2048

type OSBootParams struct {
	Debug         bool
	Memory        memory.Memory
	InputHandler  *InputHandler
	OutputHandler *OutputHandler
	Applications  []opcode.Application
}

func Boot(params OSBootParams) (*OS, error) {
	var err error
	mapApps := map[int]opcode.Application{}

	for _, app := range params.Applications {
		code := app.Opcode()
		mapApps[code] = app
	}

	ih := params.InputHandler
	oh := params.OutputHandler

	if ih == nil {
		ih, err = NewInputHandler(ImmediateInputMode, nil)
		if err != nil {
			return nil, fmt.Errorf("cannot make default input handler: %w", err)
		}
	}

	if oh == nil {
		oh, err = NewOutputHandler(ImmediateOutputMode)
		if err != nil {
			return nil, fmt.Errorf("cannot make default ouput handler: %w", err)
		}
	}

	return &OS{
		debug:         params.Debug,
		memory:        params.Memory,
		inputHandler:  ih,
		outputHandler: oh,
		Applications:  mapApps,
	}, nil
}
