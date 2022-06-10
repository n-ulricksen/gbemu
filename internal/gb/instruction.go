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
	instructions[0x18] = instruction{"JR", 2, 3, cpu.op18}
	instructions[0x21] = instruction{"LD", 3, 3, cpu.op21}
	instructions[0x31] = instruction{"LD", 3, 3, cpu.op31}
	instructions[0x3E] = instruction{"LD", 2, 2, cpu.op3E}
	instructions[0x7C] = instruction{"LD", 1, 1, cpu.op7C}
	instructions[0x7D] = instruction{"LD", 1, 1, cpu.op7D}
	instructions[0xC3] = instruction{"JP", 3, 4, cpu.opC3}
	instructions[0xC9] = instruction{"RET", 1, 4, cpu.opC9}
	instructions[0xCD] = instruction{"CALL", 3, 6, cpu.opCD}
	instructions[0xE0] = instruction{"LDH", 2, 3, cpu.opE0}
	instructions[0xEA] = instruction{"LD", 3, 4, cpu.opEA}
	instructions[0xF3] = instruction{"DI", 1, 1, cpu.opF3}
}

// NOP
func (cpu *CPU) op00() {}

// JR e
func (cpu *CPU) op18() {
	offset := cpu.read(cpu.PC + 1)
	cpu.PC += uint16(offset)

	cpu.PC -= instructions[0x18].length
}

// LD HL,nn
func (cpu *CPU) op21() {
	data := cpu.readWord(cpu.PC + 1)
	cpu.HL.set(data)
}

// LD SP,nn
func (cpu *CPU) op31() {
	cpu.SP = cpu.readWord(cpu.PC + 1)
}

// LD A,n
func (cpu *CPU) op3E() {
	data := cpu.read(cpu.PC + 1)
	cpu.AF.setHi(data)
}

// LD A,H
func (cpu *CPU) op7C() {
	data := cpu.HL.hiVal()
	cpu.AF.setHi(data)
}

// LD A,L
func (cpu *CPU) op7D() {
	data := cpu.HL.loVal()
	cpu.AF.setHi(data)
}

// JP nn
func (cpu *CPU) opC3() {
	addr := cpu.readWord(cpu.PC + 1)
	cpu.jp(addr)

	cpu.PC -= instructions[0xC3].length
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

// LDH (n),A
func (cpu *CPU) opE0() {
	lo := cpu.read(cpu.PC + 1)
	addr := u16(lo, 0xFF)
	cpu.ld8(addr, cpu.AF.hiVal())
}

// LD (nn),A
func (cpu *CPU) opEA() {
	addr := cpu.readWord(cpu.PC + 1)
	cpu.ld8(addr, cpu.AF.hiVal())
}

// DI
func (cpu *CPU) opF3() {
	cpu.di()
}
