package gb

// instruction represents a CPU instruction and is used to lookup instruction
// byte length, number of cycles required, and the method used to execute the
// instruction
type instruction struct {
	name   string // instruction mnemonic
	length uint16 // number of bytes
	cycles int    // number of machine cycles
	exec   func() // method to execute instruction
}

// 0x00 - 0xFF
const INSTRUCTION_COUNT = 0x100

// instructions is the instruction lookup array, used by the CPU during the
// decode stage
var instructions = [INSTRUCTION_COUNT]instruction{}

// setupInstuctionLookup defines all legal CPU instructions for the instruction
// lookup array
func (cpu *CPU) setupInstructionLookup() {
	instructions[0x00] = instruction{"NOP", 1, 1, cpu.op00}
	instructions[0x01] = instruction{"LD", 3, 3, cpu.op01}
	instructions[0x02] = instruction{"LD", 1, 2, cpu.op02}
	instructions[0x03] = instruction{"INC", 1, 2, cpu.op03}
	instructions[0x05] = instruction{"DEC", 1, 1, cpu.op05}
	instructions[0x06] = instruction{"LD", 2, 2, cpu.op06}
	instructions[0x07] = instruction{"RLCA", 1, 1, cpu.op07}
	instructions[0x18] = instruction{"JR", 2, 3, cpu.op18}
	instructions[0x1C] = instruction{"INC", 1, 1, cpu.op1C}
	instructions[0x20] = instruction{"JR", 2, 2, cpu.op20}
	instructions[0x21] = instruction{"LD", 3, 3, cpu.op21}
	instructions[0x23] = instruction{"INC", 1, 2, cpu.op23}
	instructions[0x28] = instruction{"JR", 2, 2, cpu.op28}
	instructions[0x2A] = instruction{"LD", 1, 2, cpu.op2A}
	instructions[0x30] = instruction{"JR", 2, 2, cpu.op30}
	instructions[0x31] = instruction{"LD", 3, 3, cpu.op31}
	instructions[0x32] = instruction{"LD", 1, 2, cpu.op32}
	instructions[0x36] = instruction{"LD", 2, 3, cpu.op36}
	instructions[0x38] = instruction{"JR", 2, 2, cpu.op38}
	instructions[0x3C] = instruction{"INC", 1, 1, cpu.op3C}
	instructions[0x3E] = instruction{"LD", 2, 2, cpu.op3E}
	instructions[0x5D] = instruction{"LD", 1, 1, cpu.op5D}
	instructions[0x60] = instruction{"LD", 1, 1, cpu.op60}
	instructions[0x66] = instruction{"LD", 1, 2, cpu.op66}
	instructions[0x6C] = instruction{"LD", 1, 1, cpu.op6C}
	instructions[0x6E] = instruction{"LD", 1, 2, cpu.op6E}
	instructions[0x76] = instruction{"HALT", 1, 1, cpu.op76}
	instructions[0x78] = instruction{"LD", 1, 1, cpu.op78}
	instructions[0x7B] = instruction{"LD", 1, 1, cpu.op7B}
	instructions[0x7C] = instruction{"LD", 1, 1, cpu.op7C}
	instructions[0x7D] = instruction{"LD", 1, 1, cpu.op7D}
	instructions[0x7E] = instruction{"LD", 1, 2, cpu.op7E}
	instructions[0x8E] = instruction{"ADC", 1, 2, cpu.op8E}
	instructions[0xA3] = instruction{"AND", 1, 1, cpu.opA3}
	instructions[0xB1] = instruction{"OR", 1, 1, cpu.opB1}
	instructions[0xB7] = instruction{"OR", 1, 1, cpu.opB7}
	instructions[0xC3] = instruction{"JP", 3, 4, cpu.opC3}
	instructions[0xC5] = instruction{"PUSH", 1, 4, cpu.opC5}
	instructions[0xC6] = instruction{"ADD", 2, 2, cpu.opC6}
	instructions[0xC9] = instruction{"RET", 1, 4, cpu.opC9}
	instructions[0xCD] = instruction{"CALL", 3, 6, cpu.opCD}
	instructions[0xD6] = instruction{"SUB", 2, 2, cpu.opD6}
	instructions[0xE0] = instruction{"LDH", 2, 3, cpu.opE0}
	instructions[0xE1] = instruction{"POP", 1, 3, cpu.opE1}
	instructions[0xE5] = instruction{"PUSH", 1, 4, cpu.opE5}
	instructions[0xEA] = instruction{"LD", 3, 4, cpu.opEA}
	instructions[0xEE] = instruction{"XOR", 2, 2, cpu.opEE}
	instructions[0xF0] = instruction{"LDH", 2, 3, cpu.opF0}
	instructions[0xF1] = instruction{"POP", 1, 3, cpu.opF1}
	instructions[0xF3] = instruction{"DI", 1, 1, cpu.opF3}
	instructions[0xF5] = instruction{"PUSH", 1, 4, cpu.opF5}
	instructions[0xFE] = instruction{"CP", 2, 2, cpu.opFE}
	instructions[0xFF] = instruction{"RST", 1, 4, cpu.opFF}
}

