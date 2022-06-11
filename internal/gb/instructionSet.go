package gb

// ld8 loads a byte into memory at the given address
func (cpu *CPU) ld8(addr uint16, data byte) {
	cpu.write(addr, data)
}

// jp jumps unconditionally to the given 16-bit address
func (cpu *CPU) jp(addr uint16) {
	cpu.PC = addr
}

// jrIf performs a conditional relative jump based on the given condition
func (cpu *CPU) jrIf(offset byte, condition bool) {
	if condition {
		cpu.PC += uint16(offset)
		cpu.cycles++
	}
}

// inc8 increments the given 8-big register and sets appropriate flags
func (cpu *CPU) inc8(reg *register8Bit) {
	val := reg.value
	res := reg.value + 1
	reg.value = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, halfCarryOccurs(val, 1))
}

// adc performs an adition with carry on the value in register A and the given
// value. The result is stored in register A.
func (cpu *CPU) adc(add byte) {
	carry := byte(0)
	if cpu.getFlag(FLAG_C) {
		carry = 1
	}

	a := cpu.AF.getHi()
	res := a + add + carry

	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, halfCarryOccurs(a, add+carry))
	cpu.setFlag(FLAG_C, a > res)
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
