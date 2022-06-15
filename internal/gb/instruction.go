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
	instructions   = [INSTRUCTION_COUNT]inst{}
	pfInstructions = [INSTRUCTION_COUNT]inst{} // prefixed $CB
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
	instructions[0x18] = inst{"JR", 2, 3, cpu.op18}
	instructions[0x1C] = inst{"INC", 1, 1, cpu.op1C}
	instructions[0x1F] = inst{"RRA", 1, 1, cpu.op1F}
	instructions[0x20] = inst{"JR", 2, 2, cpu.op20}
	instructions[0x21] = inst{"LD", 3, 3, cpu.op21}
	instructions[0x23] = inst{"INC", 1, 2, cpu.op23}
	instructions[0x26] = inst{"LD", 2, 2, cpu.op26}
	instructions[0x28] = inst{"JR", 2, 2, cpu.op28}
	instructions[0x2A] = inst{"LD", 1, 2, cpu.op2A}
	instructions[0x2D] = inst{"DEC", 1, 1, cpu.op2D}
	instructions[0x30] = inst{"JR", 2, 2, cpu.op30}
	instructions[0x31] = inst{"LD", 3, 3, cpu.op31}
	instructions[0x32] = inst{"LD", 1, 2, cpu.op32}
	instructions[0x36] = inst{"LD", 2, 3, cpu.op36}
	instructions[0x38] = inst{"JR", 2, 2, cpu.op38}
	instructions[0x3C] = inst{"INC", 1, 1, cpu.op3C}
	instructions[0x3D] = inst{"DEC", 1, 1, cpu.op3D}
	instructions[0x3E] = inst{"LD", 2, 2, cpu.op3E}
	instructions[0x46] = inst{"LD", 1, 2, cpu.op46}
	instructions[0x4E] = inst{"LD", 1, 2, cpu.op4E}
	instructions[0x56] = inst{"LD", 1, 2, cpu.op56}
	instructions[0x5D] = inst{"LD", 1, 1, cpu.op5D}
	instructions[0x60] = inst{"LD", 1, 1, cpu.op60}
	instructions[0x66] = inst{"LD", 1, 2, cpu.op66}
	instructions[0x6C] = inst{"LD", 1, 1, cpu.op6C}
	instructions[0x6E] = inst{"LD", 1, 2, cpu.op6E}
	instructions[0x76] = inst{"HALT", 1, 1, cpu.op76}
	instructions[0x78] = inst{"LD", 1, 1, cpu.op78}
	instructions[0x7B] = inst{"LD", 1, 1, cpu.op7B}
	instructions[0x7C] = inst{"LD", 1, 1, cpu.op7C}
	instructions[0x7D] = inst{"LD", 1, 1, cpu.op7D}
	instructions[0x7E] = inst{"LD", 1, 2, cpu.op7E}
	instructions[0x8E] = inst{"ADC", 1, 2, cpu.op8E}
	instructions[0xA3] = inst{"AND", 1, 1, cpu.opA3}
	instructions[0xAE] = inst{"XOR", 1, 2, cpu.opAE}
	instructions[0xB1] = inst{"OR", 1, 1, cpu.opB1}
	instructions[0xB7] = inst{"OR", 1, 1, cpu.opB7}
	instructions[0xC3] = inst{"JP", 3, 4, cpu.opC3}
	instructions[0xC5] = inst{"PUSH", 1, 4, cpu.opC5}
	instructions[0xC6] = inst{"ADD", 2, 2, cpu.opC6}
	instructions[0xC9] = inst{"RET", 1, 4, cpu.opC9}
	instructions[0xCB] = inst{"", 0, 0, cpu.opCB} // prefix 0xCB
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
	instructions[0xFE] = inst{"CP", 2, 2, cpu.opFE}
	instructions[0xFF] = inst{"RST", 1, 4, cpu.opFF}
}

// setupInstuctionLookup defines all legal CPU instructions prefixed with 0xCB
func (cpu *CPU) setupPfInstructionLookup() {
	pfinstructions[0xFF] = inst{"SET", 2, 2, func() { cpu.set(7, &cpu.AF.hiReg) }}

	// log.Fatal on unimplemented opcodes development
	for i, inst := range pfInstructions {
		if inst.exec == nil {
			pfInstructions[i].exec = func() {
				op := byte(i)
				cpu.bus.logger.Fatalf("unimplemented op: CB %02X %#08b\n", op, op)
			}
		}
	}
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

// INC HL
func (cpu *CPU) op23() {
	cpu.HL.inc()
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

// LD C,(HL)
func (cpu *CPU) op4E() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.BC.setLo(data)
}

// LD D,(HL)
func (cpu *CPU) op56() {
	addr := cpu.HL.get()
	data := cpu.read(addr)
	cpu.DE.setHi(data)
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
	addr := cpu.HL.get()
	hl := cpu.read(addr)
	cpu.adc(hl)
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
	b := cpu.read(cpu.PC + 1)
	pfInstructions[b].exec()
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