// NOP
func (cpu *CPU) op00() {}

// LD BC,nn
func (cpu *CPU) op01() {
	data := cpu.readWord(cpu.PC + 1)
	cpu.BC.set(data)
}

// LD (BC),A
func (cpu *CPU) op02() {
	data := cpu.AF.getHi()
	addr := cpu.BC.get()
	cpu.ld8(addr, data)
}

// INC BC
func (cpu *CPU) op03() {
	cpu.BC.inc()
}

// DEC B
func (cpu *CPU) op05() {
	cpu.dec8(&cpu.BC.hiReg)
}

// LD B,n
func (cpu *CPU) op06() {
	data := cpu.read(cpu.PC + 1)
	cpu.BC.setHi(data)
}

// RLCA
func (cpu *CPU) op07() {
	cpu.rlca()
}

// JR e
func (cpu *CPU) op18() {
	offset := cpu.read(cpu.PC + 1)
	cpu.PC += uint16(offset)
}

// INC E
func (cpu *CPU) op1C() {
	cpu.inc8(&cpu.DE.loReg)
}

// LD HL,nn
func (cpu *CPU) op21() {
	data := cpu.readWord(cpu.PC + 1)
	cpu.HL.set(data)
}

// INC HL
func (cpu *CPU) op23() {
	cpu.HL.inc()
}

// JR NZ,e
func (cpu *CPU) op20() {
	offset := cpu.read(cpu.PC + 1)
	cond := !cpu.getFlag(FLAG_Z)
	cpu.jrIf(offset, cond)
}

// JR Z,e
func (cpu *CPU) op28() {
	offset := cpu.read(cpu.PC + 1)
	cond := cpu.getFlag(FLAG_Z)
	cpu.jrIf(offset, cond)
}

// LD A,(HL+)
func (cpu *CPU) op2A() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.AF.setHi(data)

	cpu.HL.inc()
}

// JR NC,e
func (cpu *CPU) op30() {
	offset := cpu.read(cpu.PC + 1)
	cond := !cpu.getFlag(FLAG_C)
	cpu.jrIf(offset, cond)
}

// LD SP,nn
func (cpu *CPU) op31() {
	cpu.SP = cpu.readWord(cpu.PC + 1)
}

// LD (HL-),A
func (cpu *CPU) op32() {
	data := cpu.AF.getHi()
	addr := cpu.HL.get()
	cpu.ld8(addr, data)

	cpu.HL.dec()
}

// LD (HL),n
func (cpu *CPU) op36() {
	data := cpu.read(cpu.PC + 1)
	addr := cpu.HL.get()
	cpu.ld8(addr, data)
}

// JR C,e
func (cpu *CPU) op38() {
	offset := cpu.read(cpu.PC + 1)
	cond := cpu.getFlag(FLAG_C)
	cpu.jrIf(offset, cond)
}

// INC A
func (cpu *CPU) op3C() {
	cpu.inc8(&cpu.AF.hiReg)
}

// LD A,n
func (cpu *CPU) op3E() {
	data := cpu.read(cpu.PC + 1)
	cpu.AF.setHi(data)
}

// LD E,L
func (cpu *CPU) op5D() {
	data := cpu.HL.getLo()
	cpu.DE.setLo(data)
}

// LD H,B
func (cpu *CPU) op60() {
	data := cpu.BC.getHi()
	cpu.HL.setHi(data)
}

// LD H,(HL)
func (cpu *CPU) op66() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.HL.setHi(data)
}

// LD L,H
func (cpu *CPU) op6C() {
	data := cpu.HL.getHi()
	cpu.HL.setLo(data)
}

// LD L,(HL)
func (cpu *CPU) op6E() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.HL.setLo(data)
}

