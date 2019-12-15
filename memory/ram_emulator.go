package memory

type RAMEmulator struct {
	cursor       int
	relativeBase int
	v            []int
}

func NewRAMStore(startIn interface{}, defaultSize *int) (*RAMEmulator, error) {
	start, err := handleBaseMemory(startIn)
	if err != nil {
		return nil, err
	}

	if defaultSize == nil {
		return &RAMEmulator{
			v: start,
		}, nil
	}

	size := *defaultSize

	s := make([]int, size)

	for k, v := range start {
		s[k] = v
	}

	return &RAMEmulator{
		v: s,
	}, nil
}

func (ms *RAMEmulator) GetAt(pos, paramMode int) int {
	return ms.v[ms.GetIndex(pos, paramMode)]
}

func (ms *RAMEmulator) GetIndex(pos, paramMode int) int {
	switch paramMode {
	case PositionMode:
		return ms.v[pos]
	case ImmediateMde:
		return pos
	case RelativeMode:
		return ms.v[pos] + ms.relativeBase
	default:
		return 0
	}
}

func (ms *RAMEmulator) Set(pos, val int) {
	ms.v[pos] = val
}

func (ms *RAMEmulator) IncRelativeBase(i int) {
	ms.relativeBase += i
}

func (ms *RAMEmulator) GetRelativeBase() int {
	return ms.relativeBase
}

func (ms *RAMEmulator) Next() (int, bool) {
	defer func() { ms.cursor++ }()

	if ms.cursor >= len(ms.v) {
		return ms.cursor, false
	}

	return ms.cursor, true
}

func (ms *RAMEmulator) Jump(to int) {
	ms.cursor = to
}

func (ms *RAMEmulator) All() []int {
	return ms.v
}
