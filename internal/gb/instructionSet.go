package gb

// ld8 loads a byte into memory at the given address
func (cpu *CPU) ld8(addr uint16, data byte) {
	cpu.write(addr, data)
}

// jp jumps unconditionally to the given 16-bit address
func (cpu *CPU) jp(addr uint16) {
	cpu.PC = addr
}

// call performs an unconditional function all to the given address
func (cpu *CPU) call(addr uint16) {
	hi := byte(cpu.PC >> 8)
	lo := byte(cpu.PC + 1)
	cpu.stackPush(hi)
	cpu.stackPush(lo)

	cpu.PC = addr
}

// ret unconditionally returns from a function
func (cpu *CPU) ret() {
	lo := cpu.stackPop()
	hi := cpu.stackPop()
	cpu.PC = u16(lo, hi)
}

// di disables interrupt handling
func (cpu *CPU) di() {
	cpu.IME = false
}

// nop nops
func (cpu *CPU) nop() {}
