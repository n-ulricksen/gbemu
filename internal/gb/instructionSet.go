package gb

// ld8 loads a byte into memory at the given address
func (cpu *CPU) ld8(addr uint16, data byte) {
	cpu.write(addr, data)
}

// jp jumps unconditionally to the given 16-bit address
func (cpu *CPU) jp(addr uint16) {
	cpu.PC = addr
}

// push pushes a word of data to the stack
func (cpu *CPU) push(data uint16) {
	hi := byte(data >> 8)
	lo := byte(data)
	cpu.stackPush(hi)
	cpu.stackPush(lo)
}

// pop returns the top word on the stack
func (cpu *CPU) pop() uint16 {
	lo := cpu.stackPop()
	hi := cpu.stackPop()
	return u16(lo, hi)
}

// call performs an unconditional function all to the given address
func (cpu *CPU) call(addr uint16) {
	cpu.push(cpu.PC + 1)
	cpu.PC = addr
}

// ret unconditionally returns from a function
func (cpu *CPU) ret() {
	cpu.PC = cpu.pop()
}

// di disables interrupt handling
func (cpu *CPU) di() {
	cpu.IME = false
}

// nop nops
func (cpu *CPU) nop() {}
