package gb

// attachCPU attaches the given CPU to the Game Boy's bus
func (gb *GameBoy) attachCPU(cpu *CPU) {
	gb.Cpu = cpu
	gb.Cpu.bus = gb
}

// cpuRead allows the CPU to read from the system bus at the given memory
// address
// TODO: limit access
func (gb *GameBoy) cpuRead(addr uint16) byte {
	addr &= 0xFFFF
	data := gb.CartRom[addr]
	return data
}

// cpuWrite allows the CPU to write data to the system bus at the given memory
// address
// TODO: limit access
func (gb *GameBoy) cpuWrite(addr uint16, data byte) {
	addr &= 0xFFFF
	gb.CartRom[addr] = data
}
