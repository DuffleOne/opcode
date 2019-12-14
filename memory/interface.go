package memory

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
