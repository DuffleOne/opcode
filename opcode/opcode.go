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

	var err error
	var de int
	var c int
	var b int
	var a int

	strVal = fmt.Sprintf("%05s", strVal)

	if len(strVal) > 5 {
		return nil, fmt.Errorf("opcode too long")
	}

	de, err = strconv.Atoi(strVal[len(strVal)-2:])
	c, err = getAt(strVal, 3, 2)
	b, err = getAt(strVal, 4, 3)
	a, err = getAt(strVal, 5, 4)

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
