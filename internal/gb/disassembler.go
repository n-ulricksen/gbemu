package gb

import (
	"fmt"
)

const lineTemplate string = "[%#04x]:\t%s\t%s "

// disassemble disassembles the system's loaded cartridge ROM between addresses
// 'start' and 'end', and stores the result in the 'gb.disassembly'
//
// see 'instruction.go' or https://gbdev.io/gb-opcodes/optables/ for details
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

		if op == 0xCB {
			// Prefix instruction
			inst.name = getPrefixInstructionName(op1)
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
		case 0x0A:
			// LD A,(BC)
			msg += "A,(BC)"
		case 0x0E:
			// LD C,n
			msg += fmt.Sprintf("B,0x%02X", op1)
		case 0x0F:
			// RRCA
		case 0x10:
			// STOP
		case 0x11:
			// LD DE,nn
			msg += fmt.Sprintf("DE,0x%04X", word)
		case 0x12:
			// LD (DE),A
			msg += "(DE),A"
		case 0x13:
			// INC DE
			msg += "DE"
		case 0x18:
			// JR e
			msg += fmt.Sprintf("0x%02X", op1)
		case 0x1A:
			// LD A,(DE)
			msg += "A,(DE)"
		case 0x1C:
			// INC E
			msg += "E"
		case 0x1D:
			// DEC E
			msg += "E"
		case 0x1F:
			// RRA
		case 0x20:
			// JR NZ,e
			msg += fmt.Sprintf("NZ,0x%02X", op1)
		case 0x21:
			// LD HL,nn
			msg += fmt.Sprintf("HL,0x%04X", word)
		case 0x22:
			// LD (HL+),A
			msg += "(HL+),A"
		case 0x23:
			// INC HL
			msg += "HL"
		case 0x24:
			// INC H
			msg += "H"
		case 0x25:
			// DEC H
			msg += "H"
		case 0x26:
			// LD H,n
			msg += fmt.Sprintf("H,0x%02X", op1)
		case 0x28:
			// JR Z,e
			msg += fmt.Sprintf("Z,0x%02X", op1)
		case 0x29:
			// ADD HL,HL
			msg += "HL,HL"
		case 0x2A:
			// LD A,(HL+)
			msg += "A,(HL+)"
		case 0x2C:
			// INC L
			msg += "L"
		case 0x2D:
			// DEC L
			msg += "L"
		case 0x2F:
			// CPL
		case 0x30:
			// JR NC,e
			msg += fmt.Sprintf("NC,0x%02X", op1)
		case 0x31:
			// LD SP,nn
			msg += fmt.Sprintf("SP,0x%04X", word)
		case 0x32:
			// LD (HL-),A
			msg += "(HL-),A"
		case 0x35:
			// DEC (HL)
			msg += "(HL)"
		case 0x36:
			// LD (HL),n
			msg += fmt.Sprintf("(HL),0x%02X", op1)
		case 0x38:
			// JR C,e
			msg += fmt.Sprintf("C,0x%02X", op1)
		case 0x3A:
			// LD A,(HL-)
			msg += "A,(HL-)"
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
		case 0x47:
			// LD B,A
			msg += "B,A"
		case 0x4E:
			// LD C,(HL)
			msg += "C,(HL)"
		case 0x4F:
			// LD C,A
			msg += "C,A"
		case 0x56:
			// LD D,(HL)
			msg += "D,(HL)"
		case 0x57:
			// LD D,A
			msg += "D,A"
		case 0x5D:
			// LD E,L
			msg += "E,L"
		case 0x5F:
			// LD E,A
			msg += "E,A"
		case 0x60:
			// LD H,B
			msg += "H,B"
		case 0x62:
			// LD H,D
			msg += "H,D"
		case 0x63:
			// LD H,E
			msg += "H,E"
		case 0x66:
			// LD H,(HL)
			msg += "H,(HL)"
		case 0x67:
			// LD H,A
			msg += "H,A"
		case 0x69:
			// LD L,C
			msg += "L,C"
		case 0x6B:
			// LD L,E
			msg += "L,E"
		case 0x6C:
			// LD L,H
			msg += "L,H"
		case 0x6E:
			// LD L,(HL)
			msg += "L,(HL)"
		case 0x6F:
			// LD L,A
			msg += "L,A"
		case 0x70:
			// LD (HL),B
			msg += "(HL),B"
		case 0x71:
			// LD (HL),C
			msg += "(HL),C"
		case 0x72:
			// LD (HL),D
			msg += "(HL),D"
		case 0x73:
			// LD (HL),E
			msg += "(HL),E"
		case 0x74:
			// LD (HL),H
			msg += "(HL),H"
		case 0x75:
			// LD (HL),L
			msg += "(HL),L"
		case 0x76:
			// HALT
		case 0x77:
			// LD (HL),A
			msg += "(HL),A"
		case 0x78:
			// LD A,B
			msg += "A,B"
		case 0x79:
			// LD A,C
			msg += "A,C"
		case 0x7A:
			// LD A,D
			msg += "A,D"
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
		case 0x80:
			// ADD A,B
			msg += "A,B"
		case 0x8E:
			// ADC A,(HL)
			msg += "A,(HL)"
		case 0x93:
			// SUB E
			msg += "E"
		case 0xA3:
			// AND E
			msg += "E"
		case 0xA9:
			// XOR C
			msg += "C"
		case 0xAE:
			// XOR (HL)
			msg += "(HL)"
		case 0xB1:
			// OR C
			msg += "C"
		case 0xB6:
			// OR (HL)
			msg += "(HL)"
		case 0xC0:
			// RET NZ
			msg += "NZ"
		case 0xC1:
			// POP BC
			msg += "BC"
		case 0xC3:
			// JP nn
			msg += fmt.Sprintf("0x%04X", word)
		case 0xC4:
			// CALL NZ,nn
			msg += fmt.Sprintf("NZ,0x%04X", word)
		case 0xC5:
			// PUSH BC
			msg += "BC"
		case 0xC6:
			// ADD A,n
			msg += fmt.Sprintf("A,0x%02X", op1)
		case 0xC8:
			// RET Z
			msg += "Z"
		case 0xC9:
			// RET
		case 0xCA:
			// JP Z,nn
			msg += fmt.Sprintf("Z,0x%04X", word)
		case 0xCB:
			// Prefix instructions
			var reg string
			switch op1 % 8 {
			case 0:
				reg = "B"
			case 1:
				reg = "C"
			case 2:
				reg = "D"
			case 3:
				reg = "E"
			case 4:
				reg = "H"
			case 5:
				reg = "L"
			case 6:
				reg = "(HL)"
			case 7:
				reg = "A"
			}

			if op1 <= 0x3F {
				msg += fmt.Sprintf("%s", reg)
			} else {
				// Determine bit
				b := int((op1 / 8) % 8)
				bit := fmt.Sprintf("%d", b)

				msg += fmt.Sprintf("%s,%s", bit, reg)
			}
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
		case 0xD2:
			// JP NC,nn
			msg += fmt.Sprintf("NC,0x%04X", word)
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
		case 0xE6:
			// AND n
			msg += fmt.Sprintf("0x%02X", op1)
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
		case 0xF9:
			// LD SP,HL
			msg += "SP,HL"
		case 0xFA:
			// LD A,(nn)
			msg += fmt.Sprintf("A,(0x%04X)", word)
		case 0xFE:
			// CP n
			msg += fmt.Sprintf("0x%02X", op1)
		case 0xFF:
			// RST 0x38
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

func getPrefixInstructionName(op byte) string {
	var name string

	if op <= 0x3F {
		switch op / 8 {
		case 0:
			name = "RLC"
		case 1:
			name = "RRC"
		case 2:
			name = "RL"
		case 3:
			name = "RR"
		case 4:
			name = "SLA"
		case 5:
			name = "SRA"
		case 6:
			name = "SWAP"
		case 7:
			name = "SRL"
		}
	} else {
		if op >= 0x40 && op <= 0x7F {
			name = "BIT"
		} else if op >= 0x80 && op <= 0xBF {
			name = "RES"
		} else if op >= 0xC0 && op <= 0xFF {
			name = "SET"
		}
	}

	return name
}
