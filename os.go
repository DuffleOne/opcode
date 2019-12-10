package opcode

import (
	"fmt"
	"strconv"
	"strings"
)

type OS struct {
	Memory       *MemoryStore
	Applications map[int]*Application
	stdOut       []string
}

func BootFromString(start string, apps []*Application) (*OS, error) {
	var mem []int

	parts := strings.Split(start, ",")

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		mem = append(mem, n)
	}

	return Boot(mem, apps), nil
}

func Boot(memoryToLoad []int, apps []*Application) *OS {
	ms := newMemStore(memoryToLoad)

	maps := map[int]*Application{}

	os := &OS{
		Memory: ms,
		stdOut: []string{},
	}

	for _, app := range apps {
		maps[app.OPCode()] = app
	}

	os.Applications = maps

	return os
}

func (os *OS) Println(o interface{}) {
	switch v := o.(type) {
	case int:
		os.stdOut = append(os.stdOut, strconv.Itoa(v))
	case string:
		os.stdOut = append(os.stdOut, v)
	default:
		panic(fmt.Errorf("cannot push item of %t to []string", o))
	}
}

func (os *OS) StdOut() string {
	return strings.Join(os.stdOut, "\n")
}

func (os *OS) WriteOut() {
	for _, s := range os.stdOut {
		fmt.Println(s)
	}
}

func (os *OS) Run() error {
	for {
		addr, b := os.Memory.Next()
		opcode, err := BuildOPCode(os.Memory.Get(addr))
		if err != nil {
			return err
		}
		val := os.Memory.Get(addr)
		if app, ok := os.Applications[opcode.Code]; ok {
			jumpto, err := app.Exec(os, opcode, addr)
			if err != nil {
				if err == ErrHalt {
					break
				}

				return err
			}

			if jumpto != nil {
				os.Memory.Jump(*jumpto)
			}
		} else {
			return fmt.Errorf("cannot find application for opcode %d", val)
		}

		if !b {
			return fmt.Errorf("eof without halt")
		}
	}

	return nil
}

func (os *OS) Dump() string {
	var all []string

	for _, i := range os.Memory.All() {
		all = append(all, strconv.Itoa(i))
	}

	return strings.Join(all, ",")
}

func IntP(i int) *int {
	return &i
}