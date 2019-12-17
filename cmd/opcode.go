package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	goos "os"

	"opcode"
	"opcode/applications"
	"opcode/memory"
	"opcode/os"
)

var defaultPath = "./program.txt"

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
		if !fileExists(defaultPath) {
			fmt.Println("cannot open program, use \"-path {path}\" to specify")
			fmt.Printf("or use the default location \"%s\"\n", defaultPath)
			return
		}

		programFilePath = opcode.StringP(defaultPath)
	}

	start, err := ioutil.ReadFile(*programFilePath)
	if err != nil {
		fmt.Printf("cannot read program file: %w", err)
		return
	}

	if string(start) == "" {
		fmt.Println("program file is empty")
		return
	}

	mem, err := memory.NewRAMStore(string(start), opcode.IntP(2048))
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

func fileExists(filename string) bool {
	info, err := goos.Stat(filename)
	if goos.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
