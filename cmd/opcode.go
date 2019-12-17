package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"opcode"
	"opcode/applications"
	"opcode/memory"
	"opcode/os"
)

var DefaultApps = []opcode.Application{
	applications.Halt,
	applications.Add,
	applications.Mul,
	applications.Input,
	applications.Output,
	applications.JumpIfTrue,
	applications.JumpIfFalse,
	applications.LessThan,
	applications.Equals,
	applications.AdjustRelativeBase,
}

func main() {
	programFilePath := flag.String("path", "", "The location of the program")

	flag.Parse()

	if programFilePath == nil || *programFilePath == "" {
		fmt.Println("you need to specify a program with -path {path}")
		return
	}

	start, err := ioutil.ReadFile(*programFilePath)
	if err != nil {
		fmt.Printf("cannot read program file: %w", err)
		return
	}

	mem, err := memory.NewRAMStore(start, opcode.IntP(2048))
	if err != nil {
		panic(err)
	}

	os, err := os.Boot(os.OSBootParams{
		Debug:        false,
		Memory:       mem,
		Applications: DefaultApps,
	})
	if err != nil {
		panic(err)
	}

	err = os.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(os.GetStdOut("\n"))
}
