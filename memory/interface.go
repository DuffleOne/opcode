package memory

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PositionMode = 0
	ImmediateMde = 1
	RelativeMode = 2
)

type Memory interface {
	GetAt(int, int) int
	GetIndex(int, int) int
	Set(int, int)
	IncRelativeBase(int)
	GetRelativeBase() int
	Next() (int, bool)
	Jump(int)
	All() []int
}

func handleBaseMemory(in interface{}) ([]int, error) {
	if in == nil {
		return []int{99}, nil
	}

	switch v := in.(type) {
	case string:
		var mem []int

		parts := strings.Split(v, ",")

		for _, p := range parts {
			n, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				return nil, err
			}

			mem = append(mem, n)
		}

		return mem, nil
	case []int:
		return v, nil
	case []byte:
		return handleBaseMemory(string(v))
	default:
		return nil, fmt.Errorf("cannot handle input memory type")
	}
}
