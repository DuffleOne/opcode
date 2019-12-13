package opcode

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ImmediateInputMode = 0
	CachedInputMode    = 1
)

type InputHandler struct {
	mode    int
	buffer  []int
	reader  *bufio.Reader
	LastErr error
}

func NewInputHandler(mode int, cache []int) (*InputHandler, error) {
	switch mode {
	case ImmediateInputMode:
		if cache != nil {
			return nil, errors.New("cannot provide a cache with immediate input mode")
		}
	case CachedInputMode:
		if cache == nil || len(cache) == 0 {
			return nil, errors.New("must provide a cache if using cached input mode")
		}
	default:
		return nil, errors.New("unsupported input mode")
	}

	return &InputHandler{
		mode:   mode,
		buffer: cache,
		reader: bufio.NewReader(os.Stdin),
	}, nil
}

func (i *InputHandler) SwitchMode(inputMode int) {
	i.mode = inputMode
}

func (i *InputHandler) GetInput() (int, bool) {
	switch i.mode {
	case ImmediateInputMode:
		fmt.Print(">> ")
		input, err := i.reader.ReadString('\n')
		if err != nil {
			i.LastErr = err
			return 0, false
		}

		input = strings.TrimSpace(input)

		val, err := strconv.Atoi(input)
		if err != nil {
			i.LastErr = err
			return 0, false
		}

		return val, true
	case CachedInputMode:
		if len(i.buffer) < 1 {
			return 0, false
		}

		val, rest := i.buffer[0], i.buffer[1:]
		i.buffer = rest

		return val, true
	default:
		i.LastErr = errors.New("unsupported input mode")
		return 0, false
	}
}
