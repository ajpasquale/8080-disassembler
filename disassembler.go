package disassembler

import (
	"bufio"
	"fmt"
	"os"
)

type operation struct {
	instr string
	size  int
}

func Disassemble(bs []byte) {
	f, err := os.Create("output")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for i := 0; i < len(bs); i++ {
		b := bs[i]
		op := operList[b]
		switch op.size {
		case 0:
		case 1:
			_, err := fmt.Fprintf(w, "%04d %02x %s\n", i, b, op.instr)
			if err != nil {
				panic(err)
			}
		case 2:
			_, err := fmt.Fprintf(w, "%04d %02x %02x %s $%02x\n", i, b, bs[i+1], op.instr, bs[i+1])
			if err != nil {
				panic(err)
			}
			i += op.size - 1
		case 3:
			_, err := fmt.Fprintf(w, "%04d %02x %02x %02x %s $%02x%02x\n", i, b, bs[i+1], bs[i+2], op.instr, bs[i+2], bs[i+1])
			if err != nil {
				panic(err)
			}
			i += op.size - 1
		}
	}
	w.Flush()
}
