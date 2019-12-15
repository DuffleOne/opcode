package opcode

import (
	"fmt"
	"strconv"
)

type OPCode struct {
	Code       int
	Param1Mode int
	Param2Mode int
	Param3Mode int
}

func BuildOPCode(in int) (*OPCode, error) {
	strVal := strconv.Itoa(in)

	strVal = fmt.Sprintf("%05s", strVal)

	if len(strVal) != 5 {
		return nil, fmt.Errorf("malformed opcode")
	}

	deStr := strVal[3:]
	cStr := strVal[2:3]
	bStr := strVal[1:2]
	aStr := strVal[:1]

	de, err := strconv.Atoi(deStr)
	if err != nil {
		return nil, err
	}

	c, err := strconv.Atoi(cStr)
	if err != nil {
		return nil, err
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		return nil, err
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		return nil, err
	}

	code := &OPCode{
		Code:       de,
		Param1Mode: c,
		Param2Mode: b,
		Param3Mode: a,
	}

	return code, err
}

func (oc *OPCode) String() string {
	return fmt.Sprintf("%d%d%d%02d", oc.Param3Mode, oc.Param2Mode, oc.Param1Mode, oc.Code)
}

func getAt(in string, higher, lower int) (int, error) {
	out, err := strconv.Atoi(in[len(in)-higher : len(in)-lower])
	if err != nil {
		return 0, err
	}

	return out, nil
}
