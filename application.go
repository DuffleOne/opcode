package opcode

type Application struct {
	opcode int
	Exec   func(os *OS, opcode *OPCode, cursor int) (*int, error)
}

func (a *Application) OPCode() int {
	return a.opcode
}

func MakeApp(opcode int) *Application {
	return &Application{
		opcode: opcode,
	}
}
