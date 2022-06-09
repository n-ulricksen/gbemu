package gb

// ld8 loads a byte into memory at the given address
func (cpu *CPU) ld8(addr uint16, data byte) {
	cpu.write(addr, data)
}

// jp jumps unconditionally to the given 16-bit address
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
