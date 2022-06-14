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
		case 0x01:
			// LD BC,nn
			msg += fmt.Sprintf("BC,0x%04X", word)
		case 0x02:
			// LD (BC),A
			msg += "(BC),A"
		case 0x03:
			// INC BC
			msg += "BC"
		case 0x05:
			// DEC B
			msg += "B"
		case 0x06:
			// LD B,n
			msg += fmt.Sprintf("B,0x%02X", op1)
		case 0x07:
			// RLCA
		case 0x18:
			// JR e
			msg += fmt.Sprintf("0x%02X", op1)
		case 0x1C:
			// INC E
			msg += "E"
		case 0x1F:
			// RRA
		case 0x20:
			// JR NZ,e
			msg += fmt.Sprintf("NZ,0x%02X", op1)
		case 0x21:
			// LD HL,nn
			msg += fmt.Sprintf("HL,0x%04X", word)
		case 0x23:
			// INC HL
			msg += "HL"
		case 0x26:
			// LD H,n
			msg += fmt.Sprintf("H,0x%02X", op1)
		case 0x28:
			// JR Z,e
			msg += fmt.Sprintf("Z,0x%02X", op1)
		case 0x2A:
			// LD A,(HL+)
			msg += "A,(HL+)"
		case 0x2D:
			// DEC L
			msg += "L"
		case 0x30:
			// JR NC,e
			msg += fmt.Sprintf("NC,0x%02X", op1)
		case 0x31:
			// LD SP,nn
			msg += fmt.Sprintf("SP,0x%04X", word)
		case 0x32:
			// LD (HL-),A
			msg += "(HL-),A"
		case 0x36:
			// LD (HL),n
			msg += fmt.Sprintf("(HL),0x%02X", op1)
		case 0x38:
			// JR C,e
			msg += fmt.Sprintf("C,0x%02X", op1)
		case 0x3C:
			// INC A
			msg += "A"
		case 0x3D:
			// DEC A
			msg += "A"
		case 0x3E:
			// LD A,n
			msg += fmt.Sprintf("A,0x%02X", op1)
		case 0x46:
			// LD B,(HL)
			msg += "B,(HL)"
		case 0x4E:
			// LD C,(HL)
			msg += "C,(HL)"
		case 0x56:
			// LD D,(HL)
			msg += "D,(HL)"
		case 0x5D:
			// LD E,L
			msg += "E,L"
		case 0x60:
			// LD H,B
			msg += "H,B"
		case 0x66:
			// LD H,(HL)
			msg += "H,(HL)"
		case 0x6C:
			// LD L,H
			msg += "L,H"
		case 0x6E:
			// LD L,(HL)
			msg += "L,(HL)"
		case 0x78:
			// LD A,B
			msg += "A,B"
		case 0x76:
			// HALT
		case 0x7B:
			// LD A,E
			msg += "A,E"
		case 0x7C:
			// LD A,H
			msg += "A,H"
		case 0x7D:
			// LD A,L
			msg += "A,L"
		case 0x7E:
			// LD A,(HL)
			msg += "A,(HL)"
		case 0x8E:
			// ADC A,(HL)
			msg += "A,(HL)"
		case 0xA3:
			// AND E
			msg += "E"
		case 0xAE:
			// XOR (HL)
			msg += "(HL)"
		case 0xB1:
			// OR C
			msg += "C"
		case 0xC3:
			// JP nn
			msg += fmt.Sprintf("0x%04X", word)
		case 0xC5:
			// PUSH BC
			msg += "BC"
		case 0xC6:
			// ADD A,n
			msg += fmt.Sprintf("A,0x%02X", op1)
		case 0xC9:
			// RET
		case 0xCD:
			// CALL nn
			msg += fmt.Sprintf("(0x%04X)", word)
		case 0xCE:
			// ADC A,n
			msg += fmt.Sprintf("A,0x%02X", op1)
		case 0xD0:
			// RET NC
			msg += "NC"
		case 0xD1:
			// POP DE
			msg += "DE"
		case 0xD5:
			// PUSH DE
			msg += "DE"
		case 0xD6:
			// SUB n
			msg += fmt.Sprintf("0x%02X", op1)
		case 0xE0:
			// LDH (n),A
			msg += fmt.Sprintf("(0xFF%02X),A", op1)
		case 0xE1:
			// POP HL
			msg += "HL"
		case 0xE5:
			// PUSH HL
			msg += "HL"
		case 0xEA:
			// LD (nn),A
			msg += fmt.Sprintf("(0x%04X),A", word)
		case 0xEE:
			// XOR n
			msg += fmt.Sprintf("0x%02X", op1)
		case 0xF0:
			// LDH A,(n)
			msg += fmt.Sprintf("A,(0xFF%02X)", op1)
		case 0xF1:
			// POP AF
			msg += "AF"
		case 0xF3:
			// DI
		case 0xF5:
			// PUSH AF
			msg += "AF"
		case 0xFE:
			// CP n
			msg += fmt.Sprintf("0x%02X", op1)
		case 0xFF:
			// RST n
			msg += "0x38"
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
