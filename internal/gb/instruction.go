package gb

// instruction represents a CPU instruction and is used to lookup instruction
// byte length and number of cycles required
type instruction struct {
	name   string // instruction mnemonic
	length uint16 // number of bytes
	cycles int    // number of machine cycles
}

const INSTRUCTION_COUNT = 0x100

var instructions = [INSTRUCTION_COUNT]instruction{}

func init() {
	instructions[0x00] = instruction{"NOP", 1, 1}
	instructions[0x31] = instruction{"LD", 3, 3}
	instructions[0xC3] = instruction{"JP", 3, 4}
	instructions[0xF3] = instruction{"DI", 1, 1}
}