// HALT
func (cpu *CPU) op76() {
	cpu.halted = true
}

// LD A,B
func (cpu *CPU) op78() {
	data := cpu.BC.getHi()
	cpu.AF.setHi(data)
}

// LD A,E
func (cpu *CPU) op7B() {
	data := cpu.DE.getLo()
	cpu.AF.setHi(data)
}

// LD A,H
func (cpu *CPU) op7C() {
	data := cpu.HL.getHi()
	cpu.AF.setHi(data)
}

// LD A,L
func (cpu *CPU) op7D() {
	data := cpu.HL.getLo()
	cpu.AF.setHi(data)
}

// LD A,(HL)
func (cpu *CPU) op7E() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.AF.setHi(data)
}

// ADC A,(HL)
func (cpu *CPU) op8E() {
	add := cpu.read(cpu.HL.get())
	cpu.adc(add)
}

// AND E
func (cpu *CPU) opA3() {
	res := cpu.AF.getHi() & cpu.DE.getLo()
	cpu.AF.setHi(res)

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, false)
	cpu.setFlag(FLAG_H, true)
	cpu.setFlag(FLAG_C, false)
}

// OR C
func (cpu *CPU) opB1() {
	c := cpu.BC.getLo()
	cpu.or(c)
}

// OR A
func (cpu *CPU) opB7() {
	a := cpu.AF.getHi()
	cpu.or(a)
}

// JP nn
func (cpu *CPU) opC3() {
	addr := cpu.readWord(cpu.PC + 1)
	cpu.jp(addr)

	cpu.PC -= instructions[0xC3].length
}

// PUSH BC
func (cpu *CPU) opC5() {
	data := cpu.BC.get()
	cpu.push(data)
}

// ADD A,n
func (cpu *CPU) opC6() {
	data := cpu.read(cpu.PC + 1)
	cpu.add(data)
}

// RET
func (cpu *CPU) opC9() {
	cpu.ret()

	cpu.PC -= instructions[0xC9].length
}

// CALL nn
func (cpu *CPU) opCD() {
	addr := cpu.readWord(cpu.PC + 1)
	cpu.call(addr)

	cpu.PC -= instructions[0xCD].length
}

// SUB n
func (cpu *CPU) opD6() {
	data := cpu.read(cpu.PC + 1)
	cpu.sub(data)
}

// LDH (n),A
func (cpu *CPU) opE0() {
	lo := cpu.read(cpu.PC + 1)
	addr := u16(lo, 0xFF)
	data := cpu.AF.getHi()
	cpu.ld8(addr, data)
}

// POP HL
func (cpu *CPU) opE1() {
	data := cpu.pop()
	cpu.HL.set(data)
}

// PUSH HL
func (cpu *CPU) opE5() {
	data := cpu.HL.get()
	cpu.push(data)
}

// LD (nn),A
func (cpu *CPU) opEA() {
	addr := cpu.readWord(cpu.PC + 1)
	data := cpu.AF.getHi()
	cpu.ld8(addr, data)
}

// XOR n
func (cpu *CPU) opEE() {
	n := cpu.read(cpu.PC + 1)
	cpu.xor(n)
}

// LDH A,(n)
func (cpu *CPU) opF0() {
	lo := cpu.read(cpu.PC + 1)
	addr := u16(lo, 0xFF)
	data := cpu.read(addr)
	cpu.AF.setHi(data)
}

// POP AF
func (cpu *CPU) opF1() {
	data := cpu.pop()
	cpu.AF.set(data)
}

// DI
func (cpu *CPU) opF3() {
	cpu.di()
}

// PUSH AF
func (cpu *CPU) opF5() {
	data := cpu.AF.get()
	cpu.push(data)
}

// CP n
func (cpu *CPU) opFE() {
	n := cpu.read(cpu.PC + 1)
	add1 := cpu.AF.getHi()
	add2 := byte(int(n) * -1)
	res := add1 + add2

	cpu.setFlag(FLAG_Z, res == 0)
	cpu.setFlag(FLAG_N, true)
	cpu.setFlag(FLAG_H, halfCarryOccurs(add1, add2))
	cpu.setFlag(FLAG_C, add1 > res)
}

// RST n
func (cpu *CPU) opFF() {
	n := 0x38
	addr := uint16(n)
	cpu.call(addr)

	cpu.PC -= instructions[0xFF].length
}
