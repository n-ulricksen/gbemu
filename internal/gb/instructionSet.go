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

// dec8 decrements the given 8-big register and sets appropriate flags
func (cpu *CPU) dec8(reg *register8Bit) {
	val := reg.value
	res := reg.value - 1
	reg.value = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, true)
	cpu.setFlag(FLAG_H, halfCarryOccurs(val, 0xFF))
}

// add performs an addition on the value in register A and the given value, and
// stores the result in register A
func (cpu *CPU) add(b byte) {
	a := cpu.AF.getHi()
	res := a + b
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, halfCarryOccurs(a, b))
	cpu.setFlag(FLAG_C, a > res)
}

// sub sutracts the given value from the value in register A, and stores the
// result in register A
func (cpu *CPU) sub(b byte) {
	a := cpu.AF.getHi()
	res := a - b
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, true)
	cpu.setFlag(FLAG_H, halfCarryOccurs(a, -b))
	cpu.setFlag(FLAG_C, a > res)
}

// adc performs an addition with carry on the value in register A and the given
// value, and stores the result in register A
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

// and performs a bitwise and on the value in register A and the given value,
// and stores the result in register A
func (cpu *CPU) and(b byte) {
	a := cpu.AF.getHi()
	res := a & b
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, true)
	cpu.setFlag(FLAG_C, false)
}

// or performs a bitwise or on the value in register A and the given value,
// and stores the result in register A
func (cpu *CPU) or(b byte) {
	a := cpu.AF.getHi()
	res := a | b
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, false)
}

// xor performs a bitwise xor on the value in register A and the given value,
// and stores the result in register A
func (cpu *CPU) xor(b byte) {
	a := cpu.AF.getHi()
	res := a ^ b
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, false)
}

// rlca rotates A left 1 bit with wrapping, leaving the previous MSB in the LSB
// position
func (cpu *CPU) rlca() {
	a := cpu.AF.getHi()
	res := (a << 1) | (a >> 7)
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, false)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (a>>7) > 0)
}

// rla rotates A left 1 bit through the carry flag with wrapping, leaving the
// previous carry bit in the LSB, and the previous MSB in the carry flag
func (cpu *CPU) rla() {
	a := cpu.AF.getHi()
	carry := byte(0)
	if cpu.getFlag(FLAG_C) {
		carry = 1
	}

	res := (a << 1) | carry
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, false)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (a>>7) > 0)
}

// rrca rotates A right 1 bit with wrapping, leaving the previous LSB in the
// MSB position
func (cpu *CPU) rrca() {
	a := cpu.AF.getHi()
	res := ((a & 0x1) << 7) | (a >> 1)
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, false)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (a&0x1) > 0)
}

// rra rotates A right 1 bit through the carry flag with wrapping, leaving the
// previous carry bit in the MSB, and the previous LSB in the carry flag
func (cpu *CPU) rra() {
	a := cpu.AF.getHi()
	carry := byte(0)
	if cpu.getFlag(FLAG_C) {
		carry = 1
	}

	res := (carry << 7) | (a >> 1)
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, false)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (a&0x1) > 0)
}

// rlc rotates the given data left 1 bit with wrapping, leaving the previous
// MSB in the LSB position
func (cpu *CPU) rlc(data *byte) {
	d := *data
	res := (d << 1) | (d >> 7)
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (d>>7) > 0)
}

// rl rotates the given data left 1 bit through the carry flag with wrapping,
// leaving the previous carry bit in the LSB, and the previous MSB in the carry
// flag
func (cpu *CPU) rl(data *byte) {
	d := *data
	carry := byte(0)
	if cpu.getFlag(FLAG_C) {
		carry = 1
	}

	res := (d << 1) | carry
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (d>>7) > 0)
}

// rrc rotates the given data right 1 bit with wrapping, leaving the previous
// LSB in the MSB position
func (cpu *CPU) rrc(data *byte) {
	d := *data
	res := ((d & 0x1) << 7) | (d >> 1)
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (d&0x1) > 0)
}

// rr rotates the given data right 1 bit through the carry flag with wrapping,
// leaving the previous carry bit in the MSB, and the previous LSB in the carry
// flag
func (cpu *CPU) rr(data *byte) {
	d := *data
	carry := byte(0)
	if cpu.getFlag(FLAG_C) {
		carry = 1
	}

	res := (carry << 7) | (d >> 1)
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, (d&0x1) > 0)
}

// sla performs an arithmetic left shift on the given data
func (cpu *CPU) sla(data *byte) {
	res := *data << 1
	carry := *data&(0x1<<7) > 0
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, carry)
}

// sra performs an arithmetic right shift on the given data
func (cpu *CPU) sra(data *byte) {
	msb := *data & (0x1 << 7)
	res := msb | (*data >> 1)
	carry := *data&(0x1) > 0
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, carry)
}

// srl performs a logical right shift on the given data
func (cpu *CPU) srl(data *byte) {
	res := *data >> 1
	carry := *data&(0x1) > 0
	*data = res

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, false)
	cpu.setFlag(FLAG_C, carry)
}

// swap exchanges the low and high nibble of the given byte
func (cpu *CPU) swap(data *byte) {
	hi := *data & 0xF0
	lo := *data & 0x0F
	*data = (lo << 4) | (hi >> 4)
}

// bit tests bit 'b', setting the zero flag if the bit is 0
func (cpu *CPU) bit(b int, data *byte) {
	z := (*data & (1 << b)) == 0
	cpu.setFlag(FLAG_Z, z)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, true)
}

// res resets (clears) a bit 'b' in the given data
func (cpu *CPU) res(b int, data *byte) {
	*data &^= (1 << b)
}

// set sets a bit 'b' to 1 in the given data
func (cpu *CPU) set(b int, data *byte) {
	*data &= (1 << b)
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

// retIf returns from a function if the given condtion is true
func (cpu *CPU) retIf(condition bool) {
	if condition {
		cpu.ret()
		cpu.cycles += 3
	}
}

// di disables interrupt handling
func (cpu *CPU) di() {
	cpu.IME = false
}

// nop nops
func (cpu *CPU) nop() {}
