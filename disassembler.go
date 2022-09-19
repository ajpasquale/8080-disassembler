package disassembler

import (
	"fmt"
	"io"
	"os"
	"strings"
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
			ss := strings.Split(op.instr, " ")

			if len(ss) > 1 {
				s := fmt.Sprintf("%04x %02x\t\t %s\t %s\n", i, b, ss[0], ss[1])
				io.WriteString(os.Stdout, s)
			} else {
				s := fmt.Sprintf("%04x %02x\t\t %s\n", i, b, op.instr)
				io.WriteString(os.Stdout, s)
			}

		case 2:
			ss := strings.Split(op.instr, " ")
			if len(ss) > 1 {
				s := fmt.Sprintf("%04x %02x %02x\t %s\t %s#$%02x\n", i, b, bs[i+1], ss[0], ss[1], bs[i+1])
				io.WriteString(os.Stdout, s)
			} else {
				s := fmt.Sprintf("%04x %02x %02x\t %s\t $%02x\n", i, b, bs[i+1], op.instr, bs[i+1])
				io.WriteString(os.Stdout, s)
			}

			i += op.size - 1
		case 3:
			ss := strings.Split(op.instr, " ")
			if len(ss) > 1 {
				s := fmt.Sprintf("%04x %02x %02x %02x\t %s\t %s#$%02x%02x\n", i, b, bs[i+1], bs[i+2], ss[0], ss[1], bs[i+2], bs[i+1])
				io.WriteString(os.Stdout, s)
			} else {
				s := fmt.Sprintf("%04x %02x %02x %02x\t %s\t $%02x%02x\n", i, b, bs[i+1], bs[i+2], op.instr, bs[i+2], bs[i+1])
				io.WriteString(os.Stdout, s)
			}

			i += op.size - 1
		}
	}
}
