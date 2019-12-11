package opcode

import (
	"fmt"
	"strconv"
	"strings"
)

type OS struct {
	Debug        bool
	Memory       *MemoryStore
	Applications map[int]*Application
	stdOut       []string
}

func BootFromString(start string, apps []*Application, debug bool) (*OS, error) {
	var mem []int

	parts := strings.Split(start, ",")

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		mem = append(mem, n)
	}

	return Boot(mem, apps, debug), nil
}

func Boot(memoryToLoad []int, apps []*Application, debug bool) *OS {
	ms := NewMemStore(memoryToLoad, IntP(2048))

	maps := map[int]*Application{}

	os := &OS{
		Debug:  debug,
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

func (os *OS) StdOut(seperator string) string {
	return strings.Join(os.stdOut, seperator)
}

func (os *OS) WriteOut() {
	for _, s := range os.stdOut {
		fmt.Println(s)
	}
}

func (os *OS) Run() error {
	for {
		addr, b := os.Memory.Next()
		opcode, err := BuildOPCode(os.Memory.GetIndex(addr, PositionMode))
		if err != nil {
			return err
		}
		if os.Debug {
			fmt.Printf("cur: %d - current opcode: %02d, %s\n", addr, opcode.Code, opcode)
		}
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
			return fmt.Errorf("cannot find application for opcode %d", opcode.Code)
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

	for k, i := range os.Memory.All() {
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

func IntP(i int) *int {
	return &i
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
