package os

import (
	"errors"
	"fmt"
	"strings"
)

const (
	ImmediateOutputMode = 0
	CachedOutputMode    = 1
)

type OutputHandler struct {
	mode    int
	buffer  []string
	LastErr error
}

func NewOutputHandler(mode int) (*OutputHandler, error) {
	switch mode {
	case 0:
	case 1:
	default:
		return nil, errors.New("unsupported input mode")
	}

	return &OutputHandler{
		mode:   mode,
		buffer: []string{},
	}, nil
}

func (o *OutputHandler) Println(in int) {
	switch o.mode {
	case ImmediateOutputMode:
		fmt.Println(in)
	case CachedInputMode:
		o.buffer = append(o.buffer, fmt.Sprintf("%d", in))
	}
}

func (o *OutputHandler) Printf(format string, args ...interface{}) {
	switch o.mode {
	case ImmediateOutputMode:
		fmt.Printf(format, args...)
	case CachedInputMode:
		o.buffer = append(o.buffer, fmt.Sprintf(format, args...))
	}
}

func (o *OutputHandler) GetStdOut(seperator string) string {
	return strings.Join(o.buffer, seperator)
}
