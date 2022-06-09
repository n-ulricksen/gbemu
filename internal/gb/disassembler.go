package gb

import (
	"fmt"
)

const lineTemplate string = "[%#04x]\t%s\t%s "

// disassemble disassembles the system's loaded cartridge ROM between addresses
// 'start' and 'end', and stores the result in the 'gb.disassembly'
func (gb *GameBoy) disassemble(start uint16, end uint16) error {
	if int(end) >= len(gb.CartRom) {
		return errAddrOutOfRange
	}

	disassembly := make([]string, len(gb.CartRom))

	addr := start
	for {
		op := gb.CartRom[addr]
		op1 := gb.CartRom[addr+1]
		op2 := gb.CartRom[addr+2]
		word := gb.Cpu.readWord(addr + 1) // next word
		inst := instructions[op]

		opString := fmt.Sprintf("%02X", op)
		if inst.length > 1 {
			opString += fmt.Sprintf(" %02X", op1)
		}
		if inst.length > 2 {
			opString += fmt.Sprintf(" %02X", op2)
		}

		msg := fmt.Sprintf(lineTemplate, addr, opString, inst.name)

		switch op {
		case 0x00:
		case 0x31:
			msg += fmt.Sprintf("SP, #$%x", word)
		case 0xC3:
			msg += fmt.Sprintf("#$%04x", word)
		case 0xF3:
		}

		disassembly[addr] = msg

		addr++
		if addr >= end {
			break
		}
	}

	gb.disassembly = disassembly
	return nil
}
