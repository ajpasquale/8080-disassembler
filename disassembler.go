package disassembler

import (
	"fmt"
	"io"
	"os"
)

type operation struct {
	instr string
	size  int
}

func Disassemble(file string) {
	bs, err := os.ReadFile(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
		return
	}

	for i := 0; i < len(bs); i++ {
		b := bs[i]
		op := operList[b]
		switch op.size {
		case 0:
		case 1:
			s := fmt.Sprintf("%04d %02x %s\n", i, b, op.instr)
			io.WriteString(os.Stdout, s)
		case 2:
			s := fmt.Sprintf("%04d %02x %02x %s $%02x\n", i, b, bs[i+1], op.instr, bs[i+1])
			io.WriteString(os.Stdout, s)
			i += op.size - 1
		case 3:
			s := fmt.Sprintf("%04d %02x %02x %02x %s $%02x%02x\n", i, b, bs[i+1], bs[i+2], op.instr, bs[i+2], bs[i+1])
			io.WriteString(os.Stdout, s)
			i += op.size - 1
		}
	}
}
