package gb

// CPU represents a Sharp SM83 CPU core
type CPU struct {
	AF register // Accumulator and flags
	BC register
	DE register
	HL register

	SP uint16 // Stack pointer
	PC uint16 // Program counter

	IME    bool // Interrupt master enable flag
	halted bool // Used to pause CPU execution until an interrupt occurs

	bus *GameBoy // 16-bit address, 8-bit data bus

	cycles int
}

// newCPU returns a SM83 CPU
func newCPU() *CPU {
	cpu := &CPU{
		AF: register{name: "AF", hiReg: newReg8Bit("A"), loReg: newReg8Bit("F")},
		BC: register{name: "BC", hiReg: newReg8Bit("B"), loReg: newReg8Bit("C")},
		DE: register{name: "DE", hiReg: newReg8Bit("D"), loReg: newReg8Bit("E")},
		HL: register{name: "HL", hiReg: newReg8Bit("H"), loReg: newReg8Bit("L")},
	}
	cpu.setupInstructionLookup()

	return cpu
}

// execNextInst fetches the opcode at the current program counter and executes
// the appropriate CPU instruction
func (cpu *CPU) execNextInst() {
	// fetch
	op := cpu.read(cpu.PC)

	// log
	cpu.logInstruction()

	// decode & execute
	cpu.decodeAndExecute(op)
}

// decodeAndExecude decodes the given opcode and executes its CPU instruction
func (cpu *CPU) decodeAndExecute(op byte) {
	inst := instructions[op]
	if inst.exec == nil {
		cpu.bus.logger.Fatalf("unimplemented op: %02X %#08b\n", op, op)
	}

	inst.exec()

	cpu.PC += inst.length
	cpu.cycles += inst.cycles
}

// logInstruction logs the disassembly for the current CPU instruction
func (cpu *CPU) logInstruction() {
	if cpu.bus.debugMode {
		logMsg := cpu.bus.disassembly[cpu.PC]
		cpu.bus.log(logMsg)
	}
}

// read reads 1 byte from the system bus at the given address
func (cpu *CPU) read(addr uint16) byte {
	return cpu.bus.cpuRead(addr)
}

// write writes 1 byte of data to the system bus at the given address
func (cpu *CPU) write(addr uint16, data byte) {
	cpu.bus.cpuWrite(addr, data)
}

// readWord reads 2 bytes from the system bus at the given address (little
// endian order)
func (cpu *CPU) readWord(addr uint16) uint16 {
	lo := uint16(cpu.read(addr))
	hi := uint16(cpu.read(addr + 1))
	return (hi << 8) | lo
}

// stackPush pushes a byte of data to the stack
func (cpu *CPU) stackPush(data byte) {
	cpu.SP--
	cpu.write(cpu.SP, data)
}

// stackPop pops a byte of data from the stack
func (cpu *CPU) stackPop() byte {
	data := cpu.read(cpu.SP)
	cpu.SP++
	return data
}
