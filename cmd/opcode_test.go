package main

import (
	"testing"

	"opcode"
)

type test struct {
	In  string
	Out string
	Mem string
}

var tests = []test{
	test{
		In:  "1,0,0,0,99",
		Mem: "2,0,0,0,99",
	},
	test{
		In:  "2,3,0,3,99",
		Mem: "2,3,0,6,99",
	},
	test{
		In:  "2,4,4,5,99,0",
		Mem: "2,4,4,5,99,9801",
	},
	test{
		In:  "1,1,1,4,99,5,6,0,99",
		Mem: "30,1,1,4,2,5,6,0,99",
	},
	test{
		In:  "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,5,19,23,2,9,23,27,1,27,5,31,2,31,13,35,1,35,9,39,1,39,10,43,2,43,9,47,1,47,5,51,2,13,51,55,1,9,55,59,1,5,59,63,2,6,63,67,1,5,67,71,1,6,71,75,2,9,75,79,1,79,13,83,1,83,13,87,1,87,5,91,1,6,91,95,2,95,13,99,2,13,99,103,1,5,103,107,1,107,10,111,1,111,13,115,1,10,115,119,1,9,119,123,2,6,123,127,1,5,127,131,2,6,131,135,1,135,2,139,1,139,9,0,99,2,14,0,0",
		Mem: "5110675,12,2,2,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,36,1,5,19,37,2,9,23,111,1,27,5,112,2,31,13,560,1,35,9,563,1,39,10,567,2,43,9,1701,1,47,5,1702,2,13,51,8510,1,9,55,8513,1,5,59,8514,2,6,63,17028,1,5,67,17029,1,6,71,17031,2,9,75,51093,1,79,13,51098,1,83,13,51103,1,87,5,51104,1,6,91,51106,2,95,13,255530,2,13,99,1277650,1,5,103,1277651,1,107,10,1277655,1,111,13,1277660,1,10,115,1277664,1,9,119,1277667,2,6,123,2555334,1,5,127,2555335,2,6,131,5110670,1,135,2,5110672,1,139,9,0,99,2,14,0,0",
	},
	test{
		In:  "104,1125899906842624,99",
		Mem: "104,1125899906842624,99",
	},
	test{
		In:  "1102,34915192,34915192,7,4,7,99,0",
		Mem: "1102,34915192,34915192,7,4,7,99,1219070632396864",
		Out: "1219070632396864",
	},
}

func TestMain(t *testing.T) {
	i := 0

	for _, test := range tests {
		os, err := opcode.BootFromString(test.In, DefaultApps)
		if err != nil {
			t.Error(err)
		}

		err = os.Run()
		if err != nil {
			t.Error(err)
		}

		if test.Mem != os.Dump() {
			t.Errorf("mismatching memory for test %d", i)
			t.Errorf("got:\n%+v", os.Dump())
		}

		if test.Out != "" {
			actual := os.StdOut()
			if actual != test.Out {
				t.Errorf("unexpected stdOut for test %d", i)
				t.Errorf("got:\n%s", actual)
			}
		}

		i++
	}
}
