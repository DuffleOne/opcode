package opcode

import "testing"

type opcodetest struct {
	In  int
	Out OPCode
}

var tests = []opcodetest{
	opcodetest{
		In: 2,
		Out: OPCode{
			Code:       2,
			Param1Mode: 0,
			Param2Mode: 0,
			Param3Mode: 0,
		},
	},
	opcodetest{
		In: 02,
		Out: OPCode{
			Code:       2,
			Param1Mode: 0,
			Param2Mode: 0,
			Param3Mode: 0,
		},
	},
	opcodetest{
		In: 102,
		Out: OPCode{
			Code:       2,
			Param1Mode: 1,
			Param2Mode: 0,
			Param3Mode: 0,
		},
	},
	opcodetest{
		In: 1102,
		Out: OPCode{
			Code:       2,
			Param1Mode: 1,
			Param2Mode: 1,
			Param3Mode: 0,
		},
	},
	opcodetest{
		In: 11102,
		Out: OPCode{
			Code:       2,
			Param1Mode: 1,
			Param2Mode: 1,
			Param3Mode: 1,
		},
	},
	opcodetest{
		In: 10102,
		Out: OPCode{
			Code:       2,
			Param1Mode: 1,
			Param2Mode: 0,
			Param3Mode: 1,
		},
	},
}

func TestOPCodeCreation(t *testing.T) {
	i := 1

	for _, test := range tests {
		c, err := BuildOPCode(test.In)
		if err != nil {
			t.Error(err)
		}

		if c.Code != test.Out.Code {
			t.Errorf("got a bad code for test %d, expected %d, got %d", i, test.Out.Code, c.Code)
		}

		if c.Param1Mode != test.Out.Param1Mode {
			t.Errorf("got a bad Param1Mode for test %d, expected %d, got %d", i, test.Out.Param1Mode, c.Param1Mode)
		}

		if c.Param2Mode != test.Out.Param2Mode {
			t.Errorf("got a bad Param2Mode for test %d, expected %d, got %d", i, test.Out.Param2Mode, c.Param2Mode)
		}

		if c.Param3Mode != test.Out.Param3Mode {
			t.Errorf("got a bad Param3Mode for test %d, expected %d, got %d", i, test.Out.Param3Mode, c.Param3Mode)
		}

		i++
	}
}
