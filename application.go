package opcode

type Application interface {
	Opcode() int
	Exec(os *OS, opcode *OPCode, cursor int) (*int, error)
}
