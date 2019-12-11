package opcode

type MemoryStore struct {
	cursor       int
	RelativeBase int
	v            []int
}

func NewMemStore(start []int, defaultSize *int) *MemoryStore {
	if defaultSize == nil {
		return &MemoryStore{
			v: start,
		}
	}

	size := *defaultSize

	s := make([]int, size)

	for k, v := range start {
		s[k] = v
	}

	return &MemoryStore{
		v: s,
	}
}

func (ms *MemoryStore) GetAt(pos, paramMode int) int {
	return ms.v[ms.GetIndex(pos, paramMode)]
}

func (ms *MemoryStore) GetIndex(pos, paramMode int) int {
	switch paramMode {
	case PositionMode:
		return ms.v[pos]
	case ImmediateMde:
		return pos
	case RelativeMode:
		return ms.v[pos] + ms.RelativeBase
	default:
		return 0
	}
}

func (ms *MemoryStore) Set(pos, val int) {
	ms.v[pos] = val
}

func (ms *MemoryStore) Next() (int, bool) {
	defer func() { ms.cursor++ }()

	if ms.cursor >= len(ms.v) {
		return ms.cursor, false
	}

	return ms.cursor, true
}

func (ms *MemoryStore) Jump(to int) {
	ms.cursor = to
}

func (ms *MemoryStore) All() []int {
	return ms.v
}
