package gb

// inst represents a CPU inst and is used to lookup inst
// byte length, number of cycles required, and the method used to execute the
// inst
type inst struct {
	name   string // instruction mnemonic
	length uint16 // number of bytes
	cycles int    // number of machine cycles
	exec   func() // method to execute instruction
}

// 0x00 - 0xFF
const INSTRUCTION_COUNT = 0x100

// instruction lookup arrays, used by the CPU during the decode stage
var (
	instructions = [INSTRUCTION_COUNT]inst{}
)

// setupInstuctionLookup defines all legal CPU instructions for the instruction
// lookup array
func (cpu *CPU) setupInstructionLookup() {
	instructions[0x00] = inst{"NOP", 1, 1, cpu.op00}
	instructions[0x01] = inst{"LD", 3, 3, cpu.op01}
	instructions[0x02] = inst{"LD", 1, 2, cpu.op02}
	instructions[0x03] = inst{"INC", 1, 2, cpu.op03}
	instructions[0x05] = inst{"DEC", 1, 1, cpu.op05}
	instructions[0x06] = inst{"LD", 2, 2, cpu.op06}
	instructions[0x07] = inst{"RLCA", 1, 1, cpu.op07}
	instructions[0x0A] = inst{"LD", 1, 2, cpu.op0A}
	instructions[0x0E] = inst{"LD", 2, 2, cpu.op0E}
	instructions[0x10] = inst{"STOP", 2, 1, cpu.op10}
	instructions[0x11] = inst{"LD", 3, 3, cpu.op11}
	instructions[0x12] = inst{"LD", 1, 2, cpu.op12}
	instructions[0x13] = inst{"INC", 1, 2, cpu.op13}
	instructions[0x18] = inst{"JR", 2, 3, cpu.op18}
	instructions[0x1C] = inst{"INC", 1, 1, cpu.op1C}
	instructions[0x1F] = inst{"RRA", 1, 1, cpu.op1F}
	instructions[0x20] = inst{"JR", 2, 2, cpu.op20}
	instructions[0x21] = inst{"LD", 3, 3, cpu.op21}
	instructions[0x22] = inst{"LD", 1, 2, cpu.op22}
	instructions[0x23] = inst{"INC", 1, 2, cpu.op23}
	instructions[0x25] = inst{"DEC", 1, 1, cpu.op25}
	instructions[0x26] = inst{"LD", 2, 2, cpu.op26}
	instructions[0x28] = inst{"JR", 2, 2, cpu.op28}
	instructions[0x2A] = inst{"LD", 1, 2, cpu.op2A}
	instructions[0x2D] = inst{"DEC", 1, 1, cpu.op2D}
	instructions[0x2F] = inst{"CPL", 1, 1, cpu.op2F}
	instructions[0x30] = inst{"JR", 2, 2, cpu.op30}
	instructions[0x31] = inst{"LD", 3, 3, cpu.op31}
	instructions[0x32] = inst{"LD", 1, 2, cpu.op32}
	instructions[0x36] = inst{"LD", 2, 3, cpu.op36}
	instructions[0x38] = inst{"JR", 2, 2, cpu.op38}
	instructions[0x3C] = inst{"INC", 1, 1, cpu.op3C}
	instructions[0x3D] = inst{"DEC", 1, 1, cpu.op3D}
	instructions[0x3E] = inst{"LD", 2, 2, cpu.op3E}
	instructions[0x46] = inst{"LD", 1, 2, cpu.op46}
	instructions[0x47] = inst{"LD", 1, 1, cpu.op47}
	instructions[0x4E] = inst{"LD", 1, 2, cpu.op4E}
	instructions[0x4F] = inst{"LD", 1, 1, cpu.op4F}
	instructions[0x56] = inst{"LD", 1, 2, cpu.op56}
	instructions[0x57] = inst{"LD", 1, 1, cpu.op57}
	instructions[0x5D] = inst{"LD", 1, 1, cpu.op5D}
	instructions[0x5F] = inst{"LD", 1, 1, cpu.op5F}
	instructions[0x60] = inst{"LD", 1, 1, cpu.op60}
	instructions[0x62] = inst{"LD", 1, 1, cpu.op62}
	instructions[0x63] = inst{"LD", 1, 1, cpu.op63}
	instructions[0x66] = inst{"LD", 1, 2, cpu.op66}
	instructions[0x67] = inst{"LD", 1, 1, cpu.op67}
	instructions[0x69] = inst{"LD", 1, 1, cpu.op69}
	instructions[0x6B] = inst{"LD", 1, 1, cpu.op6B}
	instructions[0x6C] = inst{"LD", 1, 1, cpu.op6C}
	instructions[0x6E] = inst{"LD", 1, 2, cpu.op6E}
	instructions[0x6F] = inst{"LD", 1, 1, cpu.op6F}
	instructions[0x70] = inst{"LD", 1, 2, cpu.op70}
	instructions[0x72] = inst{"LD", 1, 2, cpu.op72}
	instructions[0x73] = inst{"LD", 1, 2, cpu.op73}
	instructions[0x74] = inst{"LD", 1, 2, cpu.op74}
	instructions[0x75] = inst{"LD", 1, 2, cpu.op75}
	instructions[0x76] = inst{"HALT", 1, 1, cpu.op76}
	instructions[0x78] = inst{"LD", 1, 1, cpu.op78}
	instructions[0x79] = inst{"LD", 1, 1, cpu.op79}
	instructions[0x7A] = inst{"LD", 1, 1, cpu.op7A}
	instructions[0x7B] = inst{"LD", 1, 1, cpu.op7B}
	instructions[0x7C] = inst{"LD", 1, 1, cpu.op7C}
	instructions[0x7D] = inst{"LD", 1, 1, cpu.op7D}
	instructions[0x7E] = inst{"LD", 1, 2, cpu.op7E}
	instructions[0x80] = inst{"ADD", 1, 1, cpu.op80}
	instructions[0x8E] = inst{"ADC", 1, 2, cpu.op8E}
	instructions[0x93] = inst{"SUB", 1, 1, cpu.op93}
	instructions[0xA3] = inst{"AND", 1, 1, cpu.opA3}
	instructions[0xAE] = inst{"XOR", 1, 2, cpu.opAE}
	instructions[0xB1] = inst{"OR", 1, 1, cpu.opB1}
	instructions[0xB7] = inst{"OR", 1, 1, cpu.opB7}
	instructions[0xC0] = inst{"RET", 1, 2, cpu.opC0}
	instructions[0xC3] = inst{"JP", 3, 4, cpu.opC3}
	instructions[0xC5] = inst{"PUSH", 1, 4, cpu.opC5}
	instructions[0xC6] = inst{"ADD", 2, 2, cpu.opC6}
	instructions[0xC9] = inst{"RET", 1, 4, cpu.opC9}
	instructions[0xCB] = inst{"", 2, 2, cpu.opCB} // prefix 0xCB
	instructions[0xCD] = inst{"CALL", 3, 6, cpu.opCD}
	instructions[0xCE] = inst{"ADC", 2, 2, cpu.opCE}
	instructions[0xD0] = inst{"RET", 1, 2, cpu.opD0}
	instructions[0xD1] = inst{"POP", 1, 3, cpu.opD1}
	instructions[0xD5] = inst{"PUSH", 1, 4, cpu.opD5}
	instructions[0xD6] = inst{"SUB", 2, 2, cpu.opD6}
	instructions[0xE0] = inst{"LDH", 2, 3, cpu.opE0}
	instructions[0xE1] = inst{"POP", 1, 3, cpu.opE1}
	instructions[0xE5] = inst{"PUSH", 1, 4, cpu.opE5}
	instructions[0xEA] = inst{"LD", 3, 4, cpu.opEA}
	instructions[0xEE] = inst{"XOR", 2, 2, cpu.opEE}
	instructions[0xF0] = inst{"LDH", 2, 3, cpu.opF0}
	instructions[0xF1] = inst{"POP", 1, 3, cpu.opF1}
	instructions[0xF3] = inst{"DI", 1, 1, cpu.opF3}
	instructions[0xF5] = inst{"PUSH", 1, 4, cpu.opF5}
	instructions[0xF9] = inst{"LD", 1, 2, cpu.opF9}
	instructions[0xFA] = inst{"LD", 3, 4, cpu.opFA}
	instructions[0xFE] = inst{"CP", 2, 2, cpu.opFE}
	instructions[0xFF] = inst{"RST", 1, 4, cpu.opFF}
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

// LD A,(BC)
func (cpu *CPU) op0A() {
	addr := cpu.BC.get()
	data := cpu.read(addr)
	cpu.AF.setHi(data)
}

// LD C,n
func (cpu *CPU) op0E() {
	data := cpu.read(cpu.PC + 1)
	cpu.BC.setLo(data)
}

// STOP n
// According to https://gbdev.io/pandocs/CPU_Instruction_Set.html, the next
// byte of data after a 0x10 should always be a 0x00. This word (0x1000)
// signals the CPU to go into "low power standby mode."
func (cpu *CPU) op10() {
	_ = cpu.read(cpu.PC + 1)
	cpu.stopped = true
}

// LD DE,nn
func (cpu *CPU) op11() {
	data := cpu.readWord(cpu.PC + 1)
	cpu.DE.set(data)
}

// LD (DE),A
func (cpu *CPU) op12() {
	addr := cpu.DE.get()
	data := cpu.AF.getHi()
	cpu.ld8(addr, data)
}

// INC DE
func (cpu *CPU) op13() {
	cpu.DE.inc()
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

// RRA
func (cpu *CPU) op1F() {
	cpu.rra()
}

// JR NZ,e
func (cpu *CPU) op20() {
	offset := cpu.read(cpu.PC + 1)
	cond := !cpu.getFlag(FLAG_Z)
	cpu.jrIf(offset, cond)
}

// LD HL,nn
func (cpu *CPU) op21() {
	data := cpu.readWord(cpu.PC + 1)
	cpu.HL.set(data)
}

// LD (HL+),A
func (cpu *CPU) op22() {
	addr := cpu.HL.get()
	data := cpu.AF.getHi()
	cpu.write(addr, data)

	cpu.HL.inc()
}

// INC HL
func (cpu *CPU) op23() {
	cpu.HL.inc()
}

// DEC H
func (cpu *CPU) op25() {
	cpu.dec8(&cpu.HL.hiReg)
}

// LD H,n
func (cpu *CPU) op26() {
	data := cpu.read(cpu.PC + 1)
	cpu.HL.setHi(data)
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

// DEC L
func (cpu *CPU) op2D() {
	cpu.dec8(&cpu.HL.loReg)
}

// CPL
func (cpu *CPU) op2F() {
	cpu.cpl()
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

// DEC A
func (cpu *CPU) op3D() {
	cpu.dec8(&cpu.AF.hiReg)
}

// LD A,n
func (cpu *CPU) op3E() {
	data := cpu.read(cpu.PC + 1)
	cpu.AF.setHi(data)
}

// LD B,(HL)
func (cpu *CPU) op46() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.BC.setHi(data)
}

// LD B,A
func (cpu *CPU) op47() {
	data := cpu.AF.getHi()
	cpu.BC.setHi(data)
}

// LD C,(HL)
func (cpu *CPU) op4E() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.BC.setLo(data)
}

// LD C,A
func (cpu *CPU) op4F() {
	data := cpu.AF.getHi()
	cpu.BC.setLo(data)
}

// LD D,(HL)
func (cpu *CPU) op56() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.DE.setHi(data)
}

// LD D,A
func (cpu *CPU) op57() {
	data := cpu.AF.getHi()
	cpu.DE.setHi(data)
}

// LD E,L
func (cpu *CPU) op5D() {
	data := cpu.HL.getLo()
	cpu.DE.setLo(data)
}

// LD E,A
func (cpu *CPU) op5F() {
	data := cpu.AF.getHi()
	cpu.DE.setLo(data)
}

// LD H,B
func (cpu *CPU) op60() {
	data := cpu.BC.getHi()
	cpu.HL.setHi(data)
}

// LD H,D
func (cpu *CPU) op62() {
	data := cpu.DE.getHi()
	cpu.HL.setHi(data)
}

// LD H,E
func (cpu *CPU) op63() {
	data := cpu.DE.getLo()
	cpu.HL.setHi(data)
}

// LD H,(HL)
func (cpu *CPU) op66() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.HL.setHi(data)
}

