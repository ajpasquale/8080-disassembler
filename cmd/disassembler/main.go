package main

import (
	"os"

	disassembler "github.com/ajpasquale/8080-disassembler"
)

func main() {
	arg := os.Args[1]
	disassembler.Disassemble(arg)
}
