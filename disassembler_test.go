package disassembler

import (
	"os"
	"testing"
)

func TestDisassemble(t *testing.T) {
	b, _ := os.ReadFile("rom/invaders.h")
	Disassemble(b)
}
