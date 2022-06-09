package gb

// jp jumps unconditionally to the given 16-bit address
// 	JP nn
func (cpu *CPU) jp(addr uint16) {
	cpu.PC = addr
}

// di disables interrupt handling
//	DI
func (cpu *CPU) di() {
	cpu.IME = false
}

// nop nops
// 	NOP
func (cpu *CPU) nop() {}
