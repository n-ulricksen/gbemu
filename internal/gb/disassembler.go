package gb

import "fmt"

var prefix string = "%#04x\t[%02X]:\t%s "

// disassemble disassembles the system's loaded cartridge ROM between addresses
// 'start' and 'end', and stores the result in the 'gb.disassembly'
func (gb *GameBoy) disassemble(start uint16, end uint16) error {
	if int(end) >= len(gb.CartRom) {
		return errAddrOutOfRange
	}

	disassembly := make([]string, len(gb.CartRom))

	complete := false
	for addr := start; !complete; addr++ {
		op := gb.CartRom[addr]
		// op1 := gb.CartRom[addr+1]
		// op2 := gb.CartRom[addr+2]
		op3 := gb.Cpu.readWord(addr + 1) // next word
		msg := fmt.Sprintf(prefix, addr, op, instructions[op].name)

		switch op {
		case 0x00:
		case 0xC3:
			msg += fmt.Sprintf("%#04x", op3)
		}

		disassembly[addr] = msg
		if addr >= end {
			complete = true
		}
	}

	gb.disassembly = disassembly
	return nil
}
