package gb

import (
	"fmt"
)

const lineTemplate string = "[%#04x]:\t%s\t%s "

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
			// NOP
		case 0x21:
			// LD HL,nn
			msg += fmt.Sprintf("HL,0x%04X", word)
		case 0x31:
			// LD SP,nn
			msg += fmt.Sprintf("SP,0x%04X", word)
		case 0x3E:
			// LD A,n
			msg += fmt.Sprintf("A,0x%02X", op1)
		case 0xC3:
			// JP nn
			msg += fmt.Sprintf("0x%04X", word)
		case 0xCD:
			// CALL nn
			msg += fmt.Sprintf("(0x%04X)", word)
		case 0xE0:
			// LDH (n),A
			msg += fmt.Sprintf("(0xFF%02X),A", op1)
		case 0xEA:
			// LD (nn),A
			msg += fmt.Sprintf("(0x%04X),A", word)
		case 0xF3:
			// DI
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
