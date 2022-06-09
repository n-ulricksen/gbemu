package gb

// jp jumps unconditionally to the given 16-bit address
// 	JP nn
func (cpu *CPU) jp(addr uint16) {
	cpu.PC = addr
}

// nop nops
func (cpu *CPU) nop() {}