// LD H,A
func (cpu *CPU) op67() {
	data := cpu.AF.getHi()
	cpu.HL.setHi(data)
}

// LD L,C
func (cpu *CPU) op69() {
	data := cpu.BC.getLo()
	cpu.HL.setLo(data)
}

// LD L,E
func (cpu *CPU) op6B() {
	data := cpu.DE.getLo()
	cpu.HL.setLo(data)
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

// LD L,A
func (cpu *CPU) op6F() {
	data := cpu.AF.getHi()
	cpu.HL.setLo(data)
}

// LD (HL),B
func (cpu *CPU) op70() {
	addr := cpu.HL.get()
	data := cpu.BC.getHi()
	cpu.ld8(addr, data)
}

// LD (HL),D
func (cpu *CPU) op72() {
	addr := cpu.HL.get()
	data := cpu.DE.getHi()
	cpu.ld8(addr, data)
}

// LD (HL),E
func (cpu *CPU) op73() {
	addr := cpu.HL.get()
	data := cpu.DE.getLo()
	cpu.ld8(addr, data)
}

// LD (HL),H
func (cpu *CPU) op74() {
	addr := cpu.HL.get()
	data := cpu.HL.getHi()
	cpu.ld8(addr, data)
}

// LD (HL),L
func (cpu *CPU) op75() {
	addr := cpu.HL.get()
	data := cpu.HL.getLo()
	cpu.ld8(addr, data)
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

// LD A,C
func (cpu *CPU) op79() {
	data := cpu.BC.getLo()
	cpu.AF.setHi(data)
}

// LD A,D
func (cpu *CPU) op7A() {
	data := cpu.DE.getHi()
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

// ADD A,B
func (cpu *CPU) op80() {
	b := cpu.BC.getHi()
	cpu.add(b)
}

// ADC A,(HL)
func (cpu *CPU) op8E() {
	addr := cpu.HL.get()
	hl := cpu.read(addr)
	cpu.adc(hl)
}

// SUB E
func (cpu *CPU) op93() {
	e := cpu.DE.getLo()
	cpu.sub(e)
}

// AND E
func (cpu *CPU) opA3() {
	e := cpu.DE.getLo()
	cpu.and(e)
}

// XOR (HL)
func (cpu *CPU) opAE() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.xor(data)
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

// RET NZ
func (cpu *CPU) opC0() {
	cond := !cpu.getFlag(FLAG_Z)
	cpu.retIf(cond)

	cpu.PC -= instructions[0xC0].length
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

// Prefix instructions
func (cpu *CPU) opCB() {
	op := cpu.read(cpu.PC + 1)

	// Determine register
	// The registers used in these instructions are determined by the LSB of
	// opcode as follows:
	//
	// LSB: 0 1 2 3 4 5  6   7
	// reg: B C D E H L (HL) A
	//
	// This pattern repeats for opcodes with LSB 8 through F (x8-xF).
	var reg *byte
	switch op % 8 {
	case 0x0:
		reg = &cpu.BC.hiReg.value
	case 0x1:
		reg = &cpu.BC.loReg.value
	case 0x2:
		reg = &cpu.DE.hiReg.value
	case 0x3:
		reg = &cpu.DE.loReg.value
	case 0x4:
		reg = &cpu.HL.hiReg.value
	case 0x5:
		reg = &cpu.HL.loReg.value
	case 0x6:
		addr := cpu.HL.get()
		cpu.cycles++
		reg = &cpu.bus.CartRom[addr]
	case 0x7:
		reg = &cpu.AF.hiReg.value
	}

	// Determine bit/instruction, execute instruction
	if op <= 0x3F {
		// Determine instruction
		// For all opcodes <= 0x3F, a different CPU instruction is used for
		// each set of 8 consecutive opcodes (e.g.: 0x00-0x07, 0x18-0x1F)
		switch op / 8 {
		case 0:
			// RLC
			cpu.rlc(reg)
		case 1:
			// RRC
			cpu.rrc(reg)
		case 2:
			// RL
			cpu.rl(reg)
		case 3:
			// RR
			cpu.rr(reg)
		case 4:
			// SLA
			cpu.sla(reg)
		case 5:
			// SRA
			cpu.sra(reg)
		case 6:
			// SWAP
			cpu.swap(reg)
		case 7:
			// SRL
			cpu.srl(reg)
		}
	} else {
		// Determine bit
		//
		// Every $CB-prefixed instruction of opcode >= 0x40 operates on a
		// certain bit of the specified data. Opcodes 0x40-0x47 each operate on
		// bit 0. The next 8 opcodes operate on bit 1, and so on for all bits 0-7.
		// This pattern repeats a total of 3 times through 0x40 and 0xFF.
		//
		// 0 0 0 0 0 0 0 0 1 1 1 1 1 1 1 1
		// 2 2 2 2 2 2 2 2 3 3 3 3 3 3 3 3
		// 4 4 4 4 4 4 4 4 5 5 5 5 5 5 5 5
		// 6 6 6 6 6 6 6 6 7 7 7 7 7 7 7 7
		bit := int((op / 8) % 8)

		// Determine instruction
		if op >= 0x40 && op <= 0x7F {
			// BIT
			cpu.bit(bit, reg)
		} else if op >= 0x80 && op <= 0xBF {
			// RES
			cpu.res(bit, reg)
		} else if op >= 0xC0 && op <= 0xFF {
			// SET
			cpu.set(bit, reg)
		}
	}
}

// CALL nn
func (cpu *CPU) opCD() {
	addr := cpu.readWord(cpu.PC + 1)
	cpu.call(addr)

	cpu.PC -= instructions[0xCD].length
}

// ADC A,n
func (cpu *CPU) opCE() {
	n := cpu.read(cpu.PC + 1)
	cpu.adc(n)
}

// RET NC
func (cpu *CPU) opD0() {
	cond := !cpu.getFlag(FLAG_C)
	cpu.retIf(cond)

	cpu.PC -= instructions[0xD0].length
}

// POP DE
func (cpu *CPU) opD1() {
	data := cpu.pop()
	cpu.DE.set(data)
}

// PUSH DE
func (cpu *CPU) opD5() {
	data := cpu.DE.get()
	cpu.push(data)
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

// LD SP,HL
func (cpu *CPU) opF9() {
	data := cpu.HL.get()
	cpu.SP = data
}

// LD A,(nn)
func (cpu *CPU) opFA() {
	nn := cpu.readWord(cpu.PC + 1)
	data := cpu.read(nn)
	cpu.AF.setHi(data)
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
